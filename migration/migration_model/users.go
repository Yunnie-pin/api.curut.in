package migration_model

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;" json:"id"`
	Username  string    `gorm:"type:varchar(225);not null;uniqueIndex" json:"username"`
	Password  string    `gorm:"type:varchar(225);not null" json:"password"`
	Email     string    `gorm:"type:varchar(225);not null;uniqueIndex" json:"email"`
	RolesID   int       `gorm:"not null;references:ID" json:"roles_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

