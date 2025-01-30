package controllers

import (
	"ejercicio1/src/cars/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListCarController struct {
	usecase application.ListCar
}

func NewListCarController(usecase application.ListCar) *ListCarController{
	return &ListCarController{usecase:usecase}
}

func (lc_c *ListCarController)Execute(c *gin.Context){
	car, err := lc_c.usecase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve clients"})
		return
	}
	
	if len(car) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No cars available"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cars": car})
}