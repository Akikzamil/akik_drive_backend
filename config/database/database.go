package database

import (
	"akik_drive/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)
var DB *gorm.DB
var err error

func ConfigDatabase() {
	dsn := "host=localhost user=postgres password=Ak01888714929 dbname=akikdrive port=5433 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		panic(err)
	}
	DB.Logger= logger.Default.LogMode(logger.Info)
	DB.AutoMigrate(&models.User{})
}