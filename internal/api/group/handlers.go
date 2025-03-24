package group

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

func GetGroup(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("GET /group with id:", id)
	return c.Status(fiber.StatusOK).JSON(Group{ID: id})
}

func CreateGroup(c *fiber.Ctx) error {
	log.Println("POST /group")
	return c.Status(fiber.StatusCreated).JSON(Group{})
}

func UpdateGroup(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("PUT /group with id:", id)
	return c.Status(fiber.StatusOK).JSON(Group{ID: id})
}

func DeleteGroup(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("DELETE /group with id:", id)
	return c.Status(fiber.StatusNoContent).JSON("")
}
