package main

// Todo connect to postgres
// Restructure project

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func initRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/getPassengers", getPassengers).Methods("GET")
	router.HandleFunc("/getPassenger/{id}", getPassengerId).Methods("GET")
	router.HandleFunc("/createPassengers", createPassenger).Methods("POST")
	router.HandleFunc("/updatePassenger/{id}", updatePassenger).Methods("PUT")
	router.HandleFunc("/deletePassenger/{id}", deletePassenger).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}

func main() {

	initMigration()
	initRouter()

}
