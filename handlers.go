package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func AmazonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusOK, Text: "Enter Amazon Product Id"})
	if err != nil {
		panic(err)
	}
	return
}

func AmazonScrappingHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var amazonId string

	amazonId = vars["amazon_id"]

	movie, jerr := AmazonScraper(amazonId)

  // if error has been recieved
	if jerr.Text != "" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(jerr.Code)

		err := json.NewEncoder(w).Encode(jerr)
		if err != nil {
			panic(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(movie);
	if err != nil {
		panic(err)
	}
	return

}
