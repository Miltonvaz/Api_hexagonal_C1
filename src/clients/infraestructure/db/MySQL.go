package db

import (
	"database/sql"
	"ejercicio1/src/clients/domain/entities"
	"errors"
	"fmt"
)

type MySQL struct {
	conn *sql.DB
}

func NewMySQL(conn *sql.DB) *MySQL {
	return &MySQL{conn: conn}
}

func (m *MySQL) Save(client entities.Client) error {
	query := `INSERT INTO clients (name, last_name, email, cellphone, age, password) 
              VALUES (?, ?, ?, ?, ?, ?)`
	_, err := m.conn.Exec(query, client.Name, client.LastName, client.Email, client.Cellphone, client.Age, client.Password)
	if err != nil {
		return fmt.Errorf("failed to save client: %v", err)
	}
	return nil
}

func (m *MySQL) GetByEmail(email string) (entities.Client, error) {
	var client entities.Client
	query := `SELECT id, name, last_name, email, password, cellphone, age FROM clients WHERE email = ? LIMIT 1`
	err := m.conn.QueryRow(query, email).Scan(
		&client.ID, &client.Name, &client.LastName, &client.Email, &client.Password,
		&client.Cellphone, &client.Age,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.Client{}, errors.New("usuario no encontrado")
		}
		return entities.Client{}, err
	}

	return client, nil
}

func (m *MySQL) GetAll() ([]entities.Client, error) {
	query := "SELECT * FROM clients"
	rows, err := m.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve clients: %v", err)
	}
	defer rows.Close()

	var clients []entities.Client
	for rows.Next() {
		var client entities.Client
		err := rows.Scan(&client.ID, &client.Name, &client.LastName, &client.Email, &client.Cellphone, &client.Age, &client.Password)
		if err != nil {
			return nil, fmt.Errorf("failed to scan client row: %v", err)
		}
		clients = append(clients, client)
	}

	return clients, nil
}

func (m *MySQL) GetLastName(lastName string) ([]entities.Client, error) {
	query := "SELECT * FROM clinet WHERE last_name = ?"
	rows, err := m.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve clients: %v", err)
	}
	defer rows.Close()

	var clients []entities.Client
	for rows.Next() {
		var client entities.Client
		err := rows.Scan(&client.ID, &client.Name, &client.LastName, &client.Email, &client.Cellphone, &client.Age, &client.Password)
		if err != nil {
			return nil, fmt.Errorf("failed to scan car row: %v", err)
		}
		clients = append(clients, client)
	}

	if len(clients) == 0 {
		return nil, fmt.Errorf("no cars found with fuel type: %s", lastName)
	}

	return clients, nil
}

func (m *MySQL) GetById(id int) (entities.Client, error) {
	query := "SELECT * FROM clients WHERE id = ?"
	row := m.conn.QueryRow(query, id)

	var client entities.Client
	err := row.Scan(&client.ID, &client.Name, &client.LastName, &client.Email, &client.Cellphone, &client.Age, &client.Password)
	if err == sql.ErrNoRows {
		return entities.Client{}, errors.New("client not found")
	} else if err != nil {
		return entities.Client{}, fmt.Errorf("failed to retrieve client: %v", err)
	}

	return client, nil
}

func (m *MySQL) Edit(client entities.Client) error {
	query := "UPDATE clients SET name = ?, last_name = ?, email = ?, password = ?, cellphone = ?, age = ? WHERE id = ?"
	_, err := m.conn.Exec(query, client.Name, client.LastName, client.Email, client.Password, client.Cellphone, client.Age, client.ID)

	if err != nil {
		return fmt.Errorf("failed to update client: %v", err)
	}
	return nil
}

func (m *MySQL) Delete(id int) error {
	query := "DELETE FROM clients WHERE id = ?"
	_, err := m.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete client: %v", err)
	}
	return nil
}
