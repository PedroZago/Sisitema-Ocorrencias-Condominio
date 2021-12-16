package routes

import (
	"github.com/PedroZago/condominitech/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Usuário
	app.Post("/api/usuario/cadastrar", controllers.CadastrarUsuario)
	app.Post("/api/usuario/login", controllers.LogarUsuario)
	app.Post("/api/usuario/sair", controllers.Sair)
	app.Get("/api/usuario/foto-perfil/:id", controllers.DownloadFotoPerfil)
	app.Get("/api/usuario/meus-dados", controllers.BuscarDadosUsuario)
	app.Get("/api/usuario/:id", controllers.BuscarUsuarioID)
	app.Get("/api/usuario", controllers.BuscarTodosUsuarios)
	app.Delete("/api/usuario/:id", controllers.DeletarUsuario)
	app.Put("/api/usuario/senha/:id", controllers.AtualizarSenha)
	app.Put("/api/usuario/:id", controllers.AtualizarUsuario)
	app.Put("/api/usuario/foto-perfil/:id", controllers.UploadFotoPerfil)

	// Unidade
	app.Post("/api/unidade/cadastrar", controllers.CadastrarUnidade)
	app.Get("/api/unidade/:id", controllers.BuscarUnidadeID)
	app.Get("/api/unidade", controllers.BuscarTodasUnidades)
	app.Delete("/api/unidade/:id", controllers.DeletarUnidade)
	app.Put("/api/unidade/:id", controllers.AtualizarUnidade)

	// Bloco
	app.Post("/api/bloco/cadastrar", controllers.CadastrarBloco)
	app.Get("/api/bloco/:id", controllers.BuscarBlocoID)
	app.Get("/api/bloco", controllers.BuscarTodosBlocos)
	app.Delete("/api/bloco/:id", controllers.DeletarBloco)
	app.Put("/api/bloco/:id", controllers.AtualizarBloco)

	// Apartamento
	app.Post("/api/apartamento/cadastrar", controllers.CadastrarApartamento)
	app.Get("/api/apartamento/:id", controllers.BuscarApartamentoID)
	app.Get("/api/apartamento", controllers.BuscarTodosApartamentos)
	app.Delete("/api/apartamento/:id", controllers.DeletarApartamento)
	app.Put("/api/apartamento/:id", controllers.AtualizarApartamento)

	// Tipo de Ocorrência
	app.Post("/api/tipo-ocorrencia/cadastrar", controllers.CadastrarTipoOcorrencia)
	app.Get("/api/tipo-ocorrencia/:id", controllers.BuscarTipoOcorrenciaID)
	app.Get("/api/tipo-ocorrencia", controllers.BuscarTodosTipoOcorrencias)
	app.Delete("/api/tipo-ocorrencia/:id", controllers.DeletarTipoOcorrencia)
	app.Put("/api/tipo-ocorrencia/:id", controllers.AtualizarTipoOcorrencia)

	// Status de Ocorrência
	app.Post("/api/status-ocorrencia/cadastrar", controllers.CadastrarStatusOcorrencia)
	app.Get("/api/status-ocorrencia/:id", controllers.BuscarStatusOcorrenciaID)
	app.Get("/api/status-ocorrencia", controllers.BuscarTodosStatusOcorrencias)
	app.Delete("/api/status-ocorrencia/:id", controllers.DeletarStatusOcorrencia)
	app.Put("/api/status-ocorrencia/:id", controllers.AtualizarStatusOcorrencia)

	// Tipo de Anexo
	app.Post("/api/tipo-anexo/cadastrar", controllers.CadastrarTipoAnexo)
	app.Get("/api/tipo-anexo/:id", controllers.BuscarTipoAnexoID)
	app.Get("/api/tipo-anexo", controllers.BuscarTodosTipoAnexos)
	app.Delete("/api/tipo-anexo/:id", controllers.DeletarTipoAnexo)
	app.Put("/api/tipo-anexo/:id", controllers.AtualizarTipoAnexo)

	// Anexo
	app.Post("/api/anexo/cadastrar/:id", controllers.CadastrarAnexo)
	app.Get("/api/anexo/ocorrencia/:id", controllers.BuscarTodosAnexosPorOcorrencia)
	app.Get("/api/anexo/download/:id", controllers.DownloadAnexo)
	app.Get("/api/anexo/:id", controllers.BuscarAnexoID)
	app.Get("/api/anexo", controllers.BuscarTodosAnexos)
	app.Delete("/api/anexo/:id", controllers.DeletarAnexo)
	app.Put("/api/anexo/:id", controllers.AtualizarAnexo)

	// Ocorrencia
	app.Post("/api/ocorrencia/cadastrar", controllers.CadastrarOcorrencia)
	app.Get("/api/ocorrencia/status", controllers.BuscarContagemStatusDashboard)
	app.Get("/api/ocorrencia/:id", controllers.BuscarOcorrenciaID)
	app.Get("/api/ocorrencia", controllers.BuscarTodasOcorrencias)
	app.Delete("/api/ocorrencia/:id", controllers.DeletarOcorrencia)
	app.Put("/api/ocorrencia/:id", controllers.AtualizarOcorrencia)

	// Responsável
	app.Post("/api/responsavel/cadastrar", controllers.CadastrarResponsavel)
	app.Get("/api/responsavel/unidades", controllers.BuscarTodasUnidadesPorResponsavel)
	app.Get("/api/responsavel/:idUni/:idUsu", controllers.BuscarResponsavelID)
	app.Get("/api/responsavel", controllers.BuscarTodosResponsaveis)
	app.Delete("/api/responsavel/:idUni/:idUsu", controllers.DeletarResponsavel)
	app.Put("/api/responsavel/:idUni/:idUsu", controllers.AtualizarResponsavel)
}
