package generic_dto_dao

import "fmt"

// User

type UserDAO struct {
	FirstName string
	LastName  string
}

func (dao *UserDAO) CreateDTO() UserDTO {
	return UserDTO{
		FirstName: dao.FirstName,
		LastName:  dao.LastName,
		FullName:  fmt.Sprintf("%s %s", dao.FirstName, dao.LastName),
	}
}

type UserDTO struct {
	FirstName string `bson:"firstName"`
	LastName  string `bson:"lastName"`
	FullName  string `bson:"fullName"`
}

// Company

type CompanyDAO struct {
	Name string
}

func (dao *CompanyDAO) CreateDTO() CompanyDTO {
	return CompanyDTO{
		Name: dao.Name,
	}
}

type CompanyDTO struct {
	Name string `bson:"name"`
}

// Animal

type AnimalDAO struct {
	Name string
}

func (dao *AnimalDAO) CreateDTO() AnimalDTO {
	return AnimalDTO{
		Name: dao.Name,
	}
}

type AnimalDTO struct {
	Name string `bson:"name"`
}
