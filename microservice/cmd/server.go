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

func Execute() {
	config.InitTimeZone()
	config.InitConfig()
	config.InitDatabase()

	db := config.InitDatabase()
	db.AutoMigrate(&repository.AccountGroup{}, &repository.Account{}, &repository.Gender{},
		&repository.NamePrefix{},
		&repository.Invoice{}, &repository.Subscription{},
		&repository.PlanFeature{}, &repository.SubscriptionPlan{},
		&repository.Feature{}, &repository.Plan{}, &repository.Address{}, &repository.PartnerBranch{}, &repository.Partner{},
		&repository.AddressGroup{}, &repository.AddressType{},
	)
	// AccRepo := repository.NewAccountRepositoryDB(db)

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
