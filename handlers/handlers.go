package handlers

import (
	entity "planeTicketApi/Entity"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// database configurations

func createPassenger() {

}

func GetPassengers(c *fiber.Ctx) error {
	var passengers []entity.Passengers
	DB.Find(&passengers)
	return c.Status(200).JSON(passengers)
}
func getPassengerId() {

}
func updatePassenger() {

}
func deletePassenger() {

}
