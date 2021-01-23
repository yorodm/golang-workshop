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

type item struct {
	ID     int    `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	IsDone bool   `json:"isdone,omitempty"`
}

var items = []item {}

func homeEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the To-Do Web API")
}

func createItemEndpoint(w http.ResponseWriter, r *http.Request) {
	var newItem item
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	json.Unmarshal(reqBody, &newItem)
	items = append(items, newItem)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newItem)
}

func getItemEndpoint(w http.ResponseWriter, r *http.Request) {
	itemID, _ := strconv.Atoi(mux.Vars(r)["id"])

	for _, singleItem := range items {
		if singleItem.ID == itemID {
			json.NewEncoder(w).Encode(singleItem)
		}
	}
}

func getItemsEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(items)
}

func updateItemEndpoint(w http.ResponseWriter, r *http.Request) {
	itemID, _ := strconv.Atoi(mux.Vars(r)["id"])
	var updatedItem item

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	json.Unmarshal(reqBody, &updatedItem)

	for i, singleItem := range items {
		if singleItem.ID == itemID {
			singleItem.Title = updatedItem.Title
			singleItem.IsDone = updatedItem.IsDone
			items = append(items[:i], singleItem)
			json.NewEncoder(w).Encode(singleItem)
		}
	}
}

func deleteItemEndpoint(w http.ResponseWriter, r *http.Request) {
	itemID, _ := strconv.Atoi(mux.Vars(r)["id"])

	for i, singleItem := range items {
		if singleItem.ID == itemID {
			items = append(items[:i], items[i+1:]...)
			fmt.Fprintf(w, "The item with ID %v has been deleted successfully", itemID)
			break
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeEndpoint)
	router.HandleFunc("/api", createItemEndpoint).Methods("POST")
	router.HandleFunc("/api", getItemsEndpoint).Methods("GET")
	router.HandleFunc("/api/{id}", getItemEndpoint).Methods("GET")
	router.HandleFunc("/api/{id}", updateItemEndpoint).Methods("PATCH")
	router.HandleFunc("/api/{id}", deleteItemEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
