package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		fmt.Println("Responding with 500 error", msg)
	}

	type userError struct {
		Error string `json:"error"`
	}

	respondWithJson(w,code,userError{Error:msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {

	dat, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("error marshalling json", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}

