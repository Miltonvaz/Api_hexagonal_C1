package controllers

import (
	"ejercicio1/src/cars/application"
	"ejercicio1/src/cars/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateCarController struct {
	usecase application.CreateCar
}

func NewCreateCarController(usecase application.CreateCar) *CreateCarController{
	return &CreateCarController{usecase: usecase}
}

func (cc_c *CreateCarController)Execute(c *gin.Context){
	var input struct{        
    	Make          string  `json:"make"`       
    	Model         string  `json:"model"`       
    	Year          int32   `json:"year"`       
    	Mileage       int32   `json:"mileage"`    
    	FuelType      string  `json:"fuel_type"` 
	}
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	car := entities.Car{
		Make:     input.Make,
		Model:    input.Model,
		Year: 	  input.Year,
		Mileage:  input.Mileage,
		FuelType: input.FuelType,
	}

	
	createdCar, err := cc_c.usecase.Execute(car)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"car": createdCar})
}