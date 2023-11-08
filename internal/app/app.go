package app

import (
	"users-ms/internal/transport/rest"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "users-ms/docs"
)

// Структура приложения
type Application struct {
	app *fiber.App
}

// Создать новый объект приложения.
func NewApplication() *Application {

	return &Application{app: fiber.New(fiber.Config{EnablePrintRoutes: true})}
}

// Запустить приложение.
func (a *Application) Run() error {
	// Роутинг TODO: перенести в другое место

	// ----- Пользователи -----
	uh := rest.NewUserHandlers()
	a.app.Post("/api/v1/users", uh.CreateHandler)

	// ----- Документация -----
	a.app.Get("/docs/*", swagger.HandlerDefault)

	return a.app.Listen("localhost:8000")
}
