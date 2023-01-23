package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initMigration() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)

	DB, err := gorm.Open(postgres.Open(dbinfo), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("connection unsuccesful")
	} else {
		fmt.Println("connected")
	}

	DB.AutoMigrate(&Passengers{})
}
