package controllers

import (
	"context"
	"ejercicio1/src/cars/application"
	"ejercicio1/src/cars/domain/entities"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

type CreateCarController struct {
	usecase application.CreateCar
}

func NewCreateCarController(usecase application.CreateCar) *CreateCarController {
	return &CreateCarController{usecase: usecase}
}

func (cc_c *CreateCarController) Execute(c *gin.Context) {
	car := entities.Car{}

	contentType := c.ContentType()
	if strings.Contains(contentType, "application/json") {
		if err := c.ShouldBindJSON(&car); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
			return
		}
	} else {
		car.Make = c.PostForm("make")
		car.Model = c.PostForm("model")
		car.FuelType = c.PostForm("fuel_type")
		car.ImageURL = c.PostForm("image_url")

		if yearStr := c.PostForm("year"); yearStr != "" {
			if y, err := strconv.Atoi(yearStr); err == nil {
				car.Year = int32(y)
			}
		}
		if mileageStr := c.PostForm("mileage"); mileageStr != "" {
			if m, err := strconv.Atoi(mileageStr); err == nil {
				car.Mileage = int32(m)
			}
		}
	}

	if car.ImageURL == "" {
		fileHeader, err := c.FormFile("image")
		if err == nil && fileHeader != nil {
			file, err := fileHeader.Open()
			if err == nil {
				defer file.Close()
				cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
				if err == nil {
					uploadRes, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{Folder: "cars"})
					if err == nil {
						car.ImageURL = uploadRes.SecureURL
					}
				}
			}
		}
	}

	createdCar, err := cc_c.usecase.Execute(car)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"car": createdCar})
}
