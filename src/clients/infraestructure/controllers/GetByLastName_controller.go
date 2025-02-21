package controllers

import (
	"ejercicio1/src/clients/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetByLastNameController struct {
	usecase application.GetByLastName
}

func NewGetByLastNameController(usecase application.GetByLastName) *GetByLastNameController {
	return &GetByLastNameController{usecase: usecase}
}

func (vc_c *GetByLastNameController) Execute(c *gin.Context) {
	fuel := c.Param("lastname")

	clients, err := vc_c.usecase.Execute(fuel)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No cars found for this fuel type"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Clients": clients})
}
