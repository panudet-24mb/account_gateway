package cmd

import (
	"account_gateway/internal/config"
	"account_gateway/internal/repository"
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
	db := config.InitDatabase()

	userdb := repository.NewUserRepositoryDB(db)
	userdb.Create(repository.User{
		UserName: "test",
	})
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
