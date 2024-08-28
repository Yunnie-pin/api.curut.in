package controllers

import (
	"api-curut-in/data"
	"api-curut-in/services"
	"fmt"
	"net/http"

	"api-curut-in/data/request"
	"api-curut-in/data/response"
	"api-curut-in/helpers"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

var validate = validator.New()

type AuthenticationController struct {
	authenticationService services.AuthenticationService
}

func NewAuthenticationController(service services.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{authenticationService: service}
}

func (controller *AuthenticationController) Login(ctx *gin.Context) {

	loginRequest := request.LoginRequest{}
	ctx.ShouldBindJSON(&loginRequest)
	err := validate.Struct(loginRequest)
	helpers.ErrorPanic(err)

	token, err_token := controller.authenticationService.Login(loginRequest)
	fmt.Println(err_token)
	if err_token != nil {
		webResponse := data.ResponseModel{
			Response:   http.StatusBadRequest,
			AppID:      "api.curut.id",
			Error:      "Invalid username or password",
			Controller: "controller.AuthenticationController",
			Action:     "Login",
			Result:     nil,
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	resp := response.LoginResponse{
		TokenType: "Bearer",
		Token:     token,
	}

	webResponse := data.ResponseModel{
		Response:   http.StatusOK,
		AppID:      "api.curut.id",
		Error:      "",
		Controller: "controller.AuthenticationController",
		Action:     "Login",
		Result:     resp,
	}

	// ctx.SetCookie("token", token, config.TokenMaxAge*60, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AuthenticationController) Register(ctx *gin.Context) {
	createUsersRequest := request.CreateUserRequest{}
	ctx.ShouldBindJSON(&createUsersRequest)
	err := validate.Struct(createUsersRequest)
	if err != nil {
		helpers.ErrorBinding(ctx, err, http.StatusBadRequest, "Registration Failed!")
		return
	}

	savedUser, err := controller.authenticationService.Register(createUsersRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, data.ResponseModel{
			Response:   http.StatusInternalServerError,
			AppID:      "api.curut.id",
			Error:      []string{err.Error()},
			Controller: "controller.AuthenticationController",
			Action:     "Register",
			Result:     nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, data.ResponseModel{
		Response:   http.StatusOK,
		AppID:      "api.curut.id",
		Error:      "",
		Controller: "controller.AuthenticationController",
		Action:     "Register",
		Result:     savedUser,
	})
}
