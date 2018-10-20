//Author Mohammad Naser Abbasanadi
//Creating Date 2018-10-20
// response.go is to reformat and provide function for reformating responses

package helpers

import (
	"encoding/json"
	"net/http"
)

//RespondWithError is provide http response based on http status and message
// it shows some errors happened
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

//RespondWithJSON this is for creating response with format
//payload could be anything that allowd to e marshal
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	if payload == nil {
		payload = map[string]string{"status": "SUCCESS"}
	}
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
