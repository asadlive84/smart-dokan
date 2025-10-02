package main

import (
	"log"
	"os"
	config "smart-dokan/usersvc/config"
	"smart-dokan/usersvc/internal/adapters/db"
	"smart-dokan/usersvc/internal/adapters/grpc"
	"smart-dokan/usersvc/internal/application/core/api"
)

func exec() {
	
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	APP_PORT := os.Getenv("APP_PORT")

	dbAdapter, err := db.NewAdapter(db.DBConfig{
		DB_HOST:     DB_HOST,
		DB_PORT:     DB_PORT,
		DB_USERNAME: DB_USERNAME,
		DB_NAME:     DB_NAME,
		DB_PASSWORD: DB_PASSWORD,
	})

	if err != nil {
		log.Fatalf("")
	}

	application := api.NewApplication(dbAdapter)

	grpcAdapter, err := grpc.NewAdapter(application, APP_PORT)

	if err != nil {
		log.Fatalf("grpc error: %+v", err)
	}

	grpcAdapter.Run()

}

func init() {
	config.SetConfig()
}
