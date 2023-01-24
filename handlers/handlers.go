package handlers

import (
	entity "planeTicketApi/Entity"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// database configurations

func CreatePassenger(c *fiber.Ctx) error {
	passenger := new(entity.Passengers)
	if err := c.BodyParser(passenger); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	DB.Create(&passenger)
	return c.Status(200).JSON(passenger)
}

func GetPassengers(c *fiber.Ctx) error {
	var passengers []entity.Passengers
	DB.Find(&passengers)
	return c.Status(200).JSON(passengers)
}
func GetPassengerById(c *fiber.Ctx) error {
	id := c.Params("id")
	var passenger entity.Passengers
	match := DB.Find(&passenger, id)

	if match.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&passenger)

}
func UpdatePassenger(c *fiber.Ctx) error {
	passenger := new(entity.Passengers)
	id := c.Params("id")

	if err := c.BodyParser(passenger); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	DB.Where("id=?", id).Updates(&passenger)
	return c.Status(200).JSON(passenger)

}
func DeletePassenger(c *fiber.Ctx) error {
	var passenger entity.Passengers
	id := c.Params("id")
	delete := DB.Delete(&passenger, id)

	if delete.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.SendStatus(200)
}
