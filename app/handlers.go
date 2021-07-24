package app

import (
	"net/http"

	"github.com/thedevsaddam/gojsonq"
)

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jq := gojsonq.New().File("./listings.json")
		//res := jq.From("vendor.items").Where("price", ">", 1200).OrWhere("id", "=", nil).Only("name", "price")
		res := jq.From("cars").Where("year", ">", 2012).Get()
		sendResponse(w, res, http.StatusOK)
	}
}
