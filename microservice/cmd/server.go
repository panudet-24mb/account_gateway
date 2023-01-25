package cmd

import (
	"account_services/internal/config"
	"account_services/internal/handler"
	"account_services/internal/repository"
	"account_services/internal/services"

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

	roleRepository := repository.NewRoleRepositoryDB(db)
	roleService := services.NewRoleService(roleRepository)
	roleHandler := handler.NewRoleHandler(roleService)

	userRoleRepository := repository.NewUserRoleRepositoryDB(db)
	userRoleService := services.NewUserRoleService(userRoleRepository)
	userRoleHandler := handler.NewUserRoleHandler(userRoleService)

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
	usergroup := api.Group("/user-group")
	userrole := api.Group("/user-role")

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
	// usergroup.Get("/find-user-group", userGroupHandler.FindGroup)

	userrole.Post("/create-new-role", roleHandler.CreateNewRole)
	userrole.Get("/get-role", roleHandler.GetRoles)
	userrole.Get("/get-role/:_id", roleHandler.GetRoleByID)
	userrole.Put("/update-role", roleHandler.UpdateRoles)
	userrole.Put("/update-role/:_id", roleHandler.UpdateRole)
	userrole.Put("/delete-role", roleHandler.DeleteRole)

	userrole.Post("/create-new-user-role", userRoleHandler.CreateNewUserRole)
	userrole.Get("/get-user-role", userRoleHandler.GetUserRole)
	userrole.Get("/get-user-role/:_id", userRoleHandler.GetUserRoleByID)
	userrole.Put("/update-user-role", userRoleHandler.UpdateUsersRole)
	userrole.Put("/update-user-role/:_id", userRoleHandler.UpdateUserRole)
	userrole.Put("/delete-user-role/:_id", userRoleHandler.DeleteUserRole)

	port := viper.GetString("app.port")
	app.Listen(":" + port)

}
