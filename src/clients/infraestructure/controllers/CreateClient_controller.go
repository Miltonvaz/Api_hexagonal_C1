package controllers

import (
	"ejercicio1/src/clients/application"
	"ejercicio1/src/clients/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateClientController struct {
	useCase application.CreateClient
}

func NewCreateClientController(useCase application.CreateClient) *CreateClientController {
	return &CreateClientController{useCase: useCase}
}

func (cc_c *CreateClientController) Execute(c *gin.Context) {
	var input struct {
		Name      string `json:"name"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		Cellphone string `json:"cellphone"`
		Age       int32  `json:"age"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	client := entities.Client{
		Name:      input.Name,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
		Cellphone: input.Cellphone,
		Age:       input.Age,
	}

	err := cc_c.useCase.Execute(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Client": client})
}
