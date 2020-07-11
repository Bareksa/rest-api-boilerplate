package repositories

import "github.com/Bareksa/rest-api-boilerplate/models"

type IExampleRepository interface {
	GetUser() [] models.User
}

type ExampleRepository struct {}

func (r *ExampleRepository) GetUser() (users []models.User) {
	users = []models.User{
		models.User{
			Username:  "robert",
			FirstName: "Robert",
			Age:       40,
		},
		models.User{
			Username:  "griesemer",
			FirstName: "griesemer",
			Age:       40,
		},
	}
	return
}