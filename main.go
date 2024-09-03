package main

import (
	"api-curut-in/config"
	"api-curut-in/controllers"
	"api-curut-in/services"
	"flag"
	"log"

	"api-curut-in/migration"

	"api-curut-in/repository"
	"api-curut-in/router"
)

func main() {
	loadConfig, _ := config.LoadConfig()

	//Init Database
	db := config.StartDB(&loadConfig)

	// Definisikan flag --migrate
	input := flag.Bool("migrate", false, "Run migration scripts")

	// Parsing flag
	flag.Parse()
	if *input {
		//Migration
		err := migration.MigrationDB(db)
		if err != nil {
			log.Fatal("🧊 Tidak bisa melakukan migrasi database. ", err)
		}
		return
	}

	//Init Repository
	userRepository := repository.NewUsersRepositoryImpl(db)
	shortenRepository := repository.NewShortenRepositoryImpl(db)

	//init Service
	authService := services.NewAuthenticationServiceImpl(userRepository)

	//Init Controller
	userController := controllers.NewUsersController(userRepository)
	authController := controllers.NewAuthenticationController(authService)
	shortenController := controllers.NewShortenController(shortenRepository)

	//Init Router
	r := router.NewRouter(
		userRepository,
		userController,
		authController,
		shortenController,
	)

	log.Println("🧊 ENV: ", loadConfig.Env)
	log.Println("🧊 Menjalankan di port :", loadConfig.Port)
	r.Run(":" + loadConfig.Port)
}
