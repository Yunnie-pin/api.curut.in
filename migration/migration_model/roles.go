package migration_model

type Roles struct {
	ID   int    `gorm:"primary_key;type:uuid;" json:"id"`
	Name string `gorm:"type:varchar(225);not null" json:"name"`
}
