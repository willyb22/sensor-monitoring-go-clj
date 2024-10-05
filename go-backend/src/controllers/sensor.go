package controllers

import (
	"go-backend/models"
	"github.com/gofiber/fiber/v2"
	"log"
)

func PostSensorHandler(c *fiber.Ctx) error {
	var sensorData models.SensorData
	if err := c.BodyParser(&sensorData); err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Failed to parse request body",
			})
	}
	log.Printf("Received: %+v\n", sensorData)
	return c.JSON(fiber.Map{
		"message": "Data received successfully",
		"data": sensorData,
	})
}