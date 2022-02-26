package api

import "github.com/gofiber/fiber/v2"

type Handler interface {
	Init(router fiber.Router)
}
