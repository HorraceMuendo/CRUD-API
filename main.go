package main

import (
	"fmt"
	"log"
	database "planeTicketApi/database_config"
	"planeTicketApi/handlers"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func initRouter() {
	app := fiber.New(fiber.Config{AppName: "PlaneTicketApi v1.0.0"})

	app.Get("/getPassengers", handlers.GetPassengers)
	app.Get("/getPassengers/:id", handlers.GetPassengerById)
	app.Post("/getPassengers", handlers.CreatePassenger)
	app.Put("/updatePassengers/:id", handlers.UpdatePassenger)
	app.Delete("/deletePassengers/:id", handlers.DeletePassenger)

	fmt.Printf("Starting server at port 8000\n")

	log.Fatal(app.Server().ListenAndServe(":8000"))
}

func main() {
	database.InitMigration()

	// fmt.Printf("Starting server at port 8000\n")
	// if err := http.ListenAndServe(":8000", nil); err != nil {
	// 	log.Fatal(err)
	// }
	initRouter()
	//log.Fatal(app.Listen(":3000"))

}
