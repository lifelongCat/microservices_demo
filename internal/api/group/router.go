package group

import (
	"github.com/gofiber/fiber/v2"
)

func CreateRoutes(api fiber.Router) {
	apiGroup := api.Group("")

	apiGroup.Get("/:id", GetGroup)
	apiGroup.Post("/", CreateGroup)
	apiGroup.Put("/:id", UpdateGroup)
	apiGroup.Delete("/:id", DeleteGroup)
}
