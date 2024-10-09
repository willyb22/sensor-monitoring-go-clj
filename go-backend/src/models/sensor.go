package models

import (
	"time"
	"log"
	_ "errors"
	"go-backend/services"
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
	// Validate(i int)
	InsertToDB(s *services.Service) error
}

func GenerateAvailableModelPointer() map[string]interface{} {
	return map[string]interface{}{
		"bs": new(BSSensorData),
		"aqs": new(AQSSensorData),
		"msi": new(MSISensorData),
	}
}

func (bs *BSSensorData) InsertToDB (s *services.Service) error {
	measurementsQuery := "INSERT INTO measurements(sensor_name,sensor_type,timestamp) VALUES ($1,$2,$3) RETURNING id"
	var measurementsID int
	err := s.DB.QueryRow(measurementsQuery, bs.SensorName, bs.SensorType, bs.Timestamp).Scan(&measurementsID); 
	if err!=nil {
		log.Fatal("Failed to insert into table measurements: ", err)
		return err
	}
	bsMeasurementsQuery := "INSERT INTO bs_measurements(id,temperature,humidity) VALUES ($1,$2,$3)"
	_, err = s.DB.Exec(
		bsMeasurementsQuery, 
		measurementsID, 
		bs.Temperature,
		bs.Humidity,
	)
	if err!=nil {
		log.Fatal("Failed to insert into table bs_measurements: ", err)
		return err
	}
	return nil
}

func (aqs *AQSSensorData) InsertToDB (s *services.Service) error {
	measurementsQuery := "INSERT INTO measurements(sensor_name,sensor_type,timestamp) VALUES ($1,$2,$3) RETURNING id"
	var measurementsID int
	err := s.DB.QueryRow(measurementsQuery, aqs.SensorName, aqs.SensorType, aqs.Timestamp).Scan(&measurementsID)
	if err!=nil {
		log.Fatal("Failed to insert into table measurements: ", err)
		return err
	}
	aqsMeasurementsQuery := "INSERT INTO aqs_measurements(id,temperature,humidity,co2_level) VALUES ($1,$2,$3,$4)"
	_, err = s.DB.Exec(
		aqsMeasurementsQuery, 
		measurementsID, 
		aqs.Temperature,
		aqs.Humidity,
		aqs.CO2Level,
	)
	if err!=nil {
		log.Fatal("Failed to insert into table aqs_measurements: ", err)
		return err
	}
	return nil
}

func (msi *MSISensorData) InsertToDB (s *services.Service) error {
	measurementsQuery := "INSERT INTO measurements(sensor_name,sensor_type,timestamp) VALUES ($1,$2,$3) RETURNING id"
	var measurementsID int
	err := s.DB.QueryRow(measurementsQuery, msi.SensorName, msi.SensorType, msi.Timestamp).Scan(&measurementsID)
	if err!=nil {
		log.Fatal("Failed to insert into table measurements: ", err)
		return err
	}
	msiMeasurementsQuery := "INSERT INTO msi_measurements(id,temperature,humidity,air_pressure,wind_speed)  VALUES ($1,$2,$3,$4,$5)"
	_, err = s.DB.Exec(
		msiMeasurementsQuery, 
		measurementsID, 
		msi.Temperature,
		msi.Humidity,
		msi.AirPressure,
		msi.WindSpeed,
	)
	if err!=nil {
		log.Fatal("Failed to insert into table aqs_measurements: ", err)
		return err
	}
	return nil
}