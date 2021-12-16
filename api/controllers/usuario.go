package controllers

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/PedroZago/condominitech/configs"
	"github.com/PedroZago/condominitech/database"
	"github.com/PedroZago/condominitech/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func CadastrarUsuario(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	e_admin, _ := strconv.ParseBool(data["e_admin"])
	senha, _ := bcrypt.GenerateFromPassword([]byte(data["senha"]), 14)
	usuario := models.Usuario{
		Nome:   data["nome"],
		Email:  data["email"],
		Senha:  senha,
		EAdmin: e_admin,
	}

	err := database.DB.Create(&usuario).Error
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "näo foi possível cadastrar o usuario: " + err.Error(),
		})
	}

	return c.JSON(usuario)
}

func LogarUsuario(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var usuario models.Usuario

	database.DB.Where("email = ?", data["email"]).First(&usuario)

	if usuario.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "usuário não encontrado.",
		})
	}

	if err := bcrypt.CompareHashAndPassword(usuario.Senha, []byte(data["senha"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "senha incorreta.",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(usuario.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(configs.SecretKey))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "não foi possível realizar o login.",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "usuário logado com sucesso.",
	})
}

func BuscarDadosUsuario(c *fiber.Ctx) error {
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

	var usuario models.Usuario

	database.DB.Where("id = ?", claims.Issuer).First(&usuario)

	return c.JSON(usuario)
}

func Sair(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "usuário deslogado com sucesso.",
	})
}

func BuscarUsuarioID(c *fiber.Ctx) error {
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

	var usuario models.Usuario

	database.DB.Where("id = ?", id).First(&usuario)

	if usuario.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "usuário não encontrado.",
		})
	}

	return c.JSON(usuario)
}

func BuscarTodosUsuarios(c *fiber.Ctx) error {
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

	var usuarios []models.Usuario

	database.DB.Find(&usuarios)

	return c.JSON(usuarios)
}

func DeletarUsuario(c *fiber.Ctx) error {
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

	var usuarioToken models.Usuario

	database.DB.Where("id = ?", claims.Issuer).First(&usuarioToken)

	id := c.Params("id")

	var usuarioBanco models.Usuario

	database.DB.Where("id = ?", id).First(&usuarioBanco)

	if usuarioBanco.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "usuário não encontrado.",
		})
	}

	if usuarioBanco.ID != usuarioToken.ID {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "não é possível excluir um usuário que não seja o seu.",
		})
	}

	database.DB.Delete(&models.Usuario{}, usuarioBanco.ID)

	return c.JSON(fiber.Map{
		"message": "usuário excluído com sucesso.",
	})
}

func AtualizarUsuario(c *fiber.Ctx) error {
	// Usuário no Token.
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

	var usuarioToken models.Usuario

	database.DB.Where("id = ?", claims.Issuer).First(&usuarioToken)

	// Usuário na Requisição.
	id := c.Params("id")

	var usuarioBanco models.Usuario

	database.DB.Where("id = ?", id).First(&usuarioBanco)

	if usuarioBanco.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "usuário não encontrado.",
		})
	}

	if usuarioBanco.ID != usuarioToken.ID {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "não é possível atualizar um usuário que não seja o seu.",
		})
	}

	// Atualização do usuário.
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	database.DB.Model(&usuarioBanco).Select("nome", "email").Updates(
		models.Usuario{
			Nome:  data["nome"],
			Email: data["email"],
		})

	return c.JSON(fiber.Map{
		"message": "usuário atualizado com sucesso.",
		"usuario": usuarioBanco,
	})
}

func AtualizarSenha(c *fiber.Ctx) error {
	// Usuário no Token.
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

	var usuarioToken models.Usuario

	database.DB.Where("id = ?", claims.Issuer).First(&usuarioToken)

	// Usuário na Requisição.
	id := c.Params("id")

	var usuarioBanco models.Usuario

	database.DB.Where("id = ?", id).First(&usuarioBanco)

	if usuarioBanco.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "usuário não encontrado.",
		})
	}

	if usuarioBanco.ID != usuarioToken.ID {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "não é possível atualizar a senha de um usuário que não seja o seu.",
		})
	}

	// Senha no banco
	var senhaBanco models.Usuario

	database.DB.Model(&models.Usuario{}).Select("senha").Where("id = ?", id).First(&senhaBanco)

	// Atualização da senha.
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword(senhaBanco.Senha, []byte(data["senha_atual"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "a senha não condiz com a que está salva no banco.",
		})
	}

	senha, _ := bcrypt.GenerateFromPassword([]byte(data["senha_nova"]), 14)

	database.DB.Model(&usuarioBanco).Select("senha").Updates(
		models.Usuario{
			Senha: senha,
		})

	return c.JSON(fiber.Map{
		"message": "senha atualizada com sucesso.",
	})
}

func UploadFotoPerfil(c *fiber.Ctx) error {
	if form, err := c.MultipartForm(); err == nil {
		// Usuário no Token.
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

		var usuarioToken models.Usuario

		database.DB.Where("id = ?", claims.Issuer).First(&usuarioToken)

		// Usuário na Requisição.
		id := c.Params("id")

		var usuarioBanco models.Usuario

		database.DB.Where("id = ?", id).First(&usuarioBanco)

		if usuarioBanco.ID == 0 {
			c.Status(fiber.StatusNotFound)
			return c.JSON(fiber.Map{
				"message": "usuário não encontrado.",
			})
		}

		if usuarioBanco.ID != usuarioToken.ID {
			c.Status(fiber.StatusNotFound)
			return c.JSON(fiber.Map{
				"message": "não é possível atualizar um usuário que não seja o seu.",
			})
		}

		files := form.File["foto"]

		for _, file := range files {
			//fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
			os.Mkdir(fmt.Sprintf("./archives/usuarios/%s", usuarioBanco.Email), 0700)

			if err := c.SaveFile(file, fmt.Sprintf("./archives/usuarios/%s/%s", usuarioBanco.Email, file.Filename)); err != nil {
				fmt.Println(err)
				return err
			}

			database.DB.Model(&usuarioBanco).Select("url_foto_perfil").Updates(
				models.Usuario{
					URLFotoPerfil: fmt.Sprintf("./archives/usuarios/%s/%s", usuarioBanco.Email, file.Filename),
				})
		}
		return err
	}
	return c.JSON(fiber.Map{
		"message": "foto de perfil atualizada com sucesso.",
	})
}

func DownloadFotoPerfil(c *fiber.Ctx) error {
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

	var usuario models.Usuario

	database.DB.Where("id = ?", id).First(&usuario)

	return c.SendFile(usuario.URLFotoPerfil)
}
