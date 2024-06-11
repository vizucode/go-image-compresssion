package main

import (
	"encoding/base64"
	"io"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/h2non/bimg"
)

type Image struct {
	Images string `json:"images"`
}

func main() {
	app := fiber.New(
		fiber.Config{
			BodyLimit: 1024 * 1024 * 1024,
		},
	)

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

		// compress image
		resizedImg, err := bimg.NewImage(imageBuff).Resize(200, 200)
		if err != nil {
			log.Fatal(err)
		}

		compressedImage, err := bimg.NewImage(resizedImg).Process(bimg.Options{
			Quality: 40,
		})
		if err != nil {
			log.Fatal(err)
		}

		file, err := os.Create("subject-004.jpeg")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		_, err = file.Write(compressedImage)
		if err != nil {
			log.Fatal(err)
		}

		return nil
	})

	app.Post("/convert/base64", func(ctx *fiber.Ctx) error {
		imageFile, err := ctx.FormFile("image")
		if err != nil {
			log.Fatal(err)
		}

		file, err := imageFile.Open()
		if err != nil {
			log.Fatal(err)
		}

		imageBuff, err := io.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}

		base64Image := base64.StdEncoding.EncodeToString(imageBuff)

		ctx.JSON(map[string]interface{}{
			"image": base64Image,
		}, "ok")
		return nil
	})

	log.Fatal(app.Listen(":6666"))
}
