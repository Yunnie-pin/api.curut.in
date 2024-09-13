package models

type Roles struct {
	ID   int    `gorm:"primary_key;type:uuid;" json:"id"`
	Name string `gorm:"type:varchar(225);not null" json:"name"`
}

func (Roles) TableName() string {
	return "roles"
}
