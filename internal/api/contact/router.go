package contact

import "github.com/gofiber/fiber/v2"

func CreateRoutes(api fiber.Router) {
	apiGroup := api.Group("")

	apiGroup.Get("/:id", GetContact)
	apiGroup.Post("/", CreateContact)
	apiGroup.Put("/:id", UpdateContact)
	apiGroup.Delete("/:id", DeleteContact)
}
