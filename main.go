package main

import (
	"ejercicio1/src/cars/infraestructure/dependencies"
	"ejercicio1/src/cars/infraestructure/routes"
	dependenciesc "ejercicio1/src/clients/infraestructure/dependencies_C"
	routesc "ejercicio1/src/clients/infraestructure/routes_C"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	createClientController, viewClientController, editClientController, deleteClientController, viewByIdClientController, authController, getlastNameController, err := dependenciesc.Init()
	if err != nil {
		log.Fatalf("Error initializing client dependencies: %v", err)
		return
	}

	createCarController, viewCarController, deleteCarController, updateCarController, viewByIdCarController, getByFuelCarController, err := dependencies.InitCars()
	if err != nil {
		log.Fatalf("Error initializing car dependencies: %v", err)
		return
	}

	r := gin.Default()

	routes.RegisterCarRoutes(r,
		createCarController,
		viewCarController,
		viewByIdCarController,
		updateCarController,
		deleteCarController,
		getByFuelCarController,
	)

	routesc.RegisterClientRoutes(r,
		createClientController,
		viewClientController,
		editClientController,
		deleteClientController,
		viewByIdClientController,
		authController,
		getlastNameController,
	)

	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
