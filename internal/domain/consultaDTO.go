package domain

type ConsultaDTO struct {
	Consulta
	Dentista Dentista `json:"dentista" binding:"required"`
	Paciente Paciente `json:"paciente" binding:"required"`
}
