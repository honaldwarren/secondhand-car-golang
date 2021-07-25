package main

import (
	"sync"

	"secondhand-car-golang/controllers"
	"secondhand-car-golang/repositories"
	"secondhand-car-golang/services"
)

type IServiceContainer interface {
	InjectCarController() controllers.CarController
}

type kernel struct{}

func (k *kernel) InjectCarController() controllers.CarController {

	carRepository := &repositories.CarRepository{}
	carService := &services.CarService{carRepository}
	carController := controllers.CarController{carService}

	return carController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
