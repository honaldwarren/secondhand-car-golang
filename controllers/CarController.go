package controllers

import (
	"net/http"
	"secondhand-car-golang/helpers"
	"secondhand-car-golang/interfaces"
	"secondhand-car-golang/models/dto"

	"github.com/gorilla/schema"
)

type CarController struct {
	interfaces.ICarService
}

func (controller *CarController) Filter(res http.ResponseWriter, req *http.Request) {

	if err := req.ParseForm(); err != nil {
		helpers.SendResponse(res, nil, http.StatusInternalServerError, err)
		return
	}

	filter := new(dto.CarFilterDto)
	if err := schema.NewDecoder().Decode(filter, req.Form); err != nil {
		helpers.SendResponse(res, nil, http.StatusInternalServerError, err)
		return
	}

	list, err := controller.GetListByFilter(filter)
	if err != nil {
		helpers.SendResponse(res, nil, http.StatusInternalServerError, err)
		return
	}

	helpers.SendResponse(res, list, http.StatusOK, nil)

}
