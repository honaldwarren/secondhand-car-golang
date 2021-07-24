/*
package database

import (
	"encoding/json"

	"secondhand-car-golang/app/models"

	"github.com/thedevsaddam/gojsonq"
)

func GetCars() []*models.Car {
	var cars []*models.Car
	jq := gojsonq.New().File("./listings.json")
	//res := jq.From("vendor.items").Where("price", ">", 1200).OrWhere("id", "=", nil).Only("name", "price")
	res := jq.From("cars")
	json.Unmarshal([]byte(res.String()), cars)
	return cars
}
*/