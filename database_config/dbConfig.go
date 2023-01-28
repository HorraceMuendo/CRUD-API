package database

import (
	"fmt"
	entity "planeTicketApi/Entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	Host        = "localhost"
	PORT        = 5432
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "planeTicketApi"
)

func InitMigration() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s port=%d host=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME, PORT, Host)

	DB, err := gorm.Open(postgres.Open(dbinfo), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("connection unsuccesful")
	} else {
		fmt.Println("connected")
	}

	DB.AutoMigrate(&entity.Passengers{})
}
