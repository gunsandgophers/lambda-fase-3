package main

import (
	"fmt"

	"github.com/gunsandgophers/lambda-fase-3/internal/config"
	"github.com/gunsandgophers/lambda-fase-3/internal/domain"
)

func main() {
	// cognitoCliente, err := domain.NewCognitoClient(
	// 	"us-east-1",
	// 	"us-east-1_3ofqHwfxr",
	// )
	// cognitoCliente.ListUser("a4384418-a0b1-70cf-347c-34987e4a5a40")
	region := config.GetEnv("TC_AWS_REGION", "")
	userPoolId := config.GetEnv("TC_AWS_COGNITO_USER_POOL_ID", "")
	customerService, err := domain.NewAwsCustomerService(
		region,
		userPoolId,
	)
	if err != nil {
		panic(err)
	}
	// getCustomer := domain.NewGetCustomerUC(customerService)
	// customer, err := getCustomer.Execute("680.392.640-06")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(customer.Id)
	// fmt.Println(customer.Cpf)
	// fmt.Println(customer.Name)
	// fmt.Println(customer.Email)

	createCustomer := domain.NewCreateCustomerUC(customerService)
	input := &domain.CreateCustomerInput{
		Name: "Joao",
		Email: "joao@email.com",
		Cpf: "854.151.090-56",
	}
	customerCreated, err := createCustomer.Execute(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(customerCreated.Id)
	fmt.Println(customerCreated.Cpf)
	fmt.Println(customerCreated.Name)
	fmt.Println(customerCreated.Email)
}
