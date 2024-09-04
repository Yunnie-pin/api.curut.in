package controllers

import (
	"api-curut-in/data"
	"api-curut-in/helpers"
	"api-curut-in/repository"

	"github.com/gin-gonic/gin"
)

type RedirectController struct {
	shortenRepository repository.ShortenRepository
}

func NewRedirectController(repository repository.ShortenRepository) *RedirectController {
	return &RedirectController{shortenRepository: repository}
}

func (r *RedirectController) Get(ctx *gin.Context) {
	name := ctx.Param("name")
	shorten, err := r.shortenRepository.FindShortenByName(name)
	if err != nil {
		helpers.ErrorResponse(ctx, err, "Url Not Found")
		return
	}

	response := data.ResponseModel{
		Response:   200,
		Error:      "",
		AppID:      "skincare-server",
		Controller: "articles",
		Action:     "GetArticleByID",
		Result:     shorten,
	}

	ctx.JSON(200, response)
}
