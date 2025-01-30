package application

import "ejercicio1/src/cars/domain"

type DeleteCar struct {
	db domain.ICar
}

func NewDeleteCar(db domain.ICar) *DeleteCar{
	return &DeleteCar{db: db}
}

func (dc DeleteCar)Execute(id int)error{
	return dc.db.Delete(id)
}