package generic_dto_dao

import (
	"context"
	"log"
)

type CreateCommand[D DTO, T CreateCommandI[D]] struct {
	writeRepo WriteRepo[D]
}

func NewCreateCommand[D DTO, T CreateCommandI[D]]() CreateCommand[D, T] {
	return CreateCommand[D, T]{}
}

func (cc *CreateCommand[D, T]) Handle(ctx context.Context, dao T) (string, error) {
	// I Want To create the UserDTO out of DAO
	dto := dao.CreateDTO()

	log.Println(dto)

	//err := cc.writeRepo.InsertOne(dto)
	//if err != nil {
	//	return "", err
	//}

	return "1", nil
}
