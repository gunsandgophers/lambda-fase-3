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
	customerInput := &domain.CreateCustomerInput{}
	err := json.Unmarshal([]byte(request.Body), customerInput)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body: "Invalida body",
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
	createCustomer := domain.NewCreateCustomerUC(customerService)
	customer, err := createCustomer.Execute(customerInput)
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
		StatusCode: 201,
		Body: string(body),
	}
	return response, nil
}
