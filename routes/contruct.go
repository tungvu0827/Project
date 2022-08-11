package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/sixfwa/fiber-api/database"
	"github.com/sixfwa/fiber-api/models"
)

type Contruct struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Body   string `json:"body"`
	IDpage string `json: "idpage"`
}

func CreateResponseContruct(contruct models.Contruct) Contruct {
	return Contruct{ID: contruct.ID, Body: contruct.Body, IDpage: contruct.IDpage}
}

func CreateContruct(c *fiber.Ctx) error {
	var contruct models.Contruct

	if err := c.BodyParser(&contruct); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&contruct)
	responseContruct := CreateResponseContruct(contruct)

	return c.Status(200).JSON(responseContruct)
}

func GetContructs(c *fiber.Ctx) error {
	contructs := []models.Contruct{}

	database.Database.Db.Find(&contructs)

	responseContructs := []Contruct{}

	for _, contruct := range contructs {
		responseContruct := CreateResponseContruct(contruct)
		responseContructs = append(responseContructs, responseContruct)
	}

	return c.Status(200).JSON(responseContructs)
}

func findContruct(id int, contruct *models.Contruct) error {
	database.Database.Db.Find(&contruct, "id = ?", id)
	if contruct.ID == 0 {
		return errors.New("user does not exis")

	}
	return nil
}
func GetContruct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var contruct models.Contruct

	if err != nil {
		return c.Status(400).JSON("Ensure that :id ")
	}
	if err := findContruct(id, &contruct); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseContruct := CreateResponseContruct(contruct)

	return c.Status(200).JSON(responseContruct)
}

func UpdateContruct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var contruct models.Contruct

	if err != nil {
		return c.Status(400).JSON("Ensure that :id ")
	}
	if err := findContruct(id, &contruct); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type UpdateContruct struct {
		Body   string `json:"body"`
		IDpage string `json:"idpage"`
	}

	var updateData UpdateContruct

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	contruct.Body = updateData.Body
	contruct.IDpage = updateData.IDpage

	database.Database.Db.Save(&contruct)

	responseContruct := CreateResponseContruct(contruct)
	return c.Status(200).JSON(responseContruct)
}

func DeleteContruct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var contruct models.Contruct

	if err != nil {
		return c.Status(400).JSON("Ensure that :id ")
	}
	if err := findContruct(id, &contruct); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := database.Database.Db.Delete(&contruct).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).SendString("Deleted Page")
}
