package main

import (
	"fmt"
	"log"
	configuration "seacrust-backend/sdk"
	"seacrust-backend/src/handler"
	"seacrust-backend/src/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	config := configuration.Init()
	err := config.CanLoad(".env")
	if err != nil {
		log.Fatalln("file belum ada brooo")
	}

	dbParams := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Get("DB_USERNAME"),
		config.Get("DB_PASSWORD"),
		config.Get("DB_HOST"),
		config.Get("DB_PORT"),
		config.Get("DB_DATABASE"),
	)
	database, err := gorm.Open(mysql.Open(dbParams), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln(err)
	}

	database.AutoMigrate(models.User{})
	database.AutoMigrate(models.Category{})
	database.AutoMigrate(models.Toko{})
	database.AutoMigrate(models.Produk{})
	database.AutoMigrate(models.Address{})
	database.AutoMigrate(models.Pesanan{})

	handler := handler.Init(database)

	if err := handler.SeedCategory(database); err != nil {
		fmt.Println(err)
		panic("GAGAL SEED Category")
	}

	handler.Run()
}
