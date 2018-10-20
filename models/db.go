//Author Mohammad Naser Abbasanadi
//Creating Date 2018-11-20
// db.go is to handle mysql connection
// it has duties to connect to database

package models

import (
	"GolangOrdering/logger"
	"database/sql"
	"fmt"
	"sync"

	//mysql packages
	_ "github.com/go-sql-driver/mysql"
)

var instance *sql.DB
var once sync.Once

//GetDBInstance using singletone to return one instance of database
func GetDBInstance(user, password, server, port, dbname string) (*sql.DB, error) {
	haserror := false
	once.Do(func() {
		// Create connection string
		connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, password, server, port, dbname)
		var err error

		// Create connection pool
		instance, err = sql.Open("mysql", connString)
		if err != nil {
			haserror = true
			logger.Log.Fatal("Error creating connection pool: " + err.Error())
		}
		haserror = false

		logger.Log.Printf("Connected!\n")
	})
	if haserror {
		return nil, fmt.Errorf("DB can't be connected")
	}
	return instance, nil
}
