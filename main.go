package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		fmt.Fprint(w, "Get requests for id: ", id)
	}).Methods("GET")
	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Get requests")
	}).Methods("GET")
	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Post requests")
	}).Methods("POST")
	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Put requests")
	}).Methods("PUT")
	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		fmt.Fprint(w, "Delete requests for id: ", id)
	}).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
