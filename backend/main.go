package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
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

	app := fiber.New()
	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://127.0.0.1:5173",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Wrong method  you fool 👋!")
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": "products",
		})
	})

	app.Post(("/upload"), func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			log.Println(err)
			return c.SendStatus(fiber.StatusBadRequest) // => 400 "Bad Request"
		}

		// Save the file
		filePath := fmt.Sprintf("./%s", file.Filename)
		if err := c.SaveFile(file, filePath); err != nil {
			log.Println(err)
			return c.SendStatus(fiber.StatusInternalServerError) // => 500 "Internal Server Error"
		}

		// Open the file
		f, err := os.Open(filePath)
		if err != nil {
			log.Println(err)
			return c.SendStatus(fiber.StatusInternalServerError) // => 500 "Internal Server Error"
		}
		defer f.Close()

		// Read the file as a CSV
		r := csv.NewReader(f)
		records, err := r.ReadAll()
		if err != nil {
			log.Println(err)
			return c.SendStatus(fiber.StatusInternalServerError) // => 500 "Internal Server Error"
		}

		// Return the records as JSON
		return c.JSON(records)
	})
	log.Fatal(app.Listen(":3000"))
}
