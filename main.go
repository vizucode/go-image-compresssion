package main

import (
	"encoding/base64"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Image struct {
	Images string `json:"images"`
}

func main() {
	app := fiber.New()

	app.Use(recover.New())

	app.Post("/compress-image", func(ctx *fiber.Ctx) error {
		var Payload Image
		err := ctx.BodyParser(&Payload)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("This is Debug ===>>>")
		log.Println(Payload.Images)

		imageBuff, err := base64.StdEncoding.DecodeString(Payload.Images)
		if err != nil {
			log.Fatal(err)
		}

		file, err := os.Create("subject-001.jpeg")
		if err != nil {
			log.Fatal(err)
		}

		_, err = file.Write(imageBuff)
		if err != nil {
			log.Fatal(err)
		}

		return nil
	})

	log.Fatal(app.Listen(":6666"))
}
