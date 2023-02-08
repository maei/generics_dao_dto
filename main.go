package main

import (
	"context"
	"log"

	"github.com/maei/generics_dao_dto/generic_dto_dao"
)

func main() {
	user := generic_dto_dao.UserDAO{
		FirstName: "Jim",
		LastName:  "Panse",
	}

	createUserCommand := generic_dto_dao.NewCreateCommand[generic_dto_dao.UserDTO, *generic_dto_dao.UserDAO]()
	handleUser, err := createUserCommand.Handle(context.Background(), &user)
	if err != nil {
		return
	}

	log.Println(handleUser)

	company := generic_dto_dao.CompanyDAO{
		Name: "Baumhaus GmbH",
	}

	createCompanyCommand := generic_dto_dao.NewCreateCommand[generic_dto_dao.CompanyDTO, *generic_dto_dao.CompanyDAO]()
	handleCompany, err := createCompanyCommand.Handle(context.Background(), &company)
	if err != nil {
		return
	}

	log.Println(handleCompany)

	animal := generic_dto_dao.AnimalDAO{
		Name: "DOG",
	}

	createAnimalCommand := generic_dto_dao.NewCreateCommand[generic_dto_dao.AnimalDTO, *generic_dto_dao.AnimalDAO]()
	handleAnimal, err := createAnimalCommand.Handle(context.Background(), &animal)
	if err != nil {
		return
	}

	log.Println(handleAnimal)

}
