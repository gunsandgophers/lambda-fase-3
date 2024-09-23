package domain

type CustomerOutput struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Cpf   string `json:"cpf,omitempty"`
}

type CreateCustomerInput struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Cpf   string `json:"cpf,omitempty"`
}

type GetCustomerUC struct {
	customerService CustomerService
}

type CreateCustomerUC struct {
	customerService CustomerService
}

func NewGetCustomerUC(customerService CustomerService) *GetCustomerUC {
	return &GetCustomerUC{
		customerService: customerService,
	}
}

func (uc *GetCustomerUC) Execute(cpf string) (*CustomerOutput, error) {
	validCpf, err := NewCPF(cpf)
	if err != nil {
		return nil, err
	}
	customer, err := uc.customerService.GetCustomerByCPF(validCpf)
	if err != nil {
		return nil, err
	}
	return &CustomerOutput{
		Id: customer.GetId(),
		Name: customer.GetName(),
		Email: customer.GetEmail().Value(),
		Cpf: customer.GetCPF().Value(),
	}, nil
}

func NewCreateCustomerUC(customerService CustomerService) *CreateCustomerUC {
	return &CreateCustomerUC{
		customerService: customerService,
	}
}

func (uc *CreateCustomerUC) Execute(input *CreateCustomerInput) (*CustomerOutput, error) {
	customer, err := CreateCustomer(input.Name, input.Email, input.Cpf)
	if err != nil {
		return nil, err
	}
	customerCreated, err := uc.customerService.CreateCustomer(customer)
	if err != nil {
		return nil, err
	}
	return &CustomerOutput{
		Id: customerCreated.GetId(),
		Name: customerCreated.GetName(),
		Email: customerCreated.GetEmail().Value(),
		Cpf: customerCreated.GetCPF().Value(),
	}, nil
}
