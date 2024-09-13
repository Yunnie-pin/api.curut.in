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

	config, _ := config.LoadConfig()

	verify_error := helpers.VerifyPassword(new_users.Password, users.Password)
	if verify_error != nil {
		return "", errors.New("invalid username or Password")
	}

	// Generate Token
	// string to tim
	duration, err := time.ParseDuration(config.TokenExpiresIn)
	if err != nil {
		return "", err
	}

	token, err_token := helpers.GenerateToken(duration, new_users.ID, config.SecretKey)
	helpers.ErrorPanic(err_token)
	return token, nil

}

// Register implements AuthenticationService
func (a *AuthenticationServiceImpl) Register(u request.CreateUserRequest) (result models.UserResponse, err error) {

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

	if err != nil {
		return result, errors.New("Failed to save user")
	}

	idUser := string((savedUser.ID).String())

	users, err := a.UsersRepository.FindById(idUser)
	if err != nil {
		return users, err
	}

	return users, err
}
