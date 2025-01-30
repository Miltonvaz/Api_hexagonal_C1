package application

import (
    "ejercicio1/src/clients/domain"
    "ejercicio1/src/clients/domain/entities"
    "ejercicio1/src/core/security"
)

type CreateClient struct {
    db domain.IClient
}

func NewCreateClient(db domain.IClient) *CreateClient {
    return &CreateClient{db: db}
}

func (cc *CreateClient) Execute(client entities.Client) error {
    hashedPassword, err := security.HashPassword(client.Password)
    if err != nil {
        return err
    }
    client.Password = hashedPassword
    return cc.db.Save(client)
}
