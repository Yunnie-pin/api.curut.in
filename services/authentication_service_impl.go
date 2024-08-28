package services

import (
	"errors"
	"time"

	"api-curut-in/config"
	"api-curut-in/data/request"
	"api-curut-in/helpers"
	"api-curut-in/models"
	"api-curut-in/repository"
)

type AuthenticationServiceImpl struct {
	UsersRepository repository.UsersRepository
}

func NewAuthenticationServiceImpl(usersRepository repository.UsersRepository) AuthenticationService {
	return &AuthenticationServiceImpl{
		UsersRepository: usersRepository,
	}
}

func (a *AuthenticationServiceImpl) Login(users request.LoginRequest) (string, error) {
	new_users, users_err := a.UsersRepository.FindByUsername(users.Username)
	if users_err != nil {
		return "", errors.New("invalid username or Password")
	}

	config, _ := config.LoadConfig(".")

	verify_error := helpers.VerifyPassword(new_users.Password, users.Password)
	if verify_error != nil {
		return "", errors.New("invalid username or Password")
	}

	// Generate Token
	token, err_token := helpers.GenerateToken(config.TokenExpiresIn, new_users.ID, config.SecretKey)
	helpers.ErrorPanic(err_token)
	return token, nil

}

// Register implements AuthenticationService
func (a *AuthenticationServiceImpl) Register(u request.CreateUserRequest) (models.UserResponse, error) {

	hashedPassword, err := helpers.HashPassword(u.Password)
	helpers.ErrorPanic(err)

	newUser := models.Users{
		Username:  u.Username,
		Email:     u.Email,
		Password:  hashedPassword,
		RolesID:   2,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	savedUser, err := a.UsersRepository.Save(newUser)

	users := models.UserResponse{
		ID:        savedUser.ID,
		Username:  savedUser.Username,
		Email:     savedUser.Email,
		CreatedAt: savedUser.CreatedAt,
		UpdatedAt: savedUser.UpdatedAt,
		RolesID:   savedUser.RolesID,
	}

	return users, err
}
