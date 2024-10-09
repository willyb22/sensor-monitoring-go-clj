package services

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"go-backend/config"
	"log"
)

type Service struct {
	DB *sql.DB
}

var SensorService *Service

func Connect(){
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.BackendConfig.DBHost,
		config.BackendConfig.DBUser,
		config.BackendConfig.DBPass,
		config.BackendConfig.DBName,
		config.BackendConfig.DBPort,
	)
	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err!=nil {
		log.Fatal("Failed to connect to the database.", err)
	}
	SensorService = new(Service)
	SensorService.DB = db

	// Check if connection is alive
	if err:=SensorService.DB.Ping(); err!=nil {
		log.Fatal("Failed to ping the database.", err)
	}

	log.Println("Successfully connected to the database")

}