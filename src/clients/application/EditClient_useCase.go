package application

import (
	"ejercicio1/src/clients/domain"
	"ejercicio1/src/clients/domain/entities"
)

type EditClient struct {
	db domain.IClient
}

func NewEditClient(db domain.IClient) *EditClient{
	return &EditClient{db : db}
}

func (ec *EditClient)Execute(client entities.Client)error{
	return ec.db.Edit(client)
}