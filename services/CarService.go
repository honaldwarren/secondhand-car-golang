package services

import (
	"secondhand-car-golang/interfaces"
	"secondhand-car-golang/models/dto"
	"strings"
)

type CarService struct {
	interfaces.ICarRepository
}

func (service *CarService) GetListByFilter(filter *dto.CarFilterDto) (interface{}, error) {

	listJsonQ, err := service.GetList()
	if err != nil {
		return "", err
	}

	//year filter
	if filter.MinYear > 0 {
		listJsonQ = listJsonQ.Where("year", ">=", filter.MinYear)
	}

	if filter.MaxYear > 0 {
		listJsonQ = listJsonQ.Where("year", "<=", filter.MaxYear)
	}

	//price filter
	if filter.MinPrice > 0 {
		listJsonQ = listJsonQ.Where("price", ">=", filter.MinPrice)
	}

	if filter.MaxPrice > 0 {
		listJsonQ = listJsonQ.Where("price", "<=", filter.MaxPrice)
	}

	//make filter
	if strings.TrimSpace(filter.Make) != "" {
		listJsonQ = listJsonQ.Where("make", "in", strings.Split(filter.Make, ","))
	}

	//model filter
	if strings.TrimSpace(filter.Model) != "" {
		listJsonQ = listJsonQ.Where("model", "in", strings.Split(filter.Model, ","))
	}

	resp := listJsonQ.Get()
	return resp, nil
}
