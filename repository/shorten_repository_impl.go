package repository

import (
	"api-curut-in/helpers"
	"api-curut-in/models"

	"gorm.io/gorm"
)

type ShortenRepositoryImpl struct {
	Db *gorm.DB
}

func NewShortenRepositoryImpl(Db *gorm.DB) ShortenRepository {
	return &ShortenRepositoryImpl{Db: Db}
}

func (s *ShortenRepositoryImpl) FindAll() (*[]models.Shorten, error) {
	var shortens []models.Shorten
	result := s.Db.Find(&shortens)
	helpers.ErrorPanic(result.Error)

	return &shortens, result.Error
}

func (s *ShortenRepositoryImpl) Save(shorten models.Shorten) (*models.Shorten, error) {
	result := s.Db.Create(&shorten)
	return &shorten, result.Error
}
