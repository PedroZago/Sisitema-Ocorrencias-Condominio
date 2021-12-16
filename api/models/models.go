// Relacionamentos:
//   - Uma unidade possui vários blocos e um bloco possui uma unidade.
//   - Um bloco possui vários apartamentos e um  apartamento possui um bloco.
//   - Um tipo de ocorrência possui várias ocorrências e uma ocorrência possui um tipo de ocorrência.
//   - Um status de ocorrência possui várias ocorrências e uma ocorrência possui um status de ocorrência.
//   - Um anexo de ocorrência possui uma ocorrência e uma ocorrência possui um anexo.
//   - Uma unidade possui várias ocorrências e uma ocorrência possui uma unidade.
//   - Uma unidade possui um responsável e um responsável possui uma unidade.
package models

import (
	"time"

	"gorm.io/gorm"
)

type Tabler interface {
	TableName() string
}

func (Usuario) TableName() string {
	return "usuario"
}

func (Unidade) TableName() string {
	return "unidade"
}

func (Bloco) TableName() string {
	return "bloco"
}

func (Apartamento) TableName() string {
	return "apartamento"
}

func (TipoOcorrencia) TableName() string {
	return "tipo_ocorrencia"
}

func (StatusOcorrencia) TableName() string {
	return "status_ocorrencia"
}

func (Anexo) TableName() string {
	return "anexo"
}

func (TipoAnexo) TableName() string {
	return "tipo_anexo"
}

func (Ocorrencia) TableName() string {
	return "ocorrencia"
}

type Usuario struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	Nome          string         `json:"nome" gorm:"not null"`
	Email         string         `json:"email" gorm:"not null;unique"`
	Senha         []byte         `json:"-" gorm:"not null;type:longtext"`
	URLFotoPerfil string         `json:"url_foto_perfil"`
	EAdmin        bool           `json:"e_admin" gorm:"type:bool"`
	Ocorrencia    []Ocorrencia   `json:"-"`
	Unidades      []Unidade      `json:"-" gorm:"many2many:usuario_unidades"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}

type UsuarioUnidade struct {
	UsuarioID uint           `json:"usuario_id" gorm:"primaryKey"`
	UnidadeID uint           `json:"unidade_id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Unidade struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	BlocoID       uint           `json:"bloco_id" gorm:"not null"`
	ApartamentoID uint           `json:"apartamento_id" gorm:"not null"`
	Ocorrencia    []Ocorrencia   `json:"-"`
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}

type Bloco struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	Identificador string         `json:"identificador" gorm:"not null"`
	Descricao     string         `json:"descricao" gorm:"not null"`
	Unidade       []Unidade      `json:"-"`
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}

type Apartamento struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	Identificador string         `json:"identificador" gorm:"not null"`
	Descricao     string         `json:"descricao" gorm:"not null"`
	Unidade       []Unidade      `json:"-"`
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}

type TipoOcorrencia struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	Descricao  string         `json:"descricao" gorm:"not null"`
	Ocorrencia []Ocorrencia   `json:"-"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}

type StatusOcorrencia struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	Descricao  string         `json:"descricao" gorm:"not null"`
	Cor        string         `json:"cor" gorm:"not null"`
	Ocorrencia []Ocorrencia   `json:"-"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}

type TipoAnexo struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Descricao string         `json:"descricao" gorm:"not null"`
	Tipo      string         `json:"tipo" gorm:"not null"`
	Anexo     []Anexo        `json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Anexo struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Nome         string         `json:"nome" gorm:"not null"`
	Caminho      string         `json:"caminho" gorm:"not null"`
	TipoAnexoID  uint           `json:"tipo_anexo_id" gorm:"not null"`
	OcorrenciaID uint           `json:"ocorrencia_id" gorm:"not null"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

type Ocorrencia struct {
	ID                 uint           `json:"id" gorm:"primaryKey"`
	Titulo             string         `json:"titulo" gorm:"not null"`
	Descricao          string         `json:"descricao" gorm:"not null"`
	TipoOcorrenciaID   uint           `json:"tipo_ocorrencia_id" gorm:"not null"`
	StatusOcorrenciaID uint           `json:"status_ocorrencia_id" gorm:"not null"`
	UnidadeID          uint           `json:"unidade_id" gorm:"not null"`
	UsuarioID          uint           `json:"usuario_id" gorm:"not null"`
	Anexo              []Anexo        `json:"-"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `json:"delete_at" gorm:"index"`
}

type ConsultaOcorrencias struct {
	ID                 uint      `json:"id"`
	Titulo             string    `json:"titulo"`
	Descricao          string    `json:"descricao"`
	TipoOcorrenciaID   uint      `json:"tipo_ocorrencia_id"`
	TipoOcorrencia     string    `json:"tipo_ocorrencia"`
	StatusOcorrenciaID uint      `json:"status_ocorrencia_id"`
	StatusOcorrencia   string    `json:"status_ocorrencia"`
	Bloco              string    `json:"bloco"`
	Apartamento        string    `json:"apartamento"`
	Responsavel        string    `json:"responsavel"`
	ResponsavelID      uint      `json:"responsavel_id"`
	CreatedAt          time.Time `json:"created_at"`
}

type ConsultaUnidades struct {
	ID            uint   `json:"id"`
	BlocoID       uint   `json:"bloco_id"`
	Bloco         string `json:"bloco"`
	ApartamentoID uint   `json:"apartamento_id"`
	Apartamento   string `json:"apartamento"`
	UsuarioID     uint   `json:"usuario_id"`
	Usuario       string `json:"usuario"`
}

type Contagem struct {
	Pendente  int64 `json:"pendente"`
	Atrasada  int64 `json:"atrasada"`
	Aprovada  int64 `json:"aprovada"`
	Concluida int64 `json:"concluida"`
	Reprovada int64 `json:"reprovada"`
}
