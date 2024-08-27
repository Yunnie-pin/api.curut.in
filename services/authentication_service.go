package services

import (
	"api-curut-in/data/request"
	"api-curut-in/models"
)

type AuthenticationService interface {
	Login(users request.LoginRequest) (string, error)
	Register(users request.CreateUserRequest) (models.UserResponse, error)
}
