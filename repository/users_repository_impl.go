package repository

import (
	"api-curut-in/helpers"
	"api-curut-in/models"
	"errors"

	"gorm.io/gorm"
)

type UsersRepositoryImpl struct {
	Db *gorm.DB
}

func NewUsersRepositoryImpl(Db *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{Db: Db}
}

func (u *UsersRepositoryImpl) FindAll() []models.UserResponse {
	var users []models.UserResponse
	results := u.Db.Preload("Roles").Find(&users)
	helpers.ErrorPanic(results.Error)

	userResponses := make([]models.UserResponse, len(users))
	// for i, user := range users {
	// 	userResponses[i] = user
	// }

	copy(userResponses, users)

	return userResponses
}

func (u *UsersRepositoryImpl) Save(users models.Users) (*models.Users, error) {
	result := u.Db.Create(&users)
	return &users, result.Error
}

func (u *UsersRepositoryImpl) FindByUsername(username string) (models.Users, error) {
	var users models.Users
	result := u.Db.First(&users, "username = ?", username)

	if result.Error != nil {
		return users, errors.New("invalid username or Password")
	}
	return users, nil
}

func (u *UsersRepositoryImpl) FindById(usersId string) (models.UserResponse, error) {
	var users models.UserResponse
	result := u.Db.Preload("Roles").Find(&users, "id = ?", usersId)
	if result != nil {
		return users, nil
	} else {
		return models.UserResponse{}, errors.New("users is not found")
	}
}
