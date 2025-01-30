package controllers

import (
	"ejercicio1/src/cars/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteCarController struct {
	usecase application.DeleteCar
}


func NewDeleteCarController(usecase application.DeleteCar) *DeleteCarController {
	return &DeleteCarController{usecase: usecase}
}

func (dc_c *DeleteCarController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := dc_c.usecase.Execute(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Car deleted successfully"})
}

