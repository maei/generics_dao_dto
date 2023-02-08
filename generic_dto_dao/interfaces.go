package generic_dto_dao

type DTO interface {
	UserDTO | CompanyDTO | AnimalDTO
}

type WriteRepo[T any] interface {
	InsertOne(data T) error
}

type CreateCommandI[T any] interface {
	CreateDTO() T
}
