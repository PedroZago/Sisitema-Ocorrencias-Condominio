package database

import (
	"github.com/PedroZago/condominitech/configs"
	"github.com/PedroZago/condominitech/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open(configs.ConexaoBanco), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.Migrator().DropTable(&models.Bloco{})
	connection.AutoMigrate(&models.Bloco{})

	connection.Migrator().DropTable(&models.Apartamento{})
	connection.AutoMigrate(&models.Apartamento{})

	connection.Migrator().DropTable(&models.TipoOcorrencia{})
	connection.AutoMigrate(&models.TipoOcorrencia{})

	connection.Migrator().DropTable(&models.TipoAnexo{})
	connection.AutoMigrate(&models.TipoAnexo{})

	connection.Migrator().DropTable(&models.StatusOcorrencia{})
	connection.AutoMigrate(&models.StatusOcorrencia{})

	connection.Migrator().DropTable("usuario_unidades")
	connection.SetupJoinTable(&models.Usuario{}, "Unidades", &models.UsuarioUnidade{})

	connection.Migrator().DropTable(&models.Usuario{})
	connection.AutoMigrate(&models.Usuario{})

	connection.Migrator().DropTable(&models.Unidade{})
	connection.AutoMigrate(&models.Unidade{})

	connection.Migrator().DropTable(&models.Ocorrencia{})
	connection.AutoMigrate(&models.Ocorrencia{})

	connection.Migrator().DropTable(&models.Anexo{})
	connection.AutoMigrate(&models.Anexo{})

	tipoOcorrencia := []models.TipoOcorrencia{
		{ID: 1, Descricao: "Notificação financeira"},
		{ID: 2, Descricao: "Outras infrações da convenção sem multa"},
		{ID: 3, Descricao: "Outras infrações da convenção com multa"},
		{ID: 4, Descricao: "Convivência: Infração da convenção sem multa"},
		{ID: 5, Descricao: "Convivência: Infração da convenção com multa"},
		{ID: 6, Descricao: "Obra: Infração da convenção sem multa"},
		{ID: 7, Descricao: "Obra: Infração da convenção com multa"},
	}
	connection.Create(&tipoOcorrencia)

	tipoAnexo := []models.TipoAnexo{
		{ID: 1, Tipo: "application/pdf", Descricao: "PDF"},
		{ID: 2, Tipo: "image/png", Descricao: "PNG"},
		{ID: 3, Tipo: "image/jpeg", Descricao: "JPEG"},
		{ID: 4, Tipo: "video/mp4", Descricao: "MP4"},
		{ID: 5, Tipo: "audio/mp3", Descricao: "MP3"},
		{ID: 6, Tipo: "application/vnd.openxmlformats-officedocument.wordprocessingml.document", Descricao: "DOCX"},
		{ID: 7, Tipo: "application/vnd.openxmlformats-officedocument.presentationml.presentation", Descricao: "PPTX"},
		{ID: 8, Tipo: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", Descricao: "XLSX"},
		{ID: 9, Tipo: "text/plain", Descricao: "TXT"},
	}
	connection.Create(&tipoAnexo)

	statusOcorrencia := []models.StatusOcorrencia{
		{ID: 1, Descricao: "Pendente", Cor: "ffff00"},
		{ID: 2, Descricao: "Atrasada", Cor: "ff0000"},
		{ID: 3, Descricao: "Aprovada", Cor: "0000ff"},
		{ID: 4, Descricao: "Concluída", Cor: "00ff00"},
		{ID: 5, Descricao: "Reprovada", Cor: "00ff00"},
	}
	connection.Create(&statusOcorrencia)

	senha, _ := bcrypt.GenerateFromPassword([]byte("73593186"), 14)
	usuario := []models.Usuario{
		{
			ID:     1,
			Nome:   "Administrador",
			Email:  "admin@mail.com",
			Senha:  senha,
			EAdmin: true,
		},
	}
	connection.Create(&usuario)

	// apartamento := []models.Apartamento{
	// 	{
	// 		ID:            1,
	// 		Identificador: "APT001",
	// 		Descricao:     "Apartamento número 1",
	// 	},
	// 	{
	// 		ID:            2,
	// 		Identificador: "APT002",
	// 		Descricao:     "Apartamento número 2",
	// 	},
	// }
	// connection.Create(&apartamento)

	// bloco := []models.Bloco{
	// 	{
	// 		ID:            1,
	// 		Identificador: "BLC001",
	// 		Descricao:     "Bloco número 1",
	// 	},
	// 	{
	// 		ID:            2,
	// 		Identificador: "BLC002",
	// 		Descricao:     "Bloco número 2",
	// 	},
	// }
	// connection.Create(&bloco)

	// unidade := []models.Unidade{
	// 	{
	// 		ID:            1,
	// 		ApartamentoID: 1,
	// 		BlocoID:       1,
	// 	},
	// 	{
	// 		ID:            2,
	// 		ApartamentoID: 2,
	// 		BlocoID:       2,
	// 	},
	// }
	// connection.Create(&unidade)

	// usuario_unidades := []models.UsuarioUnidade{
	// 	{
	// 		UsuarioID: 2,
	// 		UnidadeID: 1,
	// 	},
	// 	{
	// 		UsuarioID: 3,
	// 		UnidadeID: 2,
	// 	},
	// }
	// connection.Create(&usuario_unidades)

	// ocorrencia := []models.Ocorrencia{
	// 	{
	// 		ID:                 1,
	// 		Titulo:             "Teste 1",
	// 		Descricao:          "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
	// 		TipoOcorrenciaID:   1,
	// 		StatusOcorrenciaID: 1,
	// 		UnidadeID:          2,
	// 		UsuarioID:          3,
	// 	},
	// 	{
	// 		ID:                 2,
	// 		Titulo:             "Teste 2",
	// 		Descricao:          "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
	// 		TipoOcorrenciaID:   2,
	// 		StatusOcorrenciaID: 1,
	// 		UnidadeID:          2,
	// 		UsuarioID:          3,
	// 	},
	// 	{
	// 		ID:                 3,
	// 		Titulo:             "Teste 3",
	// 		Descricao:          "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
	// 		TipoOcorrenciaID:   3,
	// 		StatusOcorrenciaID: 2,
	// 		UnidadeID:          1,
	// 		UsuarioID:          2,
	// 	},
	// 	{
	// 		ID:                 4,
	// 		Titulo:             "Teste 4",
	// 		Descricao:          "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
	// 		TipoOcorrenciaID:   4,
	// 		StatusOcorrenciaID: 2,
	// 		UnidadeID:          1,
	// 		UsuarioID:          2,
	// 	},
	// 	{
	// 		ID:                 5,
	// 		Titulo:             "Teste 5",
	// 		Descricao:          "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
	// 		TipoOcorrenciaID:   5,
	// 		StatusOcorrenciaID: 3,
	// 		UnidadeID:          2,
	// 		UsuarioID:          3,
	// 	},
	// 	{
	// 		ID:                 6,
	// 		Titulo:             "Teste 6",
	// 		Descricao:          "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
	// 		TipoOcorrenciaID:   6,
	// 		StatusOcorrenciaID: 1,
	// 		UnidadeID:          2,
	// 		UsuarioID:          3,
	// 	},
	// 	{
	// 		ID:                 7,
	// 		Titulo:             "Teste 7",
	// 		Descricao:          "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
	// 		TipoOcorrenciaID:   7,
	// 		StatusOcorrenciaID: 4,
	// 		UnidadeID:          1,
	// 		UsuarioID:          2,
	// 	},
	// 	{
	// 		ID:                 8,
	// 		Titulo:             "Teste 8",
	// 		Descricao:          "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
	// 		TipoOcorrenciaID:   1,
	// 		StatusOcorrenciaID: 4,
	// 		UnidadeID:          1,
	// 		UsuarioID:          2,
	// 	},
	// 	{
	// 		ID:                 9,
	// 		Titulo:             "Teste 9",
	// 		Descricao:          "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
	// 		TipoOcorrenciaID:   2,
	// 		StatusOcorrenciaID: 5,
	// 		UnidadeID:          2,
	// 		UsuarioID:          3,
	// 	},
	// 	{
	// 		ID:                 10,
	// 		Titulo:             "Teste 10",
	// 		Descricao:          "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
	// 		TipoOcorrenciaID:   3,
	// 		StatusOcorrenciaID: 5,
	// 		UnidadeID:          2,
	// 		UsuarioID:          3,
	// 	},
	// }
	// connection.Create(&ocorrencia)

}
