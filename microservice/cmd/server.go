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

	userGroupRepository := repository.NewUserGroupRepositoryDB(db)
	userGroupService := services.NewUserGroupService(userGroupRepository)
	userGroupHandler := handler.NewUserGroupHandler(userGroupService)

	authService := services.NewAuthService(userRepository)
	authHandler := handler.NewAuthHandler(authService)

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
	auth := api.Group("/auth")
	usergroup := api.Group("/user-group")

	auth.Post("/login", authHandler.DefaultLogin)

	user.Post("/create-new-account", userHandler.CreateNewUserAccount)
	user.Get("/get-account", userHandler.GetAccounts)
	user.Get("/get-account/:_id", userHandler.GetAccountByID)
	user.Put("/update-account", userHandler.UpdateAccounts)
	user.Put("/update-account/:_id", userHandler.UpdateAccount)
	user.Put("/delete-accout", userHandler.DeleteAccount)
	user.Get("/find-account", userHandler.FindAccount)

	usergroup.Post("/create-new-user-group", userGroupHandler.CreateNewGroup)
	usergroup.Get("/get-user-group", userGroupHandler.GetGroups)
	usergroup.Get("/get-user-group/:_id", userGroupHandler.GetGroupByID)
	usergroup.Put("/update-user-group", userGroupHandler.UpdateGroups)
	usergroup.Put("/update-user-group/:_id", userGroupHandler.UpdateGroup)
	usergroup.Put("/delete-user-group", userGroupHandler.DeleteGroup)
	usergroup.Get("/find-user-group", userGroupHandler.FindGroup)

	port := viper.GetString("app.port")
	app.Listen(":" + port)

}
