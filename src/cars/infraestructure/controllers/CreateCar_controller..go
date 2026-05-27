package controllers

import (
	"context"
	"ejercicio1/src/cars/application"
	"ejercicio1/src/cars/domain/entities"
	"net/http"
	"os"
	"strconv"

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
	// Expect multipart/form-data with optional `image` file and form fields
	make := c.PostForm("make")
	model := c.PostForm("model")
	yearStr := c.PostForm("year")
	mileageStr := c.PostForm("mileage")
	fuelType := c.PostForm("fuel_type")

	var year int32
	if y, err := strconv.Atoi(yearStr); err == nil {
		year = int32(y)
	}
	var mileage int32
	if m, err := strconv.Atoi(mileageStr); err == nil {
		mileage = int32(m)
	}

	imageURL := ""
	fileHeader, err := c.FormFile("image")
	if err == nil && fileHeader != nil {
		file, err := fileHeader.Open()
		if err == nil {
			defer file.Close()
			cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
			if err == nil {
				uploadRes, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{Folder: "cars"})
				if err == nil {
					imageURL = uploadRes.SecureURL
				}
			}
		}
	}

	car := entities.Car{
		Make:     make,
		Model:    model,
		Year:     year,
		Mileage:  mileage,
		FuelType: fuelType,
		ImageURL: imageURL,
	}

	createdCar, err := cc_c.usecase.Execute(car)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"car": createdCar})
}
