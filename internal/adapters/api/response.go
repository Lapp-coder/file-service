package api

import "github.com/gofiber/fiber/v2"

func RespondJSON(ctx *fiber.Ctx, statusCode int, data interface{}) error {
	ctx.Response().SetStatusCode(statusCode)
	if err, ok := data.(error); ok {
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(data)
}
