package controllers

import (
	"github.com/PedroZago/condominitech/configs"
	"github.com/PedroZago/condominitech/database"
	"github.com/PedroZago/condominitech/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func CadastrarUnidade(c *fiber.Ctx) error {
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

	var unidade models.Unidade

	if err := c.BodyParser(&unidade); err != nil {
		return err
	}

	err = database.DB.Create(&unidade).Error
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "näo foi possível criar a unidade: " + err.Error(),
		})
	}

	return c.JSON(unidade)
}

func BuscarUnidadeID(c *fiber.Ctx) error {
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

	var unidade models.ConsultaUnidades

	database.DB.Model(&models.Unidade{}).
		Select(`unidade.id AS id, unidade.bloco_id AS bloco_id, unidade.apartamento_id AS apartamento_id,
				usuario.id AS usuario_id, bloco.identificador AS bloco,
				apartamento.identificador AS apartamento, usuario.nome AS usuario`).
		Joins(`INNER JOIN bloco ON unidade.bloco_id = bloco.id`).
		Joins(`INNER JOIN apartamento ON unidade.apartamento_id = apartamento.id`).
		Joins(`INNER JOIN usuario_unidades ON unidade.id = usuario_unidades.unidade_id`).
		Joins(`INNER JOIN usuario ON usuario_unidades.usuario_id = usuario.id`).
		Where("unidade.id = ?", id).
		Order(`id`).
		Scan(&unidade)

	if unidade.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "unidade não encontrada.",
		})
	}

	return c.JSON(unidade)
}

func BuscarTodasUnidades(c *fiber.Ctx) error {
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

	var unidades []models.ConsultaUnidades

	database.DB.Model(&models.Unidade{}).
		Select(`unidade.id AS id, unidade.bloco_id AS bloco_id, unidade.apartamento_id AS apartamento_id,
				usuario.id AS usuario_id, bloco.identificador AS bloco,
				apartamento.identificador AS apartamento, usuario.nome AS usuario`).
		Joins(`INNER JOIN bloco ON unidade.bloco_id = bloco.id`).
		Joins(`INNER JOIN apartamento ON unidade.apartamento_id = apartamento.id`).
		Joins(`INNER JOIN usuario_unidades ON unidade.id = usuario_unidades.unidade_id`).
		Joins(`INNER JOIN usuario ON usuario_unidades.usuario_id = usuario.id`).
		Order(`id`).
		Scan(&unidades)

	return c.JSON(unidades)
}

func DeletarUnidade(c *fiber.Ctx) error {
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

	var unidade models.Unidade

	database.DB.Where("id = ?", id).First(&unidade)

	if unidade.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "unidade não encontrada.",
		})
	}

	database.DB.Delete(&models.Unidade{}, unidade.ID)

	return c.JSON(fiber.Map{
		"message": "unidade excluída com sucesso.",
	})
}

func AtualizarUnidade(c *fiber.Ctx) error {
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

	var unidadeJSON models.Unidade

	if err := c.BodyParser(&unidadeJSON); err != nil {
		return err
	}

	id := c.Params("id")

	var unidade models.Unidade

	database.DB.Where("id = ?", id).First(&unidade)

	if unidade.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "unidade não encontrada.",
		})
	}

	database.DB.Model(&unidade).Select("bloco_id", "apartamento_id").Updates(
		models.Unidade{
			BlocoID:       unidadeJSON.BlocoID,
			ApartamentoID: unidadeJSON.ApartamentoID,
		})

	return c.JSON(unidade)
}
