package migration

import (
	"api-curut-in/helpers"
	models "api-curut-in/migration/migration_model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func MigrationDB(db *gorm.DB) (err error) {
	// //drop table if exists
	db.Migrator().DropTable(&models.Roles{})
	db.Migrator().DropTable(&models.Users{})
	db.Migrator().DropTable(&models.Shorten{})

	//create table
	db.Table("roles").AutoMigrate(&models.Roles{})
	db.Table("users").AutoMigrate(&models.Users{})
	db.Table("shortens").AutoMigrate(&models.Shorten{})

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

	hashedPassword, _ := helpers.HashPassword("12345678")

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

	exampleShorten := []models.Shorten{
		{
			ID:        uuid.New(),
			Shorten:   "google",
			Original:  "https://www.google.com",
			Flag:      true,
			CreatedAt: time.Now(),
			CreatedBy: string(uidUser[0].String()),
			UpdatedAt: time.Now(),
			UpdatedBy: string(uidUser[0].String()),
		},
	}

	db.Create(&exampleShorten)

	return
}
