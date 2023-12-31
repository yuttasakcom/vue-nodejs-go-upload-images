package main

import (
	"fmt"
	"log"
	"mime/multipart"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/caarlos0/env/v9"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

type uploadRequest struct {
	IDCardNumber string `form:"id_card_number"`
	BankNumber   string `form:"bank_number"`
}

type AWSConfig struct {
	AccessKeyID     string `env:"AWS_ACCESS_KEY_ID" json:"aws_access_key_id"`
	AccessKeySecret string `env:"AWS_SECRET_ACCESS_KEY" json:"aws_secret_access_key"`
	Region          string `env:"AWS_REGION" json:"aws_region"`
	BucketName      string `env:"AWS_BUCKET_NAME" json:"aws_bucket_name"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %s\n", err)
	}
	var awsCfg AWSConfig
	if err := env.Parse(&awsCfg); err != nil {
		log.Printf("Error parsing .env file: %s\n", err)
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:  "*",
		AllowHeaders:  "*",
		AllowMethods:  "GET, POST, PUT, PATCH, DELETE, OPTIONS",
		ExposeHeaders: "content-disposition",
	}))

	app.Post("/upload", uploadHandler(awsCfg))

	if err := app.Listen(":8081"); err != nil {
		panic(err)
	}
}

func uploadHandler(awsCfg AWSConfig) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		apiKey := c.Get("x-api-key")
		fmt.Printf("x-api-key %s\n", apiKey)
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

		wg := new(sync.WaitGroup)
		wg.Add(2)

		go func() {
			uploadImageToS3(awsCfg, idCardImage)
			wg.Done()
		}()

		var bankImage *multipart.FileHeader
		bankImage, err = c.FormFile("bank_image")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Bad Request",
			})
		}
		go func() {
			uploadImageToS3(awsCfg, bankImage)
			wg.Done()
		}()

		wg.Wait()

		fmt.Printf("id_card_number %s\n", body.IDCardNumber)
		if idCardImage != nil {
			fmt.Printf("id_card_image name %s\n", idCardImage.Filename)
		}

		fmt.Printf("bank_number %s\n", body.BankNumber)
		if bankImage != nil {
			fmt.Printf("bank_image name %s\n", bankImage.Filename)
		}

		return c.JSON(fiber.Map{
			"message": "Success",
		})
	}
}

func uploadImageToS3(awsCfg AWSConfig, fileHeader *multipart.FileHeader) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsCfg.Region),
	})
	if err != nil {
		return err
	}

	svc := s3.New(sess)
	input := &s3.PutObjectInput{
		Body:   file,
		Bucket: aws.String(awsCfg.BucketName),
		Key:    aws.String(fileHeader.Filename),
	}
	_, err = svc.PutObject(input)
	if err != nil {
		return err
	}

	return nil
}
