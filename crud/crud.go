package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// database configurations

func createPassenger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var passenger Passengers
	_ = json.NewDecoder(r.Body).Decode(&passenger)
	DB.Create(&passenger)
	json.NewEncoder(w).Encode(passenger)
}

func getPassengers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var passengers []Passengers
	DB.Find(&passengers)
	json.NewEncoder(w).Encode(passengers)

}
func getPassengerId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var passenger []Passengers
	params := mux.Vars(r)
	DB.First(&passenger, params["id"])
	json.NewEncoder(w).Encode(passenger)

}
func updatePassenger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/json")
	var passenger []Passengers
	params := mux.Vars(r)
	DB.First(&passenger, params["id"])
	_ = json.NewDecoder(r.Body).Decode(&passenger)
	DB.Save(&passenger)
	json.NewEncoder(w).Encode(passenger)

}
func deletePassenger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/json")
	var passenger []Passengers
	params := mux.Vars(r)
	DB.Delete(&passenger, params["id"])
	json.NewEncoder(w).Encode("passenger deleted")
}
