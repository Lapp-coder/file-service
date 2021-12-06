package v1

import (
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
	v1 := h.router.Group("/v1")
	{
		files := v1.Group("/files")
		{
			files.Post("/", h.uploadFile)
			files.Get("/:uuid", h.getFileByUUID)
			files.Get("/:uuid/statistics", h.getFileStatisticByUUID)
		}
	}
}
