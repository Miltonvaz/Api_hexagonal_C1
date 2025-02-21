package db

import (
	"database/sql"
	"ejercicio1/src/cars/domain/entities"
	"fmt"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) *MySQL {
	return &MySQL{conn: conn}
}

func (m *MySQL) Save(car entities.Car) (entities.Car, error) {
	query := "INSERT INTO cars (make, model, year, mileage, fuel_type) VALUES (?, ?, ?, ?, ?)"
	result, err := m.conn.Exec(query, car.Make, car.Model, car.Year, car.Mileage, car.FuelType)
	if err != nil {
		return entities.Car{}, fmt.Errorf("failed to save car: %v", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return entities.Car{}, fmt.Errorf("failed to get last insert ID: %v", err)
	}
	car.ID = int32(lastInsertID)

	return car, nil
}

func (m *MySQL) GetAll() ([]entities.Car, error) {
	query := "SELECT * FROM cars"
	rows, err := m.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve cars: %v", err)
	}
	defer rows.Close()

	var cars []entities.Car
	for rows.Next() {
		var car entities.Car
		err := rows.Scan(&car.ID, &car.Make, &car.Model, &car.Year, &car.Mileage, &car.FuelType)
		if err != nil {
			return nil, fmt.Errorf("failed to scan car row: %v", err)
		}
		cars = append(cars, car)
	}

	return cars, nil
}
func (m *MySQL) GetByFuel(fuel string) ([]entities.Car, error) {
	query := "SELECT * FROM cars WHERE fuel_type = ?"
	rows, err := m.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve cars: %v", err)
	}
	defer rows.Close()

	var cars []entities.Car
	for rows.Next() {
		var car entities.Car
		err := rows.Scan(&car.ID, &car.Make, &car.Model, &car.Year, &car.Mileage, &car.FuelType)
		if err != nil {
			return nil, fmt.Errorf("failed to scan car row: %v", err)
		}
		cars = append(cars, car)
	}

	if len(cars) == 0 {
		return nil, fmt.Errorf("no cars found with fuel type: %s", fuel)
	}

	return cars, nil
}

func (m *MySQL) GetById(id int) (entities.Car, error) {
	query := "SELECT * FROM cars WHERE id = ?"
	row := m.conn.QueryRow(query, id)

	var car entities.Car
	err := row.Scan(&car.ID, &car.Make, &car.Model, &car.Year, &car.Mileage, &car.FuelType)
	if err == sql.ErrNoRows {
		return entities.Car{}, fmt.Errorf("car not found")
	} else if err != nil {
		return entities.Car{}, fmt.Errorf("failed to retrieve car: %v", err)
	}

	return car, nil
}

func (m *MySQL) Edit(car entities.Car) error {
	query := "UPDATE cars SET make = ?, model = ?, year = ?, mileage = ?, fuel_type = ? WHERE id = ?"
	_, err := m.conn.Exec(query, car.Make, car.Model, car.Year, car.Mileage, car.FuelType, car.ID)
	if err != nil {
		return fmt.Errorf("failed to update car: %v", err)
	}
	return nil
}

func (m *MySQL) Delete(id int) error {
	query := "DELETE FROM cars WHERE id = ?"
	_, err := m.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete car: %v", err)
	}
	return nil
}
