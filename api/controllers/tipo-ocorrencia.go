package controllers

import (
	"github.com/PedroZago/condominitech/configs"
	"github.com/PedroZago/condominitech/database"
	"github.com/PedroZago/condominitech/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func CadastrarTipoOcorrencia(c *fiber.Ctx) error {
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

	var tipoOcorrencia models.TipoOcorrencia

	if err := c.BodyParser(&tipoOcorrencia); err != nil {
		return err
	}

	err = database.DB.Create(&tipoOcorrencia).Error
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "näo foi possível criar o tipo de ocorrência: " + err.Error(),
		})
	}

	return c.JSON(tipoOcorrencia)
}

func BuscarTipoOcorrenciaID(c *fiber.Ctx) error {
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

	var tipoOcorrencia models.TipoOcorrencia

	database.DB.Where("id = ?", id).First(&tipoOcorrencia)

	if tipoOcorrencia.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "tipo de ocorrência não encontrada.",
		})
	}

	return c.JSON(tipoOcorrencia)
}

func BuscarTodosTipoOcorrencias(c *fiber.Ctx) error {
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

	var tipoOcorrencias []models.TipoOcorrencia

	database.DB.Find(&tipoOcorrencias)

	return c.JSON(tipoOcorrencias)
}

func DeletarTipoOcorrencia(c *fiber.Ctx) error {
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

	var tipoOcorrencia models.TipoOcorrencia

	database.DB.Where("id = ?", id).First(&tipoOcorrencia)

	if tipoOcorrencia.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "tipo de ocorrência não encontrado.",
		})
	}

	database.DB.Delete(&models.TipoOcorrencia{}, tipoOcorrencia.ID)

	return c.JSON(fiber.Map{
		"message": "tipo de ocorrência excluído com sucesso.",
	})
}

func AtualizarTipoOcorrencia(c *fiber.Ctx) error {
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

	var tipoOcorrenciaJSON models.TipoOcorrencia

	if err := c.BodyParser(&tipoOcorrenciaJSON); err != nil {
		return err
	}

	id := c.Params("id")

	var tipoOcorrencia models.TipoOcorrencia

	database.DB.Where("id = ?", id).First(&tipoOcorrencia)

	if tipoOcorrencia.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "tipo de ocorrência não encontrado.",
		})
	}

	database.DB.Model(&tipoOcorrencia).Select("descricao").Updates(
		models.TipoOcorrencia{
			Descricao: tipoOcorrenciaJSON.Descricao,
		})

	return c.JSON(tipoOcorrencia)
}
