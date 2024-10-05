package middleware

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

func Log(c *fiber.Ctx) error {
	log.Printf("Received %s request for %s\n", c.Method(), c.Path())
	return c.Next()
}