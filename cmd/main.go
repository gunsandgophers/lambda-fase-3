package main

import (
	"fmt"
	"github.com/gunsandgophers/lambda-fase-3/internal/domain"
)

func main() {
	customerService, err := domain.NewAwsCustomerService(
		"us-east-1",
		"us-east-1_3ofqHwfxr",
	)
	if err != nil {
		panic(err)
	}
	getCustomer := domain.NewGetCustomerIdUC(customerService)
	customer, err := getCustomer.Execute("680.392.640-06")
	if err != nil {
		panic(err)
	}
	fmt.Println(customer.Id)
	fmt.Println(customer.Cpf)
	fmt.Println(customer.Name)
	fmt.Println(customer.Email)
}
