package dependencies

import (
	"ejercicio1/src/cars/application"
	"ejercicio1/src/cars/infraestructure/controllers"
	"ejercicio1/src/core"
	"ejercicio1/src/cars/infraestructure/db"
)

func InitCars() (
	*controllers.CreateCarController,
	*controllers.ListCarController,
	*controllers.DeleteCarController,
	*controllers.UpdateCarController,
	*controllers.ViewByIdCarController,
	error,
) {

	pool := core.GetDBPool()
	ps := db.NewMySQL(pool.DB)

	createCar := application.NewCreateCar(ps)
	listCar := application.NewListCar(ps)
	updateCar := application.NewUpdateCar(ps)
	deleteCar := application.NewDeleteCar(ps)
	viewByIdCar := application.NewViewByIdCar(ps)

	createCarController := controllers.NewCreateCarController(*createCar)
	viewCarController := controllers.NewListCarController(*listCar)
	updateCarController := controllers.NewUpdateCarController(*updateCar)
	deleteCarController := controllers.NewDeleteCarController(*deleteCar)
	viewByIdCarController := controllers.NewViewByIdCarController(*viewByIdCar)

	return 	createCarController, viewCarController,deleteCarController, updateCarController, viewByIdCarController, nil
}
