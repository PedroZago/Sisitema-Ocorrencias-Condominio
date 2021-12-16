package controllers

import (
	"fmt"
	"os"
	"strconv"

	"github.com/PedroZago/condominitech/configs"
	"github.com/PedroZago/condominitech/database"
	"github.com/PedroZago/condominitech/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func CadastrarAnexo(c *fiber.Ctx) error {
	if form, err := c.MultipartForm(); err == nil {
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
		ocorrencia_id, _ := strconv.ParseUint(id, 10, 0)

		var ocorrencia models.Ocorrencia

		database.DB.Where("id = ?", id).First(&ocorrencia)

		if ocorrencia.ID == 0 {
			c.Status(fiber.StatusNotFound)
			return c.JSON(fiber.Map{
				"message": "ocorrência não encontrado.",
			})
		}

		var tipoAnexos []models.TipoAnexo

		database.DB.Find(&tipoAnexos)

		var tipos []string

		for _, tipoAnexo := range tipoAnexos {
			tipos = append(tipos, tipoAnexo.Tipo)
		}

		files := form.File["anexo"]

		for _, file := range files {
			//fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
			os.Mkdir(fmt.Sprintf("./archives/anexos/ocorrencia-%d", ocorrencia.ID), 0700)

			if err := c.SaveFile(file, fmt.Sprintf("./archives/anexos/ocorrencia-%d/%s", ocorrencia.ID, file.Filename)); err != nil {
				return err
			}

			index := IndexOf(file.Header["Content-Type"][0], tipos)

			var anexo models.Anexo

			anexo.Nome = file.Filename
			anexo.Caminho = fmt.Sprintf("./archives/anexos/ocorrencia-%d/%s", ocorrencia.ID, file.Filename)
			anexo.TipoAnexoID = uint(index + 1)
			anexo.OcorrenciaID = uint(ocorrencia_id)

			err = database.DB.Create(&anexo).Error
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(fiber.Map{
					"message": "näo foi possível adicionar o anexo: " + err.Error(),
				})
			}
		}
		return err
	}
	return nil
}

func BuscarAnexoID(c *fiber.Ctx) error {
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

	var anexo models.Anexo

	database.DB.Where("id = ?", id).First(&anexo)

	if anexo.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "anexo não encontrado.",
		})
	}

	return c.JSON(anexo)
}

func BuscarTodosAnexos(c *fiber.Ctx) error {
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

	var anexos []models.Anexo

	database.DB.Find(&anexos)

	return c.JSON(anexos)
}

func BuscarTodosAnexosPorOcorrencia(c *fiber.Ctx) error {
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

	var anexos []models.Anexo

	database.DB.Where("ocorrencia_id = ?", id).Find(&anexos)

	return c.JSON(anexos)
}

func DeletarAnexo(c *fiber.Ctx) error {
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

	var anexo models.Anexo

	database.DB.Where("id = ?", id).First(&anexo)

	if anexo.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "anexo não encontrado.",
		})
	}

	database.DB.Delete(&models.Anexo{}, anexo.ID)

	return c.JSON(fiber.Map{
		"message": "anexo excluído com sucesso.",
	})
}

func AtualizarAnexo(c *fiber.Ctx) error {
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

	var anexoJSON models.Anexo

	if err := c.BodyParser(&anexoJSON); err != nil {
		return err
	}

	id := c.Params("id")

	var anexo models.Anexo

	database.DB.Where("id = ?", id).First(&anexo)

	if anexo.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "anexo não encontrado.",
		})
	}

	database.DB.Model(&anexo).Select("nome", "descricao", "tipo_anexo_id").Updates(
		models.Anexo{
			Nome:        anexoJSON.Nome,
			Caminho:     anexoJSON.Caminho,
			TipoAnexoID: anexoJSON.TipoAnexoID,
		})

	return c.JSON(anexo)
}

func DownloadAnexo(c *fiber.Ctx) error {
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

	var anexo models.Anexo

	database.DB.Where("id = ?", id).First(&anexo)

	return c.Download(anexo.Caminho)
}

func IndexOf(elemento string, dados []string) int {
	for k, v := range dados {
		if elemento == v {
			return k
		}
	}
	return -1 //not found.
}
