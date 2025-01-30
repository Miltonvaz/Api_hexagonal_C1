package controllers

import (
	"ejercicio1/src/cars/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ViewByIdCarController struct {
	usecase application.ViewByIdCar
}


func NewViewByIdCarController(usecase application.ViewByIdCar) *ViewByIdCarController {
	return &ViewByIdCarController{usecase: usecase}
}


func (vc_c *ViewByIdCarController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	car, err := vc_c.usecase.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"car": car,
	})
}
