package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const (
	DB_USER     = "postgres"
	PORT        = 5432
	DB_PASSWORD = ""
	DB_NAME     = "planeTicketApi"
)

type Passengers struct {
	gorm.Model
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"Lastname"`
	Ticket    *Ticket `json:"ticket"`
}

type Ticket struct {
	TicketNumber int `json:"ticketnumber"`
}

// database configurations
func initMigration() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	DB, err = gorm.Open(postgres.Open(dbinfo), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("connection unsuccesful")
	}
	DB.AutoMigrate(&Passengers{})
}

func createPassenger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var passenger Passengers
	_ = json.NewDecoder(r.Body).Decode(&passenger)
	DB.Create(&passenger)
	json.NewEncoder(w).Encode(passenger)
}

func getPassengers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}
func getPassengerId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}
func updatePassenger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}
func deletePassenger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}
