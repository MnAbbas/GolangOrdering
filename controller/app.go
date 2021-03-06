//Author Mohammad Naser Abbasanadi
//Creating Date 2018-10-20
// app.go is for handelling behavior of whole application , initialize and defining routers

package controller

import (
	"GolangOrdering/helpers"
	"GolangOrdering/logger"
	"GolangOrdering/models"

	"database/sql"
	"net/http"

	//mysql package imported
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//App structure to handle the functionality of whole application
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

//Initialize application by calling db and start defining router
func (a *App) Initialize(user, password, server, port, dbname string) {

	var err error
	a.DB, err = models.GetDBInstance(user, password, server, port, dbname)
	if err != nil {
		logger.Log.Panic(err)
	}
	logger.Log.Println("DB is started to act")

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

//Run is responsable to run application based on provided addres and router
func (a *App) Run(addr string) {
	logger.Log.Fatal(http.ListenAndServe(addr, a.Router))
}

//initializeRoutes is defining routers
func (a *App) initializeRoutes() {

	a.Router.HandleFunc("/orders", middlewares(a.getOrders)).Methods("GET")
	a.Router.HandleFunc("/order", middlewares(a.createOrder)).Methods("POST")
	a.Router.HandleFunc("/order/{id:[0-9]+}", middlewares(a.updateOrder)).Methods("PUT")
	a.Router.NotFoundHandler = middlewares(http.HandlerFunc(notFound))
}

//notFound to customize not found message in gorilla
func notFound(w http.ResponseWriter, r *http.Request) {
	helpers.RespondWithError(w, http.StatusNotFound, "ROUTE_NOT_FOUND")
}
