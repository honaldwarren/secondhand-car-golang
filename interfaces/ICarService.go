package interfaces

import "secondhand-car-golang/models/dto"

type ICarService interface {
	GetListByFilter(filter *dto.CarFilterDto) (interface{}, error)
}
