package repository

import (
	"api-curut-in/helpers"
	"api-curut-in/models"
	"errors"

	"gorm.io/gorm"
)

type ShortenRepositoryImpl struct {
	Db *gorm.DB
}

func NewShortenRepositoryImpl(Db *gorm.DB) ShortenRepository {
	return &ShortenRepositoryImpl{Db: Db}
}

func (s *ShortenRepositoryImpl) FindShortenByName(name string) (*models.Shorten, error) {
	var shorten models.Shorten
	result := s.Db.Where("shorten = ?", name).First(&shorten)
	if result.Error != nil {
		return &shorten, errors.New("NOT FOUND")
	}

	return &shorten, nil
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

func (s *ShortenRepositoryImpl) IsShortenAlreadyTaken(shorten string) (bool, error) {
	var count int64
	result := s.Db.Model(&models.Shorten{}).Where("shorten = ?", shorten).Count(&count)
	helpers.ErrorPanic(result.Error)

	return count > 0, result.Error
}

func (s *ShortenRepositoryImpl) FindShortenByID(id string) (*models.Shorten, error) {
	var shorten models.Shorten
	result := s.Db.Where("id = ?", id).First(&shorten)
	if result.Error != nil {
		return &shorten, errors.New("NOT FOUND")
	}

	return &shorten, nil
}

func (s *ShortenRepositoryImpl) Update(shorten models.Shorten) (*models.Shorten, error) {
	err := s.Db.Model(&shorten).Updates(shorten).Error
	if err != nil {
		return nil, err
	}
	return &shorten, nil
}
