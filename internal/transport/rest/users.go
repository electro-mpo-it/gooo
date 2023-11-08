package rest

import (
	"users-ms/internal/crud"
	"users-ms/internal/models"
	"users-ms/internal/schemas"

	"github.com/gofiber/fiber/v2"
)

// Структура содержит в себе хендлеры для взаимодействия с пользователями
type UserHandlers struct{}

// Создать новый объект.
func NewUserHandlers() *UserHandlers {

	return &UserHandlers{}
}

// @Summary      Создать пользователя
// @Tags         users
// @Accept       json
// @Produce      json
// @Param 		 request body schemas.UserCreateSchema true "RequestBody"
// @Failure      422
// @Router       /api/v1/users [post]
func (*UserHandlers) CreateHandler(ctx *fiber.Ctx) error {
	var userData schemas.UserCreateSchema
	err := ctx.BodyParser(&userData)

	if err != nil {

		return err
	}

	if err := userData.Validate(); err != nil {

		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(err.Error())
	}

	err = crud.UsersCreate(
		&models.UserModel{
			FirstName: userData.FirstName,
			LastName:  userData.LastName,
			Password:  userData.Password,
		},
	)

	if err != nil {

		ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return nil
}
