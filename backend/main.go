package main

import (
	"backend/technologies"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

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
	app.Post("/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		// Save the file
		filePath := fmt.Sprintf("./%s", file.Filename)
		if err := c.SaveFile(file, filePath); err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Open the file
		f, err := os.Open(filePath)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		defer f.Close()

		// Read the file as a CSV
		r := csv.NewReader(f)
		records, err := r.ReadAll()
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// Convert CSV records to JSON-friendly structure
		var result []map[string]interface{}  // Use 'interface{}' to allow adding a slice of strings
		headers := records[0]                // Assume the first record is the headers
		for _, record := range records[1:] { // Skip the headers
			row := make(map[string]interface{}) // Use 'interface{}' to allow adding a slice of strings
			for i, header := range headers {
				row[header] = record[i]

				// If the header is "Website", check if the website is active and what technologies are used
				if header == "Website" {
					resp, err := http.Get(record[i])
					if err != nil || resp.StatusCode != 200 {
						row["Status"] = "inactive"
					} else {
						row["Status"] = "active"
						techs, err := technologies.CheckTechnology(record[i])
						if err != nil {
							log.Println(err)
							// handle error appropriately here
						}
						row["Technologies"] = strings.Join(techs, ", ")
						if err != nil {
							log.Println("Error checking technologies:", err)
						} else {
							row["Technologies"] = techs
						}
					}
				}
			}
			result = append(result, row)
		}
		defer os.Remove(filePath)

		// Return the result as JSON
		return c.JSON(result)
	})
	
	log.Fatal(app.Listen(":3000"))
}
