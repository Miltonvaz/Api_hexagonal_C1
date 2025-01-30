package application

import (
	"ejercicio1/src/cars/domain"
	"ejercicio1/src/cars/domain/entities"
)

type ListCar struct {
	db domain.ICar
}

func NewListCar(db domain.ICar) *ListCar{
	return &ListCar{db: db}
}

func (lc ListCar)Execute()([]entities.Car, error){
	return lc.db.GetAll()
}