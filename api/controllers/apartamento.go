package controllers

import (
	"github.com/PedroZago/condominitech/configs"
	"github.com/PedroZago/condominitech/database"
	"github.com/PedroZago/condominitech/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func CadastrarApartamento(c *fiber.Ctx) error {
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

	var apartamento models.Apartamento

	if err := c.BodyParser(&apartamento); err != nil {
		return err
	}

	err = database.DB.Create(&apartamento).Error
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "näo foi possível criar o apartamento: " + err.Error(),
		})
	}

	return c.JSON(apartamento)
}

func BuscarApartamentoID(c *fiber.Ctx) error {
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

	var apartamento models.Apartamento

	database.DB.Where("id = ?", id).First(&apartamento)

	if apartamento.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "apartamento não encontrado.",
		})
	}

	return c.JSON(apartamento)
}

func BuscarTodosApartamentos(c *fiber.Ctx) error {
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

	var apartamentos []models.Apartamento

	database.DB.Find(&apartamentos)

	return c.JSON(apartamentos)
}

func DeletarApartamento(c *fiber.Ctx) error {
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

	var apartamento models.Apartamento

	database.DB.Where("id = ?", id).First(&apartamento)

	if apartamento.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "apartamento não encontrado.",
		})
	}

	database.DB.Delete(&models.Apartamento{}, apartamento.ID)

	return c.JSON(fiber.Map{
		"message": "apartamento excluído com sucesso.",
	})
}

func AtualizarApartamento(c *fiber.Ctx) error {
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

	var apartamentoJSON models.Apartamento

	if err := c.BodyParser(&apartamentoJSON); err != nil {
		return err
	}

	id := c.Params("id")

	var apartamento models.Apartamento

	database.DB.Where("id = ?", id).First(&apartamento)

	if apartamento.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "apartamento não encontrado.",
		})
	}

	database.DB.Model(&apartamento).Select("identificador", "descricao").Updates(
		models.Apartamento{
			Identificador: apartamentoJSON.Identificador,
			Descricao:     apartamentoJSON.Descricao,
		})

	return c.JSON(apartamento)
}
