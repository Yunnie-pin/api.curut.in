package helpers

import (
	"api-curut-in/data"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrorResponse(ctx *gin.Context, err error, msg string) {
	//HANDLE GORM ERROR
	ctx.JSON(http.StatusBadRequest, data.ResponseModel{
		Response:   http.StatusBadRequest,
		Error:      msg,
		AppID:      "api.curut.id",
		Controller: "-",
		Action:     "-",
		Result:     err,
	})
}

func ErrorMultipeResponse(ctx *gin.Context, err any, msg string) {
	//HANDLE GORM ERROR
	ctx.JSON(http.StatusBadRequest, data.ResponseModel{
		Response:   http.StatusBadRequest,
		Error:      msg,
		AppID:      "api.curut.id",
		Controller: "-",
		Action:     "-",
		Result:     err,
	})
}

func ErrorBinding(ctx *gin.Context, err error, statusCode int, msg string) {
	var errorMessages []string
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			errorMessage := fmt.Sprintf("ERROR: Invalid %s (%s)!", e.Field(), e.Error())
			errorMessages = append(errorMessages, errorMessage)
		}
	} else {
		errorMessages = append(errorMessages, err.Error())
	}

	ctx.JSON(http.StatusBadRequest, data.ResponseModel{
		Response:   statusCode,
		Error:      errorMessages,
		AppID:      "api.curut.id",
		Controller: "-",
		Action:     "-",
		Result:     nil,
	})
}
