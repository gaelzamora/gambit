package handlers

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gaelzamora/gambit/auth"
	"github.com/gaelzamora/gambit/routers"
)

func Manejadores(path string, method string, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {
	
	fmt.Println("Voy a procesar "+path+" > "+method)

	id := request.PathParameters["id"]
	idn, _ := strconv.Atoi(id)

	isOK, statusCode, user := validoAuthorizacion(path, method, headers)

	if !isOK {
		fmt.Println("No esta bien")
		return statusCode, user
	}

	fmt.Println("Entro")

	if len(path) < 4 {
		fmt.Println(path)
        return 400, "Path too short"
    }

	fmt.Println(path)

	fmt.Println("path[0:4] = " + path[1:5])

	switch path[1:5] {
	case "user":
		fmt.Println("Entre a user")
		return ProcesoUser(body, path, method, user, id, request)
	case "prod":
		return ProcesoProducts(body, path, method, user, idn, request)
	case "stoc":
		return ProcesoStock(body, path, method, user, idn, request)
	case "addr":
		return ProcesoAddress(body, path, method, user, idn, request)
	case "cate":
		return ProcesoCategory(body, path, method, user, idn, request)
	case "orde":
		return ProcesoOrder(body, path, method, user, idn, request)
	}

	fmt.Println("Ni modo")

	return 400, "Method Invalid"
}

func validoAuthorizacion(path string, method string, headers map[string]string) (bool, int, string) {
	if (path == "/products" && method == "GET") || 
		(path == "/category" && method == "GET") {
			return true, 200, ""
		}

	token := headers["authorization"]
	if len(token)==0 {
		return false, 401, "Token requerido"
	}

	allOK, err, msg := auth.ValidoToken(token)

	if !allOK {
		if err != nil {
			fmt.Println("Error en el token"+err.Error())
			return false, 401, err.Error()
		} else {
			fmt.Println("Error en el token"+msg)
			return false, 401, msg
		}
	}

	fmt.Println("Token OK")

	return true, 200, msg
}

func ProcesoUser(body string, path string, method string, user string, id string, request events.APIGatewayV2HTTPRequest) (int, string) {
	if path == "/users/me" {
		switch method {
		case "PUT":
			return routers.UpdateUser(body, user)
		case "GET":
			return routers.SelectUser(body, user)
		}
	}

	if path == "/users" {
		if method == "GET" {
			return routers.SelectUsers(body, user, request)
		}
	}

	fmt.Println(path)
	fmt.Println(method)

	return 400, "Method Invalid"
}

func ProcesoProducts(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "POST":
		return routers.InsertProduct(body, user)
	case "PUT":
		return routers.UpdateProduct(body, user, id)
	case "DELETE":
		return routers.DeleteProduct(user, id)
	case "GET":
		return routers.SelectProduct(request)
	}
	
	return 400, "Method Invalid"
}

func ProcesoCategory(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
		case "POST":
			return routers.InsertCategory(body, user)

		case "PUT":
			return routers.UpdateCategory(body, user, id)
		
		case "DELETE":
			return routers.DeleteCategory(body, user, id)

		case "GET":
			return routers.SelectCategories(body, request)
	}

	return 400, "Method Invalid"
}

func ProcesoStock(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return routers.UpdateStock(body, user, id)
}

func ProcesoAddress(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "POST":
		return routers.InsertAddress(body, user)
	case "PUT":
		return routers.UpdateAddress(body, user, id)
	case "DELETE":
		return routers.DeleteAddress(user, id)
	case "GET":
		return routers.SelectAddress(user)
	}


	return 400, "Method Invalid"
}


func ProcesoOrder(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "POST":
		return routers.InsertOrder(body, user)
	case "GET":
		return routers.SelectOrders(user, request)
	}

	return 400, "Method Invalid"
}
