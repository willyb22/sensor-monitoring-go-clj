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

var (
	GO_BACKEND_URL string = os.Getenv("GO_BACKEND_URL")
	mu sync.Mutex
)

func main () {
	pingBackend(GO_BACKEND_URL)
	generateSensorData()
}

// Check if the server is ready
func pingBackend (url string) {
	for {
		resp, err := http.Get(url)
		// log.Printf("%#v\n", resp)
		if err == nil && resp.StatusCode == http.StatusOK {
			log.Println("Server is reachable. Proceeding with data generation...")
			return
		}
		log.Println("Waiting for the server to be ready...")
		time.Sleep(5 * time.Second) // Wait before retrying
	}
}

// POST data to the backend
func postDataToBackend(Data interface{}) error {
	mu.Lock()
	jsonData, err := json.Marshal(Data)
	if err != nil {
		log.Println("Error marshaling sensor data:", err)
	}
	defer mu.Unlock()
	log.Printf("Sending data: %+v \n", string(jsonData))
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
func generateSensorData() {
	rand.Seed(0)
	log.Println("inside generateSensorData")
	go generateBSSensorData()
	go generateAQSSensorData()
	go generateMSISensorData()

	select {}
}
//
func generateBSSensorData() {
	log.Println("inside generateBSSensorData")
	var numbers [30]int
	for i:=0; i<30; i++ {
		numbers[i] = i
	}
	for {
		n := 5+rand.Intn(6)
		rand.Shuffle(len(numbers), func(i, j int) {
			numbers[i], numbers[j] = numbers[j], numbers[i]
		})
		for i:=0; i<n; i++ {
			go func(j int){
				var (
					bs *BSSensorData
					sd SensorData
				)
				bs = new(BSSensorData)
				sd = bs
				sd.GenerateData(j)
				if err:=postDataToBackend(bs); err!=nil{
					log.Println("Error")
				}
			}(numbers[i])
		}
		time.Sleep(100*time.Second)
	}
}

func generateAQSSensorData() {
	log.Println("inside generateAQSSensorData")
	var numbers [10]int
	for i:=0; i<10; i++ {
		numbers[i] = i
	}
	for {
		n := 5+rand.Intn(6)
		rand.Shuffle(len(numbers), func(i, j int) {
			numbers[i], numbers[j] = numbers[j], numbers[i]
		})
		for i:=0; i<n; i++ {
			go func(j int){
				var (
					aqs *AQSSensorData
					sd SensorData
				)
				aqs = new(AQSSensorData)
				sd = aqs
				sd.GenerateData(j)
				if err:=postDataToBackend(aqs); err!=nil{
					log.Println("Error")
				}
			}(numbers[i])
		}
		time.Sleep(150*time.Second)
	}
}

func generateMSISensorData() {
	log.Println("inside generateMSISensorData")
	var numbers [5]int
	for i:=0; i<5; i++ {
		numbers[i] = i
	}
	for {
		n := 1+rand.Intn(5)
		rand.Shuffle(len(numbers), func(i, j int) {
			numbers[i], numbers[j] = numbers[j], numbers[i]
		})
		for i:=0; i<n; i++ {
			go func(j int){
				var (
					msi *MSISensorData
					sd SensorData
				)
				msi = new(MSISensorData)
				sd = msi
				sd.GenerateData(j)
				if err:=postDataToBackend(msi); err!=nil{
					log.Println("Error")
				}
			}(numbers[i])
		}
		time.Sleep(200*time.Second)
	}
}
