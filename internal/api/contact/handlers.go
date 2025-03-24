package contact

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

func GetContact(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("GET /contact with id:", id)
	return c.Status(fiber.StatusOK).JSON(Contact{ID: id})
}

func CreateContact(c *fiber.Ctx) error {
	log.Println("POST /contact")
	return c.Status(fiber.StatusCreated).JSON(Contact{})
}

func UpdateContact(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("PUT /contact with id:", id)
	return c.Status(fiber.StatusOK).JSON(Contact{ID: id})
}

func DeleteContact(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("DELETE /contact with id:", id)
	return c.Status(fiber.StatusNoContent).JSON("")
}
