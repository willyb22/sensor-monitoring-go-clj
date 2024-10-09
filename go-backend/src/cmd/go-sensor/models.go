package main

import (
	"fmt"
	"time"
	"math/rand"
)

type BSSensorData struct {
	SensorName string `json:"sensor_name"`
	SensorType string `json:"sensor_type"`
	Temperature float32 `json:"temperature"`
	Humidity float32 `json:"humidity"`
	Timestamp time.Time `json:"timestamp"`
}

type AQSSensorData struct {
	SensorName string `json:"sensor_name"`
	SensorType string `json:"sensor_type"`
	Temperature float32 `json:"temperature"`
	Humidity float32 `json:"humidity"`
	CO2Level float32 `json:"co2_level"`
	Timestamp time.Time `json:"timestamp"`
}

type MSISensorData struct {
	SensorName string `json:"sensor_name"`
	SensorType string `json:"sensor_type"`
	Temperature float32 `json:"temperature"`
	Humidity float32 `json:"humidity"`
	AirPressure float32 `json:"air_pressure"`
	WindSpeed float32 `json:"wind_speed"`
	Timestamp time.Time `json:"timestamp"`
}

type SensorData interface {
	GenerateData(i int)
}

func (data *BSSensorData) GenerateData(i int) {
	data.SensorType = "bs"
	data.SensorName = data.SensorType + fmt.Sprintf("%03d", i)
	data.Temperature = 25 + rand.Float32()*15 + 20
	data.Humidity = 20 + rand.Float32()*60
	data.Timestamp = time.Now()
}

func (data *AQSSensorData) GenerateData(i int) {
	data.SensorType = "aqs"
	data.SensorName = data.SensorType + fmt.Sprintf("%03d", i)
	data.Temperature = 25 + rand.Float32()*15 + 20
	data.Humidity = 20 + rand.Float32()*60
	data.CO2Level = 400 + 600*rand.Float32()
	data.Timestamp = time.Now()
}

func (data *MSISensorData) GenerateData(i int) {
	data.SensorType = "msi"
	data.SensorName = data.SensorType + fmt.Sprintf("%03d", i)
	data.Temperature = 25 + rand.Float32()*15 + 20
	data.Humidity = 20 + rand.Float32()*60
	data.AirPressure = 900 + 200*rand.Float32()
	data.WindSpeed = 0.1 + 5*rand.Float32()
	data.Timestamp = time.Now()
}