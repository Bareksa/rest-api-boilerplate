package services

import (
	"github.com/Bareksa/rest-api-boilerplate/models"
	"github.com/Bareksa/rest-api-boilerplate/repositories"
)

type IExampleService interface {
	GetUser() [] models.User
}

type ExampleService struct {
	Repository IExampleService
}

func InitExampleService() *ExampleService{
	repository := new(repositories.ExampleRepository)

	return &ExampleService{
		Repository: repository,
	}
}

func (s *ExampleService) GetUser() (users [] models.User) {
	users = s.Repository.GetUser()
	return
}
