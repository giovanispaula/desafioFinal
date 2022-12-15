package domain

type Consulta struct {
	Id           int    `json:"id"`
	Descricao    string `json:"descricao" binding:"required"`
	DataConsulta string `json:"dataConsulta" binding:"required"`
	DentistaId   int    `json:"dentistaId" binding:"required"`
	PacienteId   int    `json:"pacienteId" binding:"required"`
}
