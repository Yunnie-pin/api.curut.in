package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

type UserResponse struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;" json:"id"`
	Username  string    `gorm:"type:varchar(225);not null;uniqueIndex" json:"username"`
	Email     string    `gorm:"type:varchar(225);not null;uniqueIndex" json:"email"`
	RolesID   int       `gorm:"not null;references:ID" json:"roles_id"`
	Roles     Roles     `gorm:"foreignKey:RolesID;references:ID;" json:"roles"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type UserSubResponse struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;" json:"id"`
	Username  string    `gorm:"type:varchar(225);not null;uniqueIndex" json:"username"`
	Email     string    `gorm:"type:varchar(225);not null;uniqueIndex" json:"email"`
	RolesID   int       `gorm:"not null;references:ID" json:"roles_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type UpdateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	RoleID   int    `json:"role_id" validate:"required"`
}

func (UserResponse) TableName() string {
	return "users"
}

func (UserSubResponse) TableName() string {
	return "users"
}

// BeforeCreate hook to set the primary key to a new UUID
func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
