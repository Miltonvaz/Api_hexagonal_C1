package application

import "ejercicio1/src/clients/domain"

type DeleteClient struct {
	db domain.IClient
}

func NewDeleteClient(db domain.IClient) *DeleteClient{
	return &DeleteClient{db : db }
}

func (dc DeleteClient) Execute(id int ) error{
	return dc.db.Delete(id)
}