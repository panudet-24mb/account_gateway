package cmd

import (
	"account_gateway/internal/config"
	"account_gateway/internal/handler"
	"account_gateway/internal/repository"
	"account_gateway/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

func Execute() {
	config.InitTimeZone()
	config.InitConfig()
	db := config.InitDatabase()
	userRepository := repository.NewUserRepositoryDB(db)
	userService := services.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "UserService API V1",
	})
	app.Use(recover.New())
	app.Use(cors.New())

	api := app.Group("/api")
	user := api.Group("/user")
	user.Post("/create-new-account", userHandler.CreateNewUserAccount)

	port := viper.GetString("app.port")
	app.Listen(":" + port)

}
