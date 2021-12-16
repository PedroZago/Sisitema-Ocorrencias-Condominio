package controllers

import (
	"github.com/PedroZago/condominitech/configs"
	"github.com/PedroZago/condominitech/database"
	"github.com/PedroZago/condominitech/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func CadastrarTipoAnexo(c *fiber.Ctx) error {
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

	var tipoAnexo models.TipoAnexo

	if err := c.BodyParser(&tipoAnexo); err != nil {
		return err
	}

	err = database.DB.Create(&tipoAnexo).Error
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "näo foi possível criar o tipo de anexo: " + err.Error(),
		})
	}

	return c.JSON(tipoAnexo)
}

func BuscarTipoAnexoID(c *fiber.Ctx) error {
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

	var tipoAnexo models.TipoAnexo

	database.DB.Where("id = ?", id).First(&tipoAnexo)

	if tipoAnexo.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "tipo de anexo não encontrada.",
		})
	}

	return c.JSON(tipoAnexo)
}

func BuscarTodosTipoAnexos(c *fiber.Ctx) error {
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

	var tipoAnexos []models.TipoAnexo

	database.DB.Find(&tipoAnexos)

	return c.JSON(tipoAnexos)
}

func DeletarTipoAnexo(c *fiber.Ctx) error {
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

	var tipoAnexo models.TipoAnexo

	database.DB.Where("id = ?", id).First(&tipoAnexo)

	if tipoAnexo.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "tipo de anexo não encontrado.",
		})
	}

	database.DB.Delete(&models.TipoAnexo{}, tipoAnexo.ID)

	return c.JSON(fiber.Map{
		"message": "tipo de anexo excluído com sucesso.",
	})
}

func AtualizarTipoAnexo(c *fiber.Ctx) error {
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

	var tipoAnexoJSON models.TipoAnexo

	if err := c.BodyParser(&tipoAnexoJSON); err != nil {
		return err
	}

	id := c.Params("id")

	var tipoAnexo models.TipoAnexo

	database.DB.Where("id = ?", id).First(&tipoAnexo)

	if tipoAnexo.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "tipo de anexo não encontrado.",
		})
	}

	database.DB.Model(&tipoAnexo).Select("descricao", "tipo").Updates(
		models.TipoAnexo{
			Descricao: tipoAnexoJSON.Descricao,
			Tipo:      tipoAnexoJSON.Tipo,
		})

	return c.JSON(tipoAnexo)
}
