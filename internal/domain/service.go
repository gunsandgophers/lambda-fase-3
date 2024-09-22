package domain


type CustomerService interface {
	GetCustomerByCPF(cpf *CPF) (*Customer, error)
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
