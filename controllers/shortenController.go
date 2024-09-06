package controllers

import (
	"api-curut-in/data"
	"api-curut-in/helpers"
	"api-curut-in/models"
	"api-curut-in/repository"
	"net/http"
	"time"

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

func (s *ShortenController) CreateShortenByPub(ctx *gin.Context) {
	createShortenRequest := models.CreateShortenRequest{}

	ctx.ShouldBindJSON(&createShortenRequest)

	err := validate.Struct(createShortenRequest)

	if err != nil {
		helpers.ErrorResponse(ctx, err, "Error validate request")
		return
	}

	us, err := s.shortenRepository.IsShortenAlreadyTaken(createShortenRequest.Shorten)
	if err != nil {
		helpers.ErrorResponse(ctx, err, "Failed to check shorten")
		return
	}

	if us {
		helpers.ErrorResponse(ctx, err, "Shorten already taken")
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

func (s *ShortenController) CreateShortenByUser(ctx *gin.Context) {
	userId := ctx.Value("userId")
	createShortenRequest := models.CreateShortenRequest{}

	ctx.ShouldBindJSON(&createShortenRequest)

	err := validate.Struct(createShortenRequest)

	if err != nil {
		helpers.ErrorResponse(ctx, err, "Error validate request")
		return
	}

	us, err := s.shortenRepository.IsShortenAlreadyTaken(createShortenRequest.Shorten)
	if err != nil {
		helpers.ErrorResponse(ctx, err, "Failed to check shorten")
		return
	}

	if us {
		helpers.ErrorResponse(ctx, err, "Shorten already taken")
		return
	}

	shorten := models.Shorten{
		Shorten:   createShortenRequest.Shorten,
		Original:  createShortenRequest.Original,
		CreatedBy: userId.(string),
		UpdatedBy: userId.(string),
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

func (s *ShortenController) UpdateShorten(ctx *gin.Context) {
	userId := ctx.Value("userId")
	shortenId := ctx.Param("id")

	shorten, err := s.shortenRepository.FindShortenByID(shortenId)

	if err != nil {
		helpers.ErrorResponse(ctx, err, "Id shorten not found")
		return
	}

	updateShortenRequest := models.EditReqShorten{}

	ctx.ShouldBindJSON(&updateShortenRequest)

	err = validate.Struct(updateShortenRequest)

	if err != nil {
		helpers.ErrorResponse(ctx, err, "Error validate request")
		return
	}

	us, err := s.shortenRepository.IsShortenAlreadyTaken(updateShortenRequest.Shorten)
	if err != nil {
		helpers.ErrorResponse(ctx, err, "Failed to check shorten")
		return
	}

	if us {
		helpers.ErrorResponse(ctx, err, "Shorten already taken")
		return
	}

	shorten.Shorten = updateShortenRequest.Shorten
	shorten.Original = updateShortenRequest.Original
	shorten.UpdatedBy = userId.(string)
	shorten.UpdatedAt = time.Now()

	updatedShorten, err := s.shortenRepository.Update(*shorten)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, data.ResponseModel{
			Response:   http.StatusInternalServerError,
			Error:      err.Error(),
			AppID:      "api-curut-in",
			Controller: "ShortenController",
			Action:     "UpdateShorten",
			Result:     nil,
		})
		return
	}

	response := data.ResponseModel{
		Response:   http.StatusOK,
		Error:      "",
		AppID:      "api-curut-in",
		Controller: "ShortenController",
		Action:     "UpdateShorten",
		Result:     updatedShorten,
	}

	ctx.JSON(http.StatusOK, response)

}
