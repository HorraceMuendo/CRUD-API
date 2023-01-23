package database

import (
	"fmt"
	entity "planeTicketApi/Entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	PORT        = 5432
	DB_USER     = "postgres"
	DB_PASSWORD = ""
	DB_NAME     = "planeTicketApi"
)

func InitMigration() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s port=%d sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME, PORT)

	DB, err := gorm.Open(postgres.Open(dbinfo), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("connection unsuccesful")
	} else {
		fmt.Println("connected")
	}

	DB.AutoMigrate(&entity.Passengers{})
}
