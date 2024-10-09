package controllers

import (
	"go-backend/models"
	"github.com/gofiber/fiber/v2"
	"go-backend/services"
	"log"
	_ "reflect"
)

func PostSensorHandler(c *fiber.Ctx) error {
    var rawData map[string]interface{}
    log.Println("PostSensorHandler Triggered")

    // Parse the JSON
    if err := c.BodyParser(&rawData); err != nil {
        log.Printf("Error parsing request body: %v", err) // Log specific error
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Failed to parse request body",
        })
    }

    // Check for sensor_type key
    sensorTypeInterface, ok := rawData["sensor_type"]
    if !ok {
        log.Printf("Key 'sensor_type' not found in request body: %+v\n", rawData)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Request body doesn't have sensor_type key",
        })
    }

    // Assertion for sensor_type
    sensorType, ok := sensorTypeInterface.(string)
    if !ok {
        log.Printf("Value for 'sensor_type' is not a string: %+v\n", sensorTypeInterface)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Request body's sensor_type's value is not a string",
        })
    }

    log.Printf("sensor_type = %+v \n", sensorType)
	var sensorDataPtr interface{}
	if sensorType=="bs" {
		sensorDataPtr = &models.BSSensorData{}
	} else if sensorType=="aqs" {
		sensorDataPtr = &models.AQSSensorData{}
	} else if sensorType=="msi" {
		sensorDataPtr = &models.MSISensorData{}
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Request body's sensor_type's value is not available",
        })
	}

    // // Generate available model pointer and check for sensor data pointer
    // mapWithPtr := models.GenerateAvailableModelPointer()
    // sensorDataPtr, ok := mapWithPtr[sensorType]
    // if !ok {
    //     log.Printf("sensor_type '%s' is not available in models: %+v\n", sensorType, mapWithPtr)
    //     return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
    //         "error": "Request body's sensor_type's value is not available",
    //     })
    // }

    // Parse sensor data ptr to the actual data type
    if err := c.BodyParser(sensorDataPtr); err != nil {
        log.Printf("Error parsing additional sensor data: %v", err) // Log specific error
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Failed to parse request body for sensor data",
        })
    }

    log.Printf("Received sensor data: %+v\n", sensorDataPtr)

	// Using interface to the pointer
	// var sensorData models.SensorData
	sensorData, ok := sensorDataPtr.(models.SensorData)
	if !ok {
		log.Printf("Error Assertion")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to assert sensorDataPtr",
		}) 
	}
	// Insert data to database
	if err := sensorData.InsertToDB(services.SensorService); err != nil {
		log.Printf("Error inserting data into the database: %v", err) // Log specific error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert data to database",
		})
	}


    // // Assertion for sensor data
	// value := reflect.ValueOf(sensorDataPtr)
	// if value.Kind() == reflect.Ptr {
	// 	element := value.Elem()

	// 	if sensorData, ok := element.Interface().(models.SensorData); ok {
	// 		// Insert Data to Database
			// if err := sensorData.InsertToDB(services.SensorService); err != nil {
			// 	log.Printf("Error inserting data into the database: %v", err) // Log specific error
			// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			// 		"error": "Failed to insert data to database",
			// 	})
			// }
	// 	} else {
	// 		log.Printf("element type: %s\n", element.Type())
	// 		log.Printf("element value: %#v\n", element.Interface())
	// 		log.Printf("--- Type assertion failed: sensorDataPtr doesn't implement SensorData: %#v\n", element)
    //     	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 			"error": "Type Assertion failed: sensorDataPtr doesn't implement SensorData",
	// 		})
	// 	}
	// } else {
	// 	log.Printf("The reflect value is not a pointer: %+v", value.Kind()) // Log specific error
    //     return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
    //         "error": "Type Assertion failed",
    //     })
	// }

    return c.JSON(fiber.Map{
        "message": "Data inserted successfully",
    })
}
