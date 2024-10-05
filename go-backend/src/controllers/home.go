package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetHome(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Hello From Backend")
}