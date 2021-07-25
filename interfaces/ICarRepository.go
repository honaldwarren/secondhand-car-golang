package interfaces

import "github.com/thedevsaddam/gojsonq"

type ICarRepository interface {
	GetList() (*gojsonq.JSONQ, error)
}
