//Author Mohammad Naser Abbasanadi
//Creating Date 2018-10-20
// main.go

package main

import (
	"GolangOrdering/config"
	"GolangOrdering/controller"
)

func main() {
	cnf := config.GetConfigInstance()

	a := controller.App{}
	a.Initialize(
		cnf.DB.User,
		cnf.DB.Password,
		cnf.DB.Address,
		cnf.DB.Port,
		cnf.DB.Dbname)
	a.Run(cnf.SERVER.Address)
}
