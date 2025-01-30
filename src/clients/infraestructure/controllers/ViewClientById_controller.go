package controllers

import (
	"ejercicio1/src/clients/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ViewClientByIdController struct {
	usecase application.ViewByIdClient
}

func NewViewClientByIdController(usecase application.ViewByIdClient) *ViewClientByIdController {
	return &ViewClientByIdController{usecase: usecase}
}

func (vc_c *ViewClientByIdController) Execute(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	client, err := vc_c.usecase.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"client": client,
	})
}
