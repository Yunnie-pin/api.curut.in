package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Shorten struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;" json:"id"`
	Shorten   string    `gorm:"type:varchar(225);not null" json:"shorten"`
	Original  string    `gorm:"type:varchar(225);not null" json:"original"`
	Flag      bool      `gorm:"default:true" json:"flag"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	CreatedBy string    `gorm:"type:varchar(225);not null" json:"created_by"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	UpdatedBy string    `gorm:"type:varchar(225);not null" json:"updated_by"`
}

type CreateShortenRequest struct {
	Shorten  string `json:"shorten" validate:"required"`
	Original string `json:"original" validate:"required"`
}

// BeforeCreate hook to set the primary key to a new UUID
func (u *Shorten) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}
