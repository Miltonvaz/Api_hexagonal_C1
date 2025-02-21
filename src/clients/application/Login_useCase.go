package application

import (
	"ejercicio1/src/clients/domain"
	"ejercicio1/src/clients/domain/entities"
	"ejercicio1/src/core/security"
	"errors"
)

type AuthService struct {
	clientRepo domain.IClient
}

func NewAuthService(clientRepo domain.IClient) *AuthService {
	return &AuthService{clientRepo: clientRepo}
}

func (as *AuthService) Login(email, password string) (string, error) {
	client, err := as.clientRepo.GetByEmail(email)
	if err != nil {
		return "", errors.New("usuario no encontrado")
	}

	if !security.CheckPassword(client.Password, password) {
		return "", errors.New("contraseña incorrecta")
	}

	token, err := security.GenerateJWT(int(client.ID), client.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (as *AuthService) Register(client entities.Client) error {

	hashedPassword, err := security.HashPassword(client.Password)
	if err != nil {
		return err
	}
	client.Password = hashedPassword

	return as.clientRepo.Save(client)
}
