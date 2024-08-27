package controllers

import (
	"api-curut-in/structs"
	"net/http"

	"api-curut-in/repository"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepository repository.UsersRepository
}

func NewUsersController(repository repository.UsersRepository) *UserController {
	return &UserController{userRepository: repository}
}

func (controller *UserController) GetAllUsers(ctx *gin.Context) {
	users := controller.userRepository.FindAll()
	webResponse := structs.ResponseModel{
		Response:   http.StatusOK,
		Error:      "",
		AppID:      "api.curut.id",
		Controller: "controller.UsersController",
		Action:     "GetUsers",
		Result:     users,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UserController) GetUserByToken(ctx *gin.Context) {
	id, _ := ctx.Get("userId")

	user, err := controller.userRepository.FindById(string(id.(string)))
	if err != nil {
		webResponse := structs.ResponseModel{
			Response:   http.StatusNotFound,
			Error:      "User not found",
			AppID:      "api.curut.id",
			Controller: "controller.UsersController",
			Action:     "GetUserByToken",
			Result:     nil,
		}
		ctx.JSON(http.StatusNotFound, webResponse)
		return
	}

	webResponse := structs.ResponseModel{
		Response:   http.StatusOK,
		Error:      "",
		AppID:      "api.curut.id",
		Controller: "controller.UsersController",
		Action:     "GetUser",
		Result:     user,
	}

	ctx.JSON(http.StatusOK, webResponse)

}
