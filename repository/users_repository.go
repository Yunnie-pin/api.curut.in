package repository

import (
	"api-curut-in/models"
)

type UsersRepository interface {
	Save(users models.Users) (*models.Users, error)
	// Update(users models.Users)
	// Delete(usersId int)
	FindById(usersId string) (models.UserResponse, error)
	FindAll() []models.UserResponse
	FindByUsername(username string) (models.Users, error)
}
