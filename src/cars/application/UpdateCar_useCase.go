package application

import (
	"ejercicio1/src/cars/domain"
	"ejercicio1/src/cars/domain/entities"
)

type UpdateCar struct {
	db domain.ICar
}

func NewUpdateCar(db domain.ICar) * UpdateCar{
	return &UpdateCar{db: db}
}

func (uc UpdateCar)Execute( car entities.Car)error{
	return uc.db.Edit(car)}