package crud

import (
	"users-ms/internal/models"
	"users-ms/pkg/database"
)

func UsersCreate(user *models.UserModel) error {
	tx := database.DB.Create(user)

	if tx.Error != nil {

		return tx.Error
	}

	return nil
}
