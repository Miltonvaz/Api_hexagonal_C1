package domain

import "ejercicio1/src/clients/domain/entities"

type IClient interface {
    Save(entities.Client) error
    GetAll() ([]entities.Client, error)
    GetById(id int) (entities.Client, error)
    GetByEmail(email string) (entities.Client, error)
    Edit(entities.Client) error
    Delete(id int) error
}
