package handler

import (
	v1 "github.com/Lapp-coder/file-service/internal/handler/v1"
	"github.com/Lapp-coder/file-service/internal/service"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	router  fiber.Router
	service service.Service
}

func New(router fiber.Router, service service.Service) Handler {
	return Handler{
		router:  router,
		service: service,
	}
}

func (h *Handler) Init() {
	api := h.router.Group("/api")
	{
		handlerV1 := v1.New(api, h.service)
		handlerV1.Init()
	}
}
