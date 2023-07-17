package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

// FormFile represents multipart form file
type FormFile struct {
    // Fieldname is form file's field name
    Fieldname string
    // Name is form file's name
    Name string
    // Content is form file's content
    Content []byte
}



func uploadFile(c *fiber.Ctx) error {
	ff1 := &FormFile{"filename1", "field name1", []byte("content")}
	fmt.Println("File Upload Endpoint Hit", ff1)

	
	if c.Response().StatusCode() != 200 {
		return c.JSON(fiber.Map{
			"error": "Error",
		})
	}
	return c.JSON(fiber.Map{
		"success": "Success",
	})
}




  func main() {

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

		fmt.Println(product)


	  app := fiber.New()
	  app.Use(cors.New())
	  app.Use(cors.New(cors.Config{
		AllowOrigins: "http://127.0.0.1:5174",
	}))

    app.Post("/", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
			"data": product,
		})
    })

	app.Post(("/upload"), uploadFile)

    app.Listen(":3000")
}