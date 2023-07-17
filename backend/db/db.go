package db
import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func Db() Product {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database")
		panic("failed to connect database")
	}
	
	// Migrate the schema
	db.AutoMigrate((Product{}))
	
	// Create 
	db.Create(&Product{Code: "D42", Price: 100})     
	
	var product Product
	db.First(&product, 1)

	return product
}