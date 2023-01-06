package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Passenger struct {
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"Lastname"`
	Ticket    *Ticket `json:"ticket"`
}

type Ticket struct {
	TicketNumber int `json:"ticketnumber"`
}

var passenger []Passenger

func getPassengers(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(passenger)

}
func getPassengerId(w http.ResponseWriter, r *http.Request) {

}
func updatePassenger(w http.ResponseWriter, r *http.Request) {

}
func deletePassenger(w http.ResponseWriter, r *http.Request) {

}

func main() {
	router := mux.NewRouter()
	//mock data
	passenger = append(passenger, Passenger{Firstname: "horrace", Lastname: "you know", Ticket: &Ticket{TicketNumber: 10}})

	router.HandleFunc("/getPassengers", getPassengers).Methods("GET")
	router.HandleFunc("/getPassenger/{id}", getPassengerId).Methods("GET")
	router.HandleFunc("/updatePassenger/{id}", updatePassenger).Methods("PUT")
	router.HandleFunc("/deletePassenger/{id}", deletePassenger).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}
