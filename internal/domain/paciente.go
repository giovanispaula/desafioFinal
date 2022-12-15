package domain

type Paciente struct {
	Id           int    `json:"id"`
	Nome         string `json:"nome" binding:"required"`
	Sobrenome    string `json:"sobrenome" binding:"required"`
	Rg           string `json:"rg" binding:"required"`
	DataCadastro string `json:"dataCadastro" binding:"required"`
}
