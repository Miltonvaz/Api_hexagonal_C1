package application

import (
	"ejercicio1/src/clients/domain"
	"ejercicio1/src/clients/domain/entities"
	"errors"
)

type ViewByIdClient struct {
	db domain.IClient
}

func NewClientById(db domain.IClient) *ViewByIdClient{
	return &ViewByIdClient{db: db}
}

func (vc *ViewByIdClient)Execute(id int)(entities.Client,error){
	client, err := vc.db.GetById(id)
	if err != nil {
		return entities.Client{}, errors.New("client not found")
	}
	return client, nil
}