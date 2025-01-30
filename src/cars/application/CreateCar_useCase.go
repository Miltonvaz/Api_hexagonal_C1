package application

import (
	"ejercicio1/src/cars/domain"
	"ejercicio1/src/cars/domain/entities"
)

type CreateCar struct {
	db domain.ICar
}

func NewCreateCar(db domain.ICar) *CreateCar{
	return &CreateCar{db: db}
}

func (cc *CreateCar)Execute(car entities.Car)(entities.Car, error){
	return cc.db.Save(car)
}