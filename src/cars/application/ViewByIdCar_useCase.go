package application

import (
	"ejercicio1/src/cars/domain"
	"ejercicio1/src/cars/domain/entities"
)

type ViewByIdCar struct {
	db domain.ICar
}

func NewViewByIdCar(db  domain.ICar) *ViewByIdCar{
	return &ViewByIdCar{db:db}
}

func (vc ViewByIdCar)Execute(id int )(entities.Car, error){
	return vc.db.GetById(id)


}