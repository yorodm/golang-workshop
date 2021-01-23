package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type data struct {
	Num  int    `json:"Num"`
	Text string `json:"Text"`
}

func create(w http.ResponseWriter, r *http.Request) {
	var newData data
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	json.Unmarshal(reqBody, &newData)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newData)
}

func getOne(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The item with %v has been requested", mux.Vars(r)["num"])
}

func getAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The all items has been requested")
}

func update(w http.ResponseWriter, r *http.Request) {
	var updatedData data
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	json.Unmarshal(reqBody, &updatedData)

	json.NewEncoder(w).Encode(updatedData)
}

func delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The item with %v has been deleted successfully", mux.Vars(r)["num"])
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api", create).Methods("POST")
	router.HandleFunc("/api", getAll).Methods("GET")
	router.HandleFunc("/api/{num}", getOne).Methods("GET")
	router.HandleFunc("/api/{num}", update).Methods("PATCH")
	router.HandleFunc("/api/{num}", delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
