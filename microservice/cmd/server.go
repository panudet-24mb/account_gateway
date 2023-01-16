package cmd

import (
	"account_gateway/internal/config"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

type Book struct {
	Title  string
	Author string
}

func Execute() {
	config.InitTimeZone()
	config.InitConfig()
	// client := config.InitDatabase()
	// coll := client.Database("db").Collection("books")
	// doc := Book{Title: "Atonement", Author: "Ian McEw"}
	// result, _ := coll.InsertOne(context.TODO(), doc)
	// fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Account API V1",
	})
	app.Use(recover.New())
	app.Use(cors.New())
	app.Get("/", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ hello")
		return c.SendString(msg) // => ✋ register
	})

	port := viper.GetString("app.port")
	app.Listen(":" + port)

}
