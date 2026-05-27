package routes

import (
	"ejercicio1/src/cars/infraestructure/controllers"
	"ejercicio1/src/core/security"
	"github.com/gin-gonic/gin"
)

func RegisterCarRoutes(r *gin.Engine,
	createCarController *controllers.CreateCarController,
	viewCarController *controllers.ListCarController,
	viewByIdCarController *controllers.ViewByIdCarController,
	updateCarController *controllers.UpdateCarController,
	deleteCarController *controllers.DeleteCarController,
	getByFuelController *controllers.GetBYfuelController,
) {
	r.POST("/cars", security.JWTMiddleware(), createCarController.Execute)
	r.GET("/cars", security.JWTMiddleware(), viewCarController.Execute)
	r.GET("/cars/:id", security.JWTMiddleware(), viewByIdCarController.Execute)
	r.PUT("/cars/:id", security.JWTMiddleware(), updateCarController.Execute)
	r.DELETE("/cars/:id", security.JWTMiddleware(), deleteCarController.Execute)
	r.GET("/cars/fuel/:fuel", security.JWTMiddleware(), getByFuelController.Execute)
}
