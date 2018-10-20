package test

import (
	"GolangOrdering/controller"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gorilla/mux"
)

type Order struct {
	ID       int    `json:"id"`
	Distance int    `json:"distance"`
	Status   string `json:"status"`
}

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

//executeRequest provide endpoint to serve http request
func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	a := controller.App{}
	a.Initialize(
		"joe",
		"U_Xz$M,(3${gpTcZ",
		"185.159.153.22",
		"3306",
		"orders")
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

//Test to response notfound response ; for Endpoint : GET /orders/
func TestNotFound(t *testing.T) {
	req, _ := http.NewRequest("GET", "/users", nil)
	response := executeRequest(req)
	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)

	equals(t, http.StatusNotFound, response.Code)
	equals(t, m["error"], "ROUTE_NOT_FOUND")

}

//TestGetOrders_1 test if there is no parameter for orders it must return InternalServerError ; for Endpoint : GET /orders/
func TestGetOrders_1(t *testing.T) {
	req, _ := http.NewRequest("GET", "/orders", nil)
	response := executeRequest(req)
	// t.Log(response.Body)
	equals(t, http.StatusBadRequest, response.Code)
}

//TestGetOrders_2 is testing; the correct answer with output ; for Endpoint : GET /orders/
func TestGetOrders_2(t *testing.T) {
	page, limit := 1, 3
	req, _ := http.NewRequest("GET", fmt.Sprintf("/orders?page=%d&limit=%d", page, limit), nil)
	response := executeRequest(req)

	var m []Order
	json.Unmarshal(response.Body.Bytes(), &m)
	equals(t, http.StatusOK, response.Code)
	equals(t, len(m), limit)

}

//TestGetOrders_3 is testing : invalid page number for orders ; for Endpoint : GET /orders/
func TestGetOrders_3(t *testing.T) {
	page, limit := "1r", 3
	req, _ := http.NewRequest("GET", fmt.Sprintf("/orders?page=%v&limit=%v", page, limit), nil)
	response := executeRequest(req)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)

	equals(t, http.StatusBadRequest, response.Code)
	equals(t, m["error"], "INVALID_PAGE_NUMBER")

}

//TestGetOrders_4 is testing : invalid limit number for orders; for Endpoint : GET /orders/
func TestGetOrders_4(t *testing.T) {
	page, limit := 1, "3e"
	req, _ := http.NewRequest("GET", fmt.Sprintf("/orders?page=%v&limit=%v", page, limit), nil)
	response := executeRequest(req)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)

	equals(t, http.StatusBadRequest, response.Code)
	equals(t, m["error"], "INVALID_LIMIT_NUMBER")

}

//TestCreateOrder_1 is testing : invalid limit provided points for calculation distance for Endpoint : post /order
func TestCreateOrder_1(t *testing.T) {

	req, _ := http.NewRequest("POST", "/order", nil)
	form := url.Values{}
	form.Add("origin", `["m40.6905615","-73.9976592"]`)
	form.Add("destination", `["40.6655101","-73.89188969999998"]`)
	req.PostForm = form
	response := executeRequest(req)
	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	equals(t, http.StatusBadRequest, response.Code)
	equals(t, m["error"], "PROVIDED_POINT_NOT_CORRECT")

}

//TestCreateOrder_2 is testing : the correct answer for Endpoint : post /order
func TestCreateOrder_2(t *testing.T) {
	req, _ := http.NewRequest("POST", "/order", nil)
	form := url.Values{}
	form.Add("origin", `["40.6905615","-73.9976592"]`)
	form.Add("destination", `["40.6655101","-73.89188969999998"]`)
	req.PostForm = form
	response := executeRequest(req)
	// t.Log(response.Body)
	var m Order
	json.Unmarshal(response.Body.Bytes(), &m)
	equals(t, http.StatusOK, response.Code)
	equals(t, m.Distance, 33)
}

//TestCreateOrder_3 is testing : the wrong orging point for Endpoint : post /order
func TestCreateOrder_3(t *testing.T) {
	req, _ := http.NewRequest("POST", "/order", nil)
	form := url.Values{}
	// form.Add("origin", `["40.6905615","-73.9976592"]`)
	form.Add("destination", `["40.6655101","-73.89188969999998"]`)
	req.PostForm = form
	response := executeRequest(req)
	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	equals(t, http.StatusBadRequest, response.Code)
	equals(t, m["error"], "INVALID_ORIGIN")
}

//TestCreateOrder_5 is testing : the wrong destination point for Endpoint : post /order
func TestCreateOrder_5(t *testing.T) {
	req, _ := http.NewRequest("POST", "/order", nil)
	form := url.Values{}
	form.Add("origin", `["40.6905615","-73.9976592"]`)
	// form.Add("destination", `40.6655101`)
	req.PostForm = form
	response := executeRequest(req)
	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	equals(t, http.StatusBadRequest, response.Code)
	equals(t, m["error"], "INVALID_DESTINATION")
}

//TestCreateOrder_6 is testing : the invalid orging or destination point for Endpoint : post /order
func TestCreateOrder_6(t *testing.T) {
	req, _ := http.NewRequest("POST", "/order", nil)
	form := url.Values{}
	form.Add("origin", `["40.6905615","-73.9976592"]`)
	form.Add("destination", `40.6655101`)
	req.PostForm = form
	response := executeRequest(req)
	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	equals(t, http.StatusBadRequest, response.Code)
	equals(t, m["error"], "PROVIDED_POINT_NOT_CORRECT")
}

//TestUpdateOrder_1 is testing : the correct answer point for Endpoint , you sould use a neworderid to get writ answer : put /order/:id
// this method must be called just once the next time will not pass the error
func TestUpdateOrder_1(t *testing.T) {
	//this id must not be used
	orderid := 36
	urlput := fmt.Sprintf("/order/%d", orderid)
	req, _ := http.NewRequest("PUT", urlput, nil)
	form := url.Values{}
	form.Add("status", `taken`)
	req.PostForm = form
	response := executeRequest(req)
	// t.Log(response.Body)
	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	equals(t, m["status"], "SUCCESS")
	equals(t, http.StatusOK, response.Code)

}

//TestUpdateOrder_2 is testing : check the response for orders which already have been taken for Endpoint : put /order/:id
func TestUpdateOrder_2(t *testing.T) {
	req, _ := http.NewRequest("PUT", "/order/1", nil)
	form := url.Values{}
	form.Add("status", `taken`)
	req.PostForm = form
	response := executeRequest(req)
	// t.Log(response.Body)
	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	equals(t, m["error"], "ORDER_ALREADY_BEEN_TAKEN")
	equals(t, http.StatusConflict, response.Code)

}

//TestUpdateOrder_3 is testing : the invalid status value for status parameter for Endpoint : put /order/:id
func TestUpdateOrder_3(t *testing.T) {
	req, _ := http.NewRequest("PUT", "/order/34", nil)
	form := url.Values{}
	form.Add("status", `something`)
	req.PostForm = form
	response := executeRequest(req)
	// t.Log(response.Body)
	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	equals(t, m["error"], "INVALID_STATUS")
	equals(t, http.StatusBadRequest, response.Code)

}
