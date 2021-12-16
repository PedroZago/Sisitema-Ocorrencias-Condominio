package controllers

import (
	"github.com/PedroZago/condominitech/configs"
	"github.com/PedroZago/condominitech/database"
	"github.com/PedroZago/condominitech/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func CadastrarStatusOcorrencia(c *fiber.Ctx) error {
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

	var statusOcorrencia models.StatusOcorrencia

	if err := c.BodyParser(&statusOcorrencia); err != nil {
		return err
	}

	err = database.DB.Create(&statusOcorrencia).Error
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "näo foi possível criar o status de ocorrência: " + err.Error(),
		})
	}

	return c.JSON(statusOcorrencia)
}

func BuscarStatusOcorrenciaID(c *fiber.Ctx) error {
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

	var statusOcorrencia models.StatusOcorrencia

	database.DB.Where("id = ?", id).First(&statusOcorrencia)

	if statusOcorrencia.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "status de ocorrência não encontrada.",
		})
	}

	return c.JSON(statusOcorrencia)
}

func BuscarTodosStatusOcorrencias(c *fiber.Ctx) error {
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

	var statusOcorrencias []models.StatusOcorrencia

	database.DB.Find(&statusOcorrencias)

	return c.JSON(statusOcorrencias)
}

func DeletarStatusOcorrencia(c *fiber.Ctx) error {
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

	var statusOcorrencia models.StatusOcorrencia

	database.DB.Where("id = ?", id).First(&statusOcorrencia)

	if statusOcorrencia.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "status de ocorrência não encontrado.",
		})
	}

	database.DB.Delete(&models.StatusOcorrencia{}, statusOcorrencia.ID)

	return c.JSON(fiber.Map{
		"message": "status de ocorrência excluído com sucesso.",
	})
}

func AtualizarStatusOcorrencia(c *fiber.Ctx) error {
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

	var statusOcorrenciaJSON models.StatusOcorrencia

	if err := c.BodyParser(&statusOcorrenciaJSON); err != nil {
		return err
	}

	id := c.Params("id")

	var statusOcorrencia models.StatusOcorrencia

	database.DB.Where("id = ?", id).First(&statusOcorrencia)

	if statusOcorrencia.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "status de ocorrência não encontrado.",
		})
	}

	database.DB.Model(&statusOcorrencia).Select("descricao").Updates(
		models.StatusOcorrencia{
			Descricao: statusOcorrenciaJSON.Descricao,
		})

	return c.JSON(statusOcorrencia)
}
