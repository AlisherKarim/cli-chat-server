package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // Use http.StatusInternalServerError instead of hardcoding 500
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func RespondWithError(w http.ResponseWriter, code int, err error) {
	if code >= http.StatusInternalServerError { // Use constants from the http package
		log.Printf("Something wrong on our end: %s\n", err.Error())
	}

	type errorStruct struct {
		Error string `json:"error"`
	}

	RespondWithJson(w, code, errorStruct{
		Error: err.Error(),
	})
}

func RespondWithErrorMsg(w http.ResponseWriter, code int, msg string) {
	if code >= http.StatusInternalServerError { // Use constants from the http package
		log.Printf("Something wrong on our end: %s\n", msg)
	}

	type errorStruct struct {
		Error string `json:"error"`
	}

	RespondWithJson(w, code, errorStruct{
		Error: msg,
	})
}
