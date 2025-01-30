package controllers

import (
	"ejercicio1/src/clients/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewClientController struct {
	usecase application.ViewClient
}

func NewViewClientController(usecase application.ViewClient) *ViewClientController {
	return &ViewClientController{usecase: usecase}
}

func (vc_c *ViewClientController)Execute(c *gin.Context) {

	clients, err := vc_c.usecase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve clients"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"clients": clients, 
	})
}
