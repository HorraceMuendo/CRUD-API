package main

import (
	"fmt"
	"log"
	"net/http"
	database "planeTicketApi/database_config"

	// "github.com/gofiber/fiber"
	// "github.com/gofiber/fiber/v2"

	_ "github.com/lib/pq"
)

// func initRouter() {
// 	app := fiber.New(fiber.Config{AppName: "PlaneTicketApi v1.0.0"})

// 	app.Get("/getPassengers")
// 	//router := mux.NewRouter()

// 	// router.HandleFunc("/getPassengers", getPassengers).Methods("GET")
// 	// router.HandleFunc("/getPassenger/{id}", getPassengerId).Methods("GET")
// 	// router.HandleFunc("/createPassengers", createPassenger).Methods("POST")
// 	// router.HandleFunc("/updatePassenger/{id}", updatePassenger).Methods("PUT")
// 	// router.HandleFunc("/deletePassenger/{id}", deletePassenger).Methods("DELETE")

// 	fmt.Printf("Starting server at port 8000\n")
// 	if err := http.ListenAndServe(":8000", app); err != nil {
// 		log.Fatal(err)
// 	}
// }

func main() {
	database.InitMigration()

	fmt.Printf("Starting server at port 8000\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
	//initRouter()

}
