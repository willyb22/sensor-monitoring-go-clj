package main

import (
	"fmt"
	"log"
	"os"
	"math/rand"
	"time"
	"encoding/json"
	"net/http"
	"sync"
	"bytes"
)

type SensorData struct {
	SensorName string `json:"sensorname"`
	Temperature float32 `json:"temperature"`
	Humidity float32 `json:"humidity"`
	Timestamp time.Time `json:"timestamp"`
}

var (
	GO_BACKEND_URL string = "http://go-backend:5000" // os.Getenv("GO_BACKEND_URL")
	GO_SENSOR_URL_WITHOUT_PROTOCOL string = os.Getenv("GO_SENSOR_URL_WITHOUT_PROTOCOL")

	Data SensorData

	mu sync.Mutex
)


func main () {
	// log.SetOutput()
	serverReady := make(chan struct{}) 
	go func(url string) {
		for {
			resp, err := http.Get(url)
			// log.Printf("%#v\n", resp)
			if err == nil && resp.StatusCode == http.StatusOK {
				log.Println("Server is reachable. Proceeding with data generation...")
				close(serverReady)               // Notify that the server is ready
				return
			}
			log.Println("Waiting for the server to be ready...")
			time.Sleep(5 * time.Second) // Wait before retrying
		}
	}(GO_BACKEND_URL) // "http://" + GO_SENSOR_URL_WITHOUT_PROTOCOL

	// Generate Data
	go GenerateSensorData(serverReady, &Data)
	
	select {}

}
// POST data to the backend
func postDataToBackend() error {
	
	mu.Lock()
	jsonData, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error marshaling sensor data:", err)
	}
	defer mu.Unlock()

	resp, err := http.Post(GO_BACKEND_URL+"/sensor", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to post data: %w", err)
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-OK status code: %d", resp.StatusCode)
	}

	return nil
}

// Generate Data
func GenerateSensorData(ready chan struct{}, data *SensorData) {
	<-ready
	log.Println("Generating Data ...")
	// Begin
	for {
		mu.Lock() // Lock before updating the data
		*data = SensorData{
			SensorName: fmt.Sprintf("sensor-%d", rand.Intn(100)),
			Temperature: rand.Float32()*15 + 20,
			Humidity: rand.Float32()*100,
			Timestamp: time.Now(),
		}
		mu.Unlock() // Unlock after updating the data

		err := postDataToBackend()
		if err != nil {
			log.Println("Error posting data to backend:", err)
		} else {
			log.Println("Successfully posted data to backend:", data)
		}

		time.Sleep(10*time.Second)
	}
}