package main

import (
	"fmt"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type uploadRequest struct {
	IDCardNumber string `json:"id_card_number"`
}

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:  "*",
		AllowHeaders:  "*",
		AllowMethods:  "GET, POST, PUT, PATCH, DELETE, OPTIONS",
		ExposeHeaders: "content-disposition",
	}))

	app.Post("/upload", upload)

	if err := app.Listen(":8081"); err != nil {
		panic(err)
	}
}

func upload(c *fiber.Ctx) error {
	fmt.Println("Called Upload")
	var body uploadRequest
	var err error
	if err = c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	var idCardImage *multipart.FileHeader
	idCardImage, err = c.FormFile("id_card_image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	fmt.Println(idCardImage)

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}
