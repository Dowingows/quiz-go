package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/dowingows/quiz-go/database"
	"github.com/dowingows/quiz-go/models"
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Div Rhino Trivia App!")
}

func Create(c *fiber.Ctx) error {

	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	json, err := json.Marshal(fact)
	if err != nil {
		fmt.Println(err)
	}

	database.Redis.Client.Set(fmt.Sprintf("fact-3"), json, 0).Err()

	return c.Status(200).JSON(fact)
}

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	return c.Status(200).JSON(facts)
}

func ImportFact(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	return c.Status(200).JSON(facts)
}

func GetFact(c *fiber.Ctx) error {

	data, err := database.Redis.Client.Get("fact-3").Result()
	if err != nil {
		fmt.Println(err)
	}

	jsonData := []byte(data)
	var fact models.Fact
	if err := json.Unmarshal(jsonData, &fact); err != nil {
		fmt.Println("failed to unmarshal:", err)
	}

	return c.Status(200).JSON(fact)
}
