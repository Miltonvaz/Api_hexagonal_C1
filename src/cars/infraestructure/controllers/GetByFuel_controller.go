package controllers

import (
	"ejercicio1/src/cars/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetBYfuelController struct {
	usecase application.GetBYfuel
}

func NewGetBYfuelController(usecase application.GetBYfuel) *GetBYfuelController {
	return &GetBYfuelController{usecase: usecase}
}

func (vc_c *GetBYfuelController) Execute(c *gin.Context) {
	fuel := c.Param("fuel")

	cars, err := vc_c.usecase.Execute(fuel)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No cars found for this fuel type"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cars": cars})
}
