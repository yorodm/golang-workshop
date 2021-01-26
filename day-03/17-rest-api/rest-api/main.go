package main

import (
	"log"
	"net/http"
)

type H = func(w http.ResponseWriter, r *http.Request)

func dispatchResource(get, post, put, delete H) H {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "GET":
			get(w, r)
		case "POST":
			post(w, r)
		case "PUT":
			put(w, r)
		case "DELETE":
			delete(w, r)
		default:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"message": "not found"}`))
		}
	}
}

func main() {

	// Create a mux for routing incoming requests
	m := http.NewServeMux()

	getFunc := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "get called"}`))
	}
	postFunc := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "post called"}`))
	}
	m.HandleFunc("/resource", dispatchResource(getFunc, postFunc, nil, nil))

	// Create a server listening on port 8000
	s := &http.Server{
		Addr:    ":8000",
		Handler: m,
	}

	// Continue to process new requests until an error occurs
	log.Fatal(s.ListenAndServe())
}
