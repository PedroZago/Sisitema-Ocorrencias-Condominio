package controllers

import (
	"github.com/PedroZago/condominitech/configs"
	"github.com/PedroZago/condominitech/database"
	"github.com/PedroZago/condominitech/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func CadastrarBloco(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	_, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "usuário não autenticado.",
		})
	}

	var bloco models.Bloco

	if err := c.BodyParser(&bloco); err != nil {
		return err
	}

	err = database.DB.Create(&bloco).Error
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "näo foi possível criar o bloco: " + err.Error(),
		})
	}

	return c.JSON(bloco)
}

func BuscarBlocoID(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	_, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "usuário não autenticado.",
		})
	}

	id := c.Params("id")

	var bloco models.Bloco

	database.DB.Where("id = ?", id).First(&bloco)

	if bloco.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "bloco não encontrado.",
		})
	}

	return c.JSON(bloco)
}

func BuscarTodosBlocos(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	_, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "usuário não autenticado.",
		})
	}

	var blocos []models.Bloco

	database.DB.Find(&blocos)

	return c.JSON(blocos)
}

func DeletarBloco(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	_, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "usuário não autenticado.",
		})
	}

	id := c.Params("id")

	var bloco models.Bloco

	database.DB.Where("id = ?", id).First(&bloco)

	if bloco.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "bloco não encontrado.",
		})
	}

	database.DB.Delete(&models.Bloco{}, bloco.ID)

	return c.JSON(fiber.Map{
		"message": "bloco excluído com sucesso.",
	})
}

func AtualizarBloco(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	_, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "usuário não autenticado.",
		})
	}

	var blocoJSON models.Bloco

	if err := c.BodyParser(&blocoJSON); err != nil {
		return err
	}

	id := c.Params("id")

	var bloco models.Bloco

	database.DB.Where("id = ?", id).First(&bloco)

	if bloco.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "bloco não encontrado.",
		})
	}

	database.DB.Model(&bloco).Select("identificador", "descricao").Updates(
		models.Bloco{
			Identificador: blocoJSON.Identificador,
			Descricao:     blocoJSON.Descricao,
		})

	return c.JSON(bloco)
}
