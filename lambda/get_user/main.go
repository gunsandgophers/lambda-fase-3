package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"encoding/json"

	"github.com/gunsandgophers/lambda-fase-3/internal/domain"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	cpf, exists := request.QueryStringParameters["cpf"]
	if !exists {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body: "CPF undefined",
		}, nil
	}
	customerService, err := domain.NewAwsCustomerService(
		"us-east-1",
		"us-east-1_3ofqHwfxr",
	)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body: err.Error(),
		}, nil
	}
	getCustomer := domain.NewGetCustomerIdUC(customerService)
	customer, err := getCustomer.Execute(cpf)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body: err.Error(),
		}, nil
	}
	body, err := json.Marshal(customer)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body: err.Error(),
		}, nil
	}
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body: string(body),
	}
	return response, nil
}
