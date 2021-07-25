package repositories

import (
	"github.com/thedevsaddam/gojsonq"
)

type CarRepository struct {
}

func (repository CarRepository) GetList() (*gojsonq.JSONQ, error) {
	jq := gojsonq.New().File("./listings.json")

	//TODO: apply error handling if connecting to database
	return jq.From("cars"), nil
}
