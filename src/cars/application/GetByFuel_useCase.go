package application

import (
	"ejercicio1/src/cars/domain"
	"ejercicio1/src/cars/domain/entities"
)

type GetBYfuel struct {
	db domain.ICar
}

func NewGetBYfuel(db domain.ICar) *GetBYfuel {
	return &GetBYfuel{db: db}
}

func (vc GetBYfuel) Execute(fuel string) ([]entities.Car, error) {
	return vc.db.GetByFuel(fuel)
}
