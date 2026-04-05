package main

import (
	"book-shop/internal/app"
	"book-shop/internal/config"
	"book-shop/internal/platform/db"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	var databaseUrl string = cfg.DatabaseURL

	database, err := db.NewPostgres(databaseUrl) //* function from db.go for creating a new database connection
	if err != nil {
		log.Fatal(err)
	}

	migrateRunError := db.RunMigrate(database) //* function from db.go for running database migrations

	if migrateRunError != nil {
		log.Fatal(migrateRunError)
	}

	application := app.NewApp(cfg, database) //* function from app.go for creating a new application instance

	log.Printf("server is running on port %s", cfg.Port)

	err = application.Run()
	if err != nil {
		log.Fatal(err)
	}
}
