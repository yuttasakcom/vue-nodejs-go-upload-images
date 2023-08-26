package main

import (
	"fmt"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type uploadRequest struct {
	IDCardNumber string `form:"id_card_number"`
	BankNumber   string `form:"bank_number"`
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

	var bankImage *multipart.FileHeader
	bankImage, err = c.FormFile("bank_image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	fmt.Printf("id_card_number %s\n", body.IDCardNumber)
	if idCardImage != nil {
		fmt.Printf("id_card_image name %s\n", idCardImage.Filename)
	}

	fmt.Printf("bank_number %s\n", body.BankNumber)
	if bankImage != nil {
		fmt.Printf("id_card_image name %s\n", bankImage.Filename)
	}

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}
