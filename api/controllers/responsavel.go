package controllers

import (
	"github.com/PedroZago/condominitech/configs"
	"github.com/PedroZago/condominitech/database"
	"github.com/PedroZago/condominitech/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func CadastrarResponsavel(c *fiber.Ctx) error {
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

	var responsavel models.UsuarioUnidade

	if err := c.BodyParser(&responsavel); err != nil {
		return err
	}

	err = database.DB.Create(&responsavel).Error
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "näo foi possível criar o responsavel: " + err.Error(),
		})
	}

	return c.JSON(responsavel)
}

func BuscarResponsavelID(c *fiber.Ctx) error {
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

	idUni := c.Params("idUni")
	idUsu := c.Params("idUsu")

	var responsavel models.UsuarioUnidade

	database.DB.Where("unidade_id = ? AND usuario_id = ?", idUni, idUsu).First(&responsavel)

	if responsavel.UsuarioID == 0 || responsavel.UnidadeID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "responsavel não encontrado.",
		})
	}

	return c.JSON(responsavel)
}

func BuscarTodasUnidadesPorResponsavel(c *fiber.Ctx) error {
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

	var responsavel []models.UsuarioUnidade

	database.DB.Where("usuario_id = ?", claims.Issuer).First(&responsavel)

	return c.JSON(responsavel)
}

func BuscarTodosResponsaveis(c *fiber.Ctx) error {
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

	var responsaveis []models.UsuarioUnidade

	database.DB.Find(&responsaveis)

	return c.JSON(responsaveis)
}

func DeletarResponsavel(c *fiber.Ctx) error {
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

	idUni := c.Params("idUni")
	idUsu := c.Params("idUsu")

	var responsavel models.UsuarioUnidade

	database.DB.Where("unidade_id = ? AND usuario_id = ?", idUni, idUsu).First(&responsavel)

	if responsavel.UsuarioID == 0 || responsavel.UnidadeID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "responsavel não encontrado.",
		})
	}

	database.DB.Where("unidade_id = ? AND usuario_id = ?", idUni, idUsu).Delete(&models.UsuarioUnidade{})

	return c.JSON(fiber.Map{
		"message": "responsavel excluído com sucesso.",
	})
}

func AtualizarResponsavel(c *fiber.Ctx) error {
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

	var responsavelJSON models.UsuarioUnidade

	if err := c.BodyParser(&responsavelJSON); err != nil {
		return err
	}

	idUni := c.Params("idUni")
	idUsu := c.Params("idUsu")

	var responsavel models.UsuarioUnidade

	database.DB.Where("unidade_id = ? AND usuario_id = ?", idUni, idUsu).First(&responsavel)

	if responsavel.UsuarioID == 0 || responsavel.UnidadeID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "responsavel não encontrado.",
		})
	}

	database.DB.Model(&responsavel).Select("unidade_id", "usuario_id").Updates(
		models.UsuarioUnidade{
			UnidadeID: responsavelJSON.UnidadeID,
			UsuarioID: responsavelJSON.UsuarioID,
		})

	return c.JSON(responsavel)
}
