package main

import (
	"log"
	"users-ms/internal/app"
	"users-ms/internal/models"
	"users-ms/logging"
	"users-ms/pkg/database"
)

func init() {
	// Установим логгер
	logging.SetLogger()

	// Подключимся к БД
	if err := database.ConnectToDB(); err != nil {
		log.Fatalf("error with connect to DB: %s", err.Error())
	}

	if err := database.DB.AutoMigrate(&models.UserModel{}); err != nil {
		log.Fatalf("error with migrate tables to DB: %s", err.Error())
	}
}

// @title Users microservice
func main() {
	app := app.NewApplication()

	log.Fatal(app.Run())
}
