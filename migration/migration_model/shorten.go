package migration_model

import (
	"time"

	"github.com/google/uuid"
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
