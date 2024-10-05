package models

import (
	"time"
	"errors"
)

type SensorData struct {
	SensorName string `json:"sensorname"`
	Temperature float32 `json:"temperature"`
	Humidity float32 `json:"humidity"`
	Timestamp time.Time `json:"timestamp"`
}

func (s *SensorData) Validate() error {
	if s.SensorName==""{
		return errors.New("sensorname cannot be empty!")
	}
	return nil
}