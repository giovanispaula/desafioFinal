package handler

import (
	"desafioII/internal/domain"
	"desafioII/internal/paciente"
	"desafioII/pkg/web"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

type pacienteHandler struct {
	s paciente.Service
}

func NewPacienteHandler(s paciente.Service) *pacienteHandler {
	return &pacienteHandler{s: s}
}

func (ph *pacienteHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paciente domain.Paciente

		err := ctx.ShouldBindJSON(&paciente)

		if err != nil {
			web.Failure(ctx, 400, errors.New("Parâmetros inválidos"))
			return
		}

		isValid, err := validarPaciente(&paciente)

		if isValid {
			web.Failure(ctx, 400, err)
			return
		}

		response, err := ph.s.Post(paciente)

		if err != nil {
			web.Failure(ctx, 400, err)
			return
		}

		web.Success(ctx, 201, response)
	}
}

func (ph *pacienteHandler) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		id, err := strconv.Atoi(idParam)

		if err != nil {
			web.Failure(ctx, 400, errors.New("ID não encontrado"))
			return
		}

		response, err := ph.s.GetById(id)

		if err != nil {
			web.Failure(ctx, 404, errors.New("ID não encontrado"))
			return
		}

		web.Success(ctx, 200, response)
	}
}

func (ph *pacienteHandler) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idParam := ctx.Param("id")

		id, err := strconv.Atoi(idParam)

		if err != nil {
			web.Failure(ctx, 400, errors.New("ID não encontrado"))
			return
		}

		var paciente domain.Paciente

		err = ctx.ShouldBindJSON(&paciente)

		if err != nil {
			web.Failure(ctx, 400, errors.New("Parâmetros inválidos"))
			return
		}

		isValid, err := validarPaciente(&paciente)

		if isValid {
			web.Failure(ctx, 400, err)
			return
		}

		response, err := ph.s.Update(id, paciente)

		if err != nil {
			web.Failure(ctx, 404, err)
			return
		}

		web.Success(ctx, 200, response)
	}
}

func (ph *pacienteHandler) Patch() gin.HandlerFunc {
	type request struct {
		Nome         string `json:"nome,omitempty"`
		Sobrenome    string `json:"sobrenome,omitempty"`
		Rg           string `json:"registroGeral,omitempty"`
		DataCadastro string `json:"dataCadastro,omitempty"`
	}

	return func(ctx *gin.Context) {

		var request request

		idParam := ctx.Param("id")

		id, err := strconv.Atoi(idParam)

		if err != nil {
			web.Failure(ctx, 400, errors.New("ID não encontrado"))
			return
		}

		err = ctx.ShouldBindJSON(&request)

		if err != nil {
			web.Failure(ctx, 400, errors.New("Parâmetros inválidos"))
			return
		}

		updatePaciente := domain.Paciente{
			Nome:         request.Nome,
			Sobrenome:    request.Sobrenome,
			Rg:           request.Rg,
			DataCadastro: request.DataCadastro,
		}

		response, err := ph.s.Update(id, updatePaciente)

		if err != nil {
			web.Failure(ctx, 404, err)
			return
		}

		web.Success(ctx, 200, response)
	}
}

func (ph *pacienteHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idParam := ctx.Param("id")

		id, err := strconv.Atoi(idParam)

		if err != nil {
			web.Failure(ctx, 400, errors.New("ID não encontrado"))
			return
		}

		err = ph.s.Delete(id)

		if err != nil {
			web.Failure(ctx, 404, err)
			return
		}

		ctx.JSON(204, "")
	}
}

func validarPaciente(p *domain.Paciente) (bool, error) {

	if p.Nome != "" || p.Sobrenome != "" || p.Rg != "" || p.DataCadastro != "" {
		return false, errors.New("Todos os campos são obrigatórios")
	}

	return true, nil
}
