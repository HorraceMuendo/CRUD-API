package main

// Todo connect to postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	PORT        = 5432
	DB_PASSWORD = ""
	DB_NAME     = "planeTicketApi"
)

type Passengers struct {
	Id        string  `json:"id"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"Lastname"`
	Ticket    *Ticket `json:"ticket"`
}

type Ticket struct {
	TicketNumber int `json:"ticketnumber"`
}

var passengers []Passengers

func createPassenger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var passenger Passengers
	_ = json.NewDecoder(r.Body).Decode(&passenger)
	passenger.Id = strconv.Itoa(rand.Intn(10000000))
	passengers = append(passengers, passenger)

	json.NewEncoder(w).Encode(passenger)
}

func getPassengers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(passengers)

}
func getPassengerId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, passenger := range passengers {
		if passenger.Id == params["id"] {
			json.NewEncoder(w).Encode(passenger)
			return
		}
	}

}
func updatePassenger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, passenger := range passengers {
		if passenger.Id == params["id"] {
			passengers = append(passengers[:index], passengers[index+1:]...)

			var passenger Passengers
			_ = json.NewDecoder(r.Body).Decode(&passenger)
			passenger.Id = strconv.Itoa(rand.Intn(10000000))
			passengers = append(passengers, passenger)

			json.NewEncoder(w).Encode(passenger)

			return
		}
	}

}
func deletePassenger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, passenger := range passengers {
		if passenger.Id == params["id"] {
			passengers = append(passengers[:index], passengers[index+1:]...)
			break
		}
	}
}

func main() {
	router := mux.NewRouter()
	//mock data
	// passengers = append(passengers, Passengers{Id: "1", Firstname: "horrace", Lastname: "muendo", Ticket: &Ticket{TicketNumber: 10}})
	// passengers = append(passengers, Passengers{Id: "2", Firstname: "maggie", Lastname: "wambui", Ticket: &Ticket{TicketNumber: 13}})
	// passengers = append(passengers, Passengers{Id: "3", Firstname: "victor", Lastname: "muchui", Ticket: &Ticket{TicketNumber: 17}})
	// passengers = append(passengers, Passengers{Id: "4", Firstname: "fortune", Lastname: "timothy", Ticket: &Ticket{TicketNumber: 20}})

	router.HandleFunc("/getPassengers", getPassengers).Methods("GET")
	router.HandleFunc("/getPassenger/{id}", getPassengerId).Methods("GET")
	router.HandleFunc("/createPassengers", createPassenger).Methods("POST")
	router.HandleFunc("/updatePassenger/{id}", updatePassenger).Methods("PUT")
	router.HandleFunc("/deletePassenger/{id}", deletePassenger).Methods("DELETE")

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("connection succesful")

	fmt.Printf("Starting server at port 8000\n")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}
