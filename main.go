package main

import (
	"context"
	"fmt"
	"log"

	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gaelzamora/gambit/awsgo"
	"github.com/gaelzamora/gambit/bd"

	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/gaelzamora/gambit/handlers"
)

func main() {
	lambda.Start(EjecutoLambda)
}


func EjecutoLambda(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	log.Println("Hello World")
	awsgo.InicializoAWS()

	if !ValidoParametros() {
		panic("Error en los parámetros. debe enviar 'SecretName', 'UrlPrefix'")
	}

	var res *events.APIGatewayProxyResponse
	path := strings.Replace(request.RawPath, os.Getenv("UrlPrefix"), "", -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	header := request.Headers

	bd.ReadSecret()

	status, message := handlers.Manejadores(path, method, body, header, request)

	headersResp := map[string]string{
		"Content-Type": "application/json",
	}

	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(message),
		Headers:    headersResp,
	}

	return res, nil

}

func ValidoParametros() bool {
	_, traeParametro := os.LookupEnv("SecretName")
	if !traeParametro {
		return traeParametro
	}

	fmt.Println(traeParametro)

	_, traeParametro = os.LookupEnv("UrlPrefix")
	if !traeParametro {
		return traeParametro
	}

	return traeParametro
}