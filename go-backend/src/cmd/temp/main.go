// Temporary testing (go run -mod=vendor cmd/temp/main.go)
package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"go-backend/config"
	"go-backend/models"
	"go-backend/services"
	"time"
	"log"
)

// type Service struct {
// 	DB *sql.DB
// }

var SensorService *services.Service

func main(){
	config.LoadConfig()
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.BackendConfig.DBHost,
		config.BackendConfig.DBUser,
		config.BackendConfig.DBPass,
		config.BackendConfig.DBName,
		config.BackendConfig.DBPort,
	)
	fmt.Println(connStr)
	fmt.Printf("%#v\n",config.BackendConfig)
	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err!=nil {
		log.Fatal("Failed to connect to the database.", err)
	}
	defer db.Close()
	SensorService = new(services.Service)
	SensorService.DB = db

	// Check if connection is alive
	if err:=SensorService.DB.Ping(); err!=nil {
		log.Fatal("Failed to ping the database.", err)
	}

	log.Println("Successfully connected to the database")
	var bs interface{}
	bs = &models.BSSensorData{
		SensorName: "bs001",
		SensorType: "bs",
		Temperature: 20.5,
		Humidity: 20.5,
		Timestamp: time.Now(),
	}
	sd, ok := bs.(models.SensorData)
	if !ok {
		fmt.Println("assertion failed")
	}
	
	if err := sd.InsertToDB(SensorService); err != nil {
		log.Printf("Error inserting data into the database: %v", err) // Log specific error
	} else {
		log.Printf("Data ")
	}
	
}
