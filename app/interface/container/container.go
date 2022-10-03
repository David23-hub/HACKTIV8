package container

import (
	"swagger/app/adapter/database"
	"swagger/app/repository/person"
	"swagger/app/usecase"
)

type Container struct {
	Usecase usecase.Usecase
}

func SetUpContainer() *Container {

	db := database.SetupDatabase()

	personRepo := person.NewPersonRepository(db)
	usecase := usecase.NewUsecase(personRepo)

	return &Container{
		Usecase: usecase,
	}
}
