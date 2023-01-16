package config

import (
	"context"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Book struct {
	Title  string
	Author string
}

func InitDatabase() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(viper.GetString("db.mongo_host")))
	if err != nil {
		panic(err)
	}

	return client
}

//old database pgsql
// func InitDatabase() *gorm.DB {
// 	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
// 		viper.GetString("db.host"),
// 		viper.GetString("db.user"),
// 		viper.GetString("db.password"),
// 		viper.GetString("db.dbname"),
// 		viper.GetInt("db.port"),
// 		viper.GetString("db.sslmode"),
// 		viper.GetString("db.timezone"),
// 	)
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	return db

// }
