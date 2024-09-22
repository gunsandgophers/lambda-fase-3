package domain

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type CustomerServiceClient interface {
	GetUser(username string) (*CognitoUser, error)
}

type CognitoClient struct {
	client *cognito.CognitoIdentityProvider
	userPoolId string
}

type CognitoUser struct {
	Id string
	Username string
	Name string
	Email string
}

func NewCognitoClient(region string, userPoolId string) (*CognitoClient, error) {
	conf := &aws.Config{
		Region: aws.String(region),
	}
	sess, err := session.NewSession(conf)
	if err != nil {
		return nil, err
	}
	client := cognito.New(sess)
	return &CognitoClient{
		client: client,
		userPoolId: userPoolId,
	}, nil
}

func (cc *CognitoClient) GetUser(username string) (*CognitoUser, error) {
	input := &cognito.AdminGetUserInput{
		UserPoolId: aws.String(cc.userPoolId),
		Username: aws.String(username),
	}
	output, err := cc.client.AdminGetUser(input)
	if err != nil {
		return nil, err
	}

	user := &CognitoUser{
		Username: *output.Username,
	}

	for _, attr := range output.UserAttributes {
		if *attr.Name == "email" {
			user.Email = *attr.Value
		} else if *attr.Name == "sub" {
			user.Id = *attr.Value
		} else if *attr.Name == "name" {
			user.Name = *attr.Value
		}
	}
	return user, nil
}
