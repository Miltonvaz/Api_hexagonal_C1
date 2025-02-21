package application

import (
	"ejercicio1/src/clients/domain"
	"ejercicio1/src/clients/domain/entities"
)

type GetByLastName struct {
	db domain.IClient
}

func NewGetByLastName(db domain.IClient) *GetByLastName {
	return &GetByLastName{db: db}
}

func (vc GetByLastName) Execute(lastName string) ([]entities.Client, error) {
	return vc.db.GetLastName(lastName)
}
