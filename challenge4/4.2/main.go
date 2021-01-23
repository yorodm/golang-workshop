package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var repositoryType string

func homeEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the To-Do Web API")
}

func createItemEndpoint(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	json.Unmarshal(reqBody, &newItem)

	repo := createRepository()
	repo.CreateItem(newItem)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "OK")
}

func getItemEndpoint(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	repo := createRepository()
	json.NewEncoder(w).Encode(repo.GetItem(id))
}

func getItemsEndpoint(w http.ResponseWriter, r *http.Request) {
	repo := createRepository()
	json.NewEncoder(w).Encode(repo.GetItems())
}

func updateItemEndpoint(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var updatedItem Item

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	json.Unmarshal(reqBody, &updatedItem)

	updatedItem.ID = id

	repo := createRepository()
	repo.UpdateItem(updatedItem)

	fmt.Fprintf(w, "OK")
}

func deleteItemEndpoint(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	repo := createRepository()
	repo.DeleteItem(id)

	fmt.Fprintf(w, "OK")
}

func createRepository() TodoRepository {
	if repositoryType == "Mongo" {
		return &MongoRepository{}
	} else {
		return &InMemoryRepository{}
	}
}

func main() {
	repositoryType = ""

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeEndpoint)
	router.HandleFunc("/api", createItemEndpoint).Methods("POST")
	router.HandleFunc("/api", getItemsEndpoint).Methods("GET")
	router.HandleFunc("/api/{id}", getItemEndpoint).Methods("GET")
	router.HandleFunc("/api/{id}", updateItemEndpoint).Methods("PATCH")
	router.HandleFunc("/api/{id}", deleteItemEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
