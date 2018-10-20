//Author Mohammad Naser Abbasanadi
//Creating Date 2018-10-20
// order.go is to handle functionality of request from outside
// it has duties to connect to models and ask for data maipulation

package controller

import (
	"GolangOrdering/helpers"
	"GolangOrdering/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//getOrders to retrieve all of orders and show to user
//a sample of request is /orders?page=1&limit=5
func (a *App) getOrders(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "INVALID_PAGE_NUMBER")
		return
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "INVALID_LIMIT_NUMBER")
		return
	}

	Orders, err := models.GetOrders(a.DB, (page-1)*limit, page*limit)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	helpers.RespondWithJSON(w, http.StatusOK, Orders)
}

//createOrder is provided in order to add orders
//a sample of request  /order
// "origin": ["START_LATITUDE", "START_LONGTITUDE"],
// "destination": ["END_LATITUDE", "END_LONGTITUDE"]
// it must use google api to find out distance between two points
// provided point must be well formated other wise it won't show right answer
func (a *App) createOrder(w http.ResponseWriter, r *http.Request) {
	origin := r.FormValue("origin")
	if len(origin) == 0 {
		helpers.RespondWithError(w, http.StatusBadRequest, "INVALID_ORIGIN")
		return
	}

	destination := r.FormValue("destination")
	if len(destination) == 0 {
		helpers.RespondWithError(w, http.StatusBadRequest, "INVALID_DESTINATION")
		return
	}

	distance, err := helpers.CalcDistance(origin, destination)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if distance == 0 {
		helpers.RespondWithError(w, http.StatusBadRequest, "PROVIDED_POINT_NOT_CORRECT")
		return
	}

	p := models.Order{
		Distance: distance,
		Status:   "UNASSIGN",
	}

	orderid, err := p.CreateOrder(a.DB)

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	p.ID = orderid
	helpers.RespondWithJSON(w, http.StatusOK, p)
}

//updateOrder is responsable to change the status of order
//a sample of request is /order/:id
// by using the id it will update the corresponding record
func (a *App) updateOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "INVALID_ORDER_ID")
		return
	}
	status := r.FormValue("status")
	if status != "taken" {
		helpers.RespondWithError(w, http.StatusBadRequest, "INVALID_STATUS")
		return
	}

	p := models.Order{
		ID:     id,
		Status: status,
	}
	effected, err := p.UpdateOrder(a.DB)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if effected == 0 {
		helpers.RespondWithError(w, http.StatusConflict, "ORDER_ALREADY_BEEN_TAKEN")
		return
	}
	helpers.RespondWithJSON(w, http.StatusOK, nil)
}
