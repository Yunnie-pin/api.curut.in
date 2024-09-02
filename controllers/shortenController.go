package controllers

import (
	"api-curut-in/data"
	"api-curut-in/helpers"
	"api-curut-in/models"
	"api-curut-in/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShortenController struct {
	shortenRepository repository.ShortenRepository
}

func NewShortenController(repository repository.ShortenRepository) *ShortenController {
	return &ShortenController{shortenRepository: repository}
}

func (s *ShortenController) GetListShorten(ctx *gin.Context) {
	shortens, err := s.shortenRepository.FindAll()
	if err != nil {
		helpers.ErrorResponse(ctx, err, "Failed to get list shorten")
		return
	}

	response := data.ResponseModel{
		Response:   http.StatusOK,
		Error:      "",
		AppID:      "api-curut-in",
		Controller: "ShortenController",
		Action:     "GetListShorten",
		Result:     shortens,
	}

	ctx.JSON(http.StatusOK, response)
}

func (s *ShortenController) CreateShorten(ctx *gin.Context) {
	createShortenRequest := models.CreateShortenRequest{}

	ctx.ShouldBindJSON(&createShortenRequest)

	err := validate.Struct(createShortenRequest)

	if err != nil {
		helpers.ErrorResponse(ctx, err, "Error validate request")
		return
	}

	shorten := models.Shorten{
		Shorten:  createShortenRequest.Shorten,
		Original: createShortenRequest.Original,
	}

	newShorten, err := s.shortenRepository.Save(shorten)
	if err != nil {
		helpers.ErrorResponse(ctx, err, "Failed to save shorten")
		return
	}

	response := data.ResponseModel{
		Response:   http.StatusOK,
		Error:      "",
		AppID:      "api-curut-in",
		Controller: "ShortenController",
		Action:     "CreateShorten",
		Result:     newShorten,
	}

	ctx.JSON(http.StatusOK, response)

}
