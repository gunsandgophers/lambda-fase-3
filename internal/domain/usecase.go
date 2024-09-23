package domain

type CustomerOutput struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Cpf   string `json:"cpf,omitempty"`
}

type GetCustomerIdUC struct {
	customerService CustomerService
}

func NewGetCustomerIdUC(customerService CustomerService) *GetCustomerIdUC {
	return &GetCustomerIdUC{
		customerService: customerService,
	}
}

func (uc *GetCustomerIdUC) Execute(cpf string) (*CustomerOutput, error) {
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
