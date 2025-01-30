package controllers

import (
	"ejercicio1/src/clients/application"
	"ejercicio1/src/clients/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EditClientController struct {
	usecase application.EditClient
}

func NewEditClientController(usecase application.EditClient) *EditClientController{
	return &EditClientController{usecase: usecase}
}

func (ed_c *EditClientController)Execute(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

	var request struct {
		Name     string `json:"name"`
		LastName string `json:"lastName"`
		Email    string `json:"email"`
		Password  string `json:"password"`
		Cellphone string `json:"cellphone"`
		Age      int32  `json:"age"`
	}

    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }
	
	client := entities.Client{
		ID:        int32(id),   
		Name:      request.Name, 
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
		Cellphone: request.Cellphone,
		Age:       request.Age,
	}
    if err := ed_c.usecase.Execute(client); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{ "Cliente updated successfully": client})
}
