package database

import (
	"fmt"
	"users-ms/logging"
	"users-ms/settings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Объект подключения к Базе Данных. (Все CRUD манипуляции через него.)
var DB *gorm.DB

// Подключение к Базе Данных.
// Функция помещает в переменную DB объект *gorm.DB
func ConnectToDB() error {
	logging.Logger.Info("Connect to DB..")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s TimeZone=%s",
		settings.DBHost, settings.DBPort, settings.DBUser, settings.DBPassword, settings.DBName, settings.DBTimeZone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {

		return err
	}

	DB = db

	return nil
}
