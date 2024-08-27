package migration

import (
	models "api-curut-in/migration/migration_model"
	"api-curut-in/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func MigrationDB(db *gorm.DB) (err error) {
	// //drop table if exists
	db.Migrator().DropTable(&models.Roles{})
	db.Migrator().DropTable(&models.Users{})

	//create table
	db.Table("roles").AutoMigrate(&models.Roles{})
	db.Table("users").AutoMigrate(&models.Users{})

	//create multiple records with faker
	SendSeeder(db, 10)

	return
}

func SendSeeder(db *gorm.DB, total int) (err error) {
	//create multiple records with faker

	roles := []models.Roles{
		{
			ID:   1,
			Name: "admin",
		},
		{
			ID:   2,
			Name: "user",
		},
	}
	db.Create(&roles)

	hashedPassword, _ := utils.HashPassword("12345678")

	uidUser := []uuid.UUID{
		uuid.New(),
		uuid.New(),
	}

	users := []models.Users{
		{
			ID:        uidUser[0],
			Username:  "admin1234",
			Password:  hashedPassword,
			Email:     "admin@admin.com",
			RolesID:   1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        uidUser[1],
			Username:  "user1234",
			Password:  hashedPassword,
			Email:     "user@user.com",
			RolesID:   2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	db.Create(&users)

	return
}
