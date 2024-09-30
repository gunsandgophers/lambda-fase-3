package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"encoding/json"

	"github.com/gunsandgophers/lambda-fase-3/internal/config"
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
	region := config.GetEnv("TC_AWS_REGION", "")
	userPoolId := config.GetEnv("TC_AWS_COGNITO_USER_POOL_ID", "")
	customerService, err := domain.NewAwsCustomerService(
		region,
		userPoolId,
	)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body: err.Error(),
		}, nil
	}
	getCustomer := domain.NewGetCustomerUC(customerService)
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
