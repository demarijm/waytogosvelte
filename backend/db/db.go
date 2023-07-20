package db

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Db() Prospect {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database")
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate((Prospect{}))

	// Create
	db.Create(&Prospect{
		CustomerID:  1,
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "",
		PhoneNumber: "",
		Address:     "",
		CompanyName: "",
		Website:     "",
		Notes:       "",
		DateCreated: time.Now(),
		DateUpdated: time.Now(),
	})

	var Prospect Prospect
	db.First(&Prospect, 1)

	return Prospect
}
