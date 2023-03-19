package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hatienl0i261299/fiber_gorm_postgresql/database"
	"github.com/hatienl0i261299/fiber_gorm_postgresql/models"
)

func CreateArticle(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	article := models.Article{
		Title:       data["title"],
		Description: data["description"],
	}

	if err := database.DB.Create(&article).Error; err != nil {
		return err
	}

	return c.JSON(article)
}

func GetAllArticle(c *fiber.Ctx) error {

	var articles []models.Article

	database.DB.Model(&models.Article{}).Find(&articles)

	return c.JSON(articles)
}
