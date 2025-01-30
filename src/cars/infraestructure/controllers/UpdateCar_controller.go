package controllers

import (
	"ejercicio1/src/cars/application"
	"ejercicio1/src/cars/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateCarController struct {
	usecase application.UpdateCar
}

func NewUpdateCarController(usecase application.UpdateCar) *UpdateCarController{
	return &UpdateCarController{usecase: usecase}
}

func (uc_c *UpdateCarController)Execute(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
	var request struct{
		ID            int32   `json:"id"`          
    	Make          string  `json:"make"`       
    	Model         string  `json:"model"`       
    	Year          int32   `json:"year"`       
    	Mileage       int32   `json:"mileage"`    
    	FuelType      string  `json:"fuel_type"` 
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	car := entities.Car{
		ID:        int32(id), 
		Make:      request.Make,
		Model:     request.Model,
		Year:      request.Year,
		Mileage:   request.Mileage,
		FuelType:  request.FuelType,
	}
	
	err = uc_c.usecase.Execute(car)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }


	c.JSON(http.StatusCreated, gin.H{"car":car})

}