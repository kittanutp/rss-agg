package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal((payload))
	if err != nil {
		w.WriteHeader(500)
		log.Println(err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	err_msg := fmt.Sprintf("Responding with %v error, %v", code, msg)
	log.Println(err_msg)
	type errResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errResponse{Error: err_msg})
}
