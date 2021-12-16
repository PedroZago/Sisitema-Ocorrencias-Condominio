package controllers

import (
	"strconv"

	"github.com/PedroZago/condominitech/configs"
	"github.com/PedroZago/condominitech/database"
	"github.com/PedroZago/condominitech/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func CadastrarOcorrencia(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "usuário não autenticado.",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var ocorrencia models.Ocorrencia

	if err := c.BodyParser(&ocorrencia); err != nil {
		return err
	}

	ocorrencia.StatusOcorrenciaID = 1
	usuario_id, _ := strconv.ParseUint(claims.Issuer, 10, 0)
	ocorrencia.UsuarioID = uint(usuario_id)

	err = database.DB.Create(&ocorrencia).Error
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "näo foi possível criar a ocorrencia: " + err.Error(),
		})
	}

	return c.JSON(ocorrencia)
}

func BuscarOcorrenciaID(c *fiber.Ctx) error {
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

	var ocorrencia models.ConsultaOcorrencias

	database.DB.Model(&models.Ocorrencia{}).
		Select(`ocorrencia.id AS id, ocorrencia.titulo AS titulo, ocorrencia.descricao AS descricao,
				ocorrencia.tipo_ocorrencia_id AS tipo_ocorrencia_id, tipo_ocorrencia.descricao AS tipo_ocorrencia, ocorrencia.status_ocorrencia_id AS status_ocorrencia_id,
				status_ocorrencia.descricao AS status_ocorrencia, bloco.identificador AS bloco, apartamento.identificador AS apartamento,
				usuario.nome AS responsavel, usuario.id AS responsavel_id, ocorrencia.created_at AS created_at`).
		Joins(`INNER JOIN tipo_ocorrencia ON ocorrencia.tipo_ocorrencia_id = tipo_ocorrencia.id`).
		Joins(`INNER JOIN status_ocorrencia ON ocorrencia.status_ocorrencia_id = status_ocorrencia.id`).
		Joins(`INNER JOIN unidade ON ocorrencia.unidade_id = unidade.id`).
		Joins(`INNER JOIN usuario ON ocorrencia.usuario_id = usuario.id`).
		Joins(`INNER JOIN apartamento ON apartamento.id = unidade.apartamento_id`).
		Joins(`INNER JOIN bloco ON bloco.id = unidade.bloco_id`).
		Where("ocorrencia.id = ?", id).
		Order(`ocorrencia.id`).
		First(&ocorrencia)

	if ocorrencia.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "ocorrência não encontrada.",
		})
	}

	return c.JSON(ocorrencia)
}

func BuscarTodasOcorrencias(c *fiber.Ctx) error {
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

	var ocorrencias []models.ConsultaOcorrencias

	database.DB.Model(&models.Ocorrencia{}).
		Select(`ocorrencia.id AS id, ocorrencia.titulo AS titulo, ocorrencia.descricao AS descricao,
				ocorrencia.tipo_ocorrencia_id AS tipo_ocorrencia_id, tipo_ocorrencia.descricao AS tipo_ocorrencia, ocorrencia.status_ocorrencia_id AS status_ocorrencia_id,
				status_ocorrencia.descricao AS status_ocorrencia, bloco.identificador AS bloco, apartamento.identificador AS apartamento,
				usuario.nome AS responsavel, usuario.id AS responsavel_id, ocorrencia.created_at AS created_at`).
		Joins(`INNER JOIN tipo_ocorrencia ON ocorrencia.tipo_ocorrencia_id = tipo_ocorrencia.id`).
		Joins(`INNER JOIN status_ocorrencia ON ocorrencia.status_ocorrencia_id = status_ocorrencia.id`).
		Joins(`INNER JOIN unidade ON ocorrencia.unidade_id = unidade.id`).
		Joins(`INNER JOIN usuario ON ocorrencia.usuario_id = usuario.id`).
		Joins(`INNER JOIN apartamento ON apartamento.id = unidade.apartamento_id`).
		Joins(`INNER JOIN bloco ON bloco.id = unidade.bloco_id`).
		Order(`id`).
		Scan(&ocorrencias)

	return c.JSON(ocorrencias)
}

func DeletarOcorrencia(c *fiber.Ctx) error {
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

	var ocorrencia models.Ocorrencia

	database.DB.Where("id = ?", id).First(&ocorrencia)

	if ocorrencia.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "ocorrência não encontrada.",
		})
	}

	database.DB.Delete(&models.Ocorrencia{}, ocorrencia.ID)

	return c.JSON(fiber.Map{
		"message": "ocorrência excluída com sucesso.",
	})
}

func AtualizarOcorrencia(c *fiber.Ctx) error {
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

	var ocorrenciaJSON models.Ocorrencia

	if err := c.BodyParser(&ocorrenciaJSON); err != nil {
		return err
	}

	id := c.Params("id")

	var ocorrencia models.Ocorrencia

	database.DB.Where("id = ?", id).First(&ocorrencia)

	if ocorrencia.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "ocorrência não encontrada.",
		})
	}

	database.DB.Model(&ocorrencia).Select("titulo", "descricao", "tipo_ocorrencia_id", "status_ocorrencia_id").Updates(
		models.Ocorrencia{
			Titulo:             ocorrenciaJSON.Titulo,
			Descricao:          ocorrenciaJSON.Descricao,
			TipoOcorrenciaID:   ocorrenciaJSON.TipoOcorrenciaID,
			StatusOcorrenciaID: ocorrenciaJSON.StatusOcorrenciaID,
		})

	return c.JSON(ocorrencia)
}

func BuscarContagemStatusDashboard(c *fiber.Ctx) error {
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

	var cont models.Contagem

	database.DB.Model(&models.Ocorrencia{}).Where("status_ocorrencia_id = ?", 1).Count(&cont.Pendente)
	database.DB.Model(&models.Ocorrencia{}).Where("status_ocorrencia_id = ?", 2).Count(&cont.Atrasada)
	database.DB.Model(&models.Ocorrencia{}).Where("status_ocorrencia_id = ?", 3).Count(&cont.Aprovada)
	database.DB.Model(&models.Ocorrencia{}).Where("status_ocorrencia_id = ?", 4).Count(&cont.Concluida)
	database.DB.Model(&models.Ocorrencia{}).Where("status_ocorrencia_id = ?", 5).Count(&cont.Reprovada)

	return c.JSON(cont)
}
