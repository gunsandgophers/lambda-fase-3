package domain


type CustomerService interface {
	GetCustomerByCPF(cpf *CPF) (*Customer, error)
	CreateCustomer(customer *Customer) (*Customer, error)
}

type AwsCustomerService struct {
	serviceClient *CognitoClient
}

func NewAwsCustomerService(
	region string,
	userPoolId string,
) (*AwsCustomerService, error) {
	cognitoClient, err := NewCognitoClient(region, userPoolId)
	if err != nil {
		return nil, err
	}
	return &AwsCustomerService{
		serviceClient: cognitoClient,
	}, nil
}

func (a *AwsCustomerService) GetCustomerByCPF(cpf *CPF) (*Customer, error) {
	user, err := a.serviceClient.GetUser(cpf.Value())
	if err != nil {
		return nil, err
	}
	return RestoreCustomer(user.Id, user.Name, user.Email, user.Username)
}

func (a *AwsCustomerService) CreateCustomer(customer *Customer) (*Customer, error) {
	cognitoCreateUser := &CognitoCreateUser{
		Username: customer.GetCPF().Value(),
		Name: customer.GetName(),
		Email: customer.GetEmail().Value(),
	}
	user, err := a.serviceClient.CreateUser(cognitoCreateUser)
	if err != nil {
		return nil, err
	}
	return RestoreCustomer(user.Id, user.Name, user.Email, user.Username)
}
