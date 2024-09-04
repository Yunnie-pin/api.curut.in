package repository

import (
	"api-curut-in/models"
)

type ShortenRepository interface {
	FindAll() (*[]models.Shorten, error)
	// FindWithPagination(pagination data.Pagination) (data.Pagination, error)
	Save(shorten models.Shorten) (*models.Shorten, error)
	// FindShowedArticles() (*[]models.Shorten, error)
	// FindByID(id string) (*models.Shorten, error)
	FindShortenByName(name string) (*models.Shorten, error)
	IsShortenAlreadyTaken(shorten string) (bool, error)
}
