//Author Mohammad Naser Abbasanadi
//Creating Date 2018-10-20
// middleware.go is to handle middlewares
// it has duties to do sumaction over all request example : security check , logging

package controller

import (
	"GolangOrdering/logger"
	"net/http"
)

//withLogging log all of requested from outside
func withLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Log.Printf("Logged connection from %s calling API:%s - %s ", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	}
}

//securityCheck will check security issues if there is
func securityCheck(next http.HandlerFunc) http.HandlerFunc {
	checking := false
	return func(w http.ResponseWriter, r *http.Request) {
		if checking {
			//do something here abas
		} else {
			next.ServeHTTP(w, r)
		}
	}
}

//middlewares combile all of middlewares to one single point
func middlewares(next http.HandlerFunc) http.HandlerFunc {
	return (securityCheck(withLogging(next)))
}
