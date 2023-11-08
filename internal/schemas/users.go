package schemas

import "github.com/go-playground/validator/v10"

type UserCreateSchema struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name" validate:"min=17"`
	Password  string `json:"password" validate:"required"`
}

// Валидация полей
func (ucs *UserCreateSchema) Validate() error {
	validator := validator.New()

	return validator.Struct(ucs)
}
