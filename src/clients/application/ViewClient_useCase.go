package application

import (
	"ejercicio1/src/clients/domain"
	"ejercicio1/src/clients/domain/entities"
)

type ViewClient struct {
	db domain.IClient
}

func NewListClient(db domain.IClient) *ViewClient{
	return &ViewClient{db : db}
}

func (vc *ViewClient)Execute()([]entities.Client, error){
	return vc.db.GetAll()
}


