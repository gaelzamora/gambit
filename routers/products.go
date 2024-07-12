package routers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gaelzamora/gambit/bd"
	"github.com/gaelzamora/gambit/models"
)

func InsertProduct(body string, User string) (int, string) {
	var t models.Product
 
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos"+err.Error()
	}

	if len(t.ProdTitle)==0 {
		return 400, "Debe especificar el Nombre (Title) del Producto"
	}

	isAdmin, msg := bd.IsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	result, err2 := bd.InsertProduct(t)

	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el registro del producto "+t.ProdTitle+" > "+err2.Error()
	}

	return 200, "{ ProductID: "+strconv.Itoa(int(result))+" }"
}

func UpdateProduct(body string, User string, id int) (int, string) {
	var t models.Product

	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos "+err.Error()
	}

	isAdmin, msg := bd.IsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	t.ProdId = id
	err2 := bd.UpdateProduct(t)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el UPDATE del producto " + strconv.Itoa(id) + " > " + err2.Error()
	}

	return 200, "Update OK"
}

func DeleteProduct(User string, id int) (int, string) {
	
	isAdmin, msg := bd.IsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	err2 := bd.DeleteProduct(id)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el UPDATE del producto " + strconv.Itoa(id) + " > " + err2.Error()
	}

	return 200, "Delete OK"

}

func SelectProduct(request events.APIGatewayV2HTTPRequest) (int, string) {
	var t models.Product
	var page, pagesize int
	var orderType, orderField string

	fmt.Println("Entro a Select Product")

	param := request.QueryStringParameters
	
	page, _ = strconv.Atoi(param["page"])
	pagesize, _ = strconv.Atoi(param["pageSize"])
	orderType = param["orderType"] // D = Desc, A o nil = ASC
	orderField = param["orderField"] // 'I' Id, 'T' Title, 'D' Description, 'F' Created At
									 // 'P' Price, 'C' CategId, 'S' Stock
	if !strings.Contains("ITDFPCS", orderField) {
		orderField=""
	}

	var choice string
	if len(param["prodId"])>0 {
		choice="P"
		t.ProdId, _ = strconv.Atoi(param["prodId"])
	}

	if len(param["search"])>0 {
		choice="S"
		t.ProdSearch = param["search"]
	}

	if len(param["categId"])>0 {
		choice="C"
		t.ProdCategId, _ = strconv.Atoi(param["categId"])
	}

	if len(param["slug"])>0 {
		choice="U"
		t.ProdPath = param["slug"]
	}

	if len(param["slugCateg"])>0 {
		choice="K"
		t.ProdCategPath = param["slugCateg"]
	}

	fmt.Println(param)

	result, err2 := bd.SelectProduct(t, choice, page, pagesize, orderType, orderField)

	fmt.Println(result)

	if err2 != nil {
		return 400, "Ocurrio un error al intentar capturar los resultados de la busqueda de tipo " + choice + "' en productos > "+err2.Error()
	}

	Product, err3 := json.Marshal(result)

	fmt.Println(Product)

	if err3 != nil {
		return 400, "Ocurrio un error al intentar convertir en JSON el registro de Productos"
	}

	return 200, string(Product)
}


func UpdateStock(body string, User string, id int) (int, string) {
	var t models.Product

	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos "+err.Error()
	}

	isAdmin, msg := bd.IsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	t.ProdId = id
	err2 := bd.UpdateStock(t)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el UPDATE del stock " + strconv.Itoa(id) + " > " + err2.Error()
	}

	return 200, "Update OK"
}
