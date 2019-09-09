package main

import (
	"encoding/json"
	"log"
	"net/http"
	"super"

	"github.com/gorilla/mux"
)
//main uses super package to get the handler functions 
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/items", super.GetItems).Methods("GET")
	router.HandleFunc("/items/{id}", super.GetItem).Methods("GET")
	router.HandleFunc("/items/{id}", super.CreateItem).Methods("POST")
	router.HandleFunc("/items/{id}", super.UpdateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", super.DeleteItem).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":12345", router))
}
