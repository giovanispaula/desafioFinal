package handler

import (
	"desafioII/internal/consulta"
	"desafioII/internal/domain"
	"desafioII/pkg/web"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

type consultaHandler struct {
	s consulta.Service
}

func NewConsultaHandler(s consulta.Service) *consultaHandler {
	return &consultaHandler{s: s}
}

func (ch *consultaHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var consulta domain.Consulta

		err := ctx.ShouldBindJSON(&consulta)

		if err != nil {
			web.Failure(ctx, 400, errors.New("Parâmetros inválidos"))
			return
		}

		isValid, err := validarConsulta(&consulta)

		if isValid {
			web.Failure(ctx, 400, err)
			return
		}

		response, err := ch.s.Post(consulta)

		if err != nil {
			web.Failure(ctx, 400, err)
			return
		}

		web.Success(ctx, 201, response)
	}
}

func (ch *consultaHandler) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		id, err := strconv.Atoi(idParam)

		if err != nil {
			web.Failure(ctx, 400, errors.New("ID não encontrado"))
			return
		}

		response, err := ch.s.GetById(id)

		if err != nil {
			web.Failure(ctx, 404, errors.New("ID não encontrado"))
			return
		}

		web.Success(ctx, 200, response)
	}
}

func (ch *consultaHandler) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idParam := ctx.Param("id")

		id, err := strconv.Atoi(idParam)

		if err != nil {
			web.Failure(ctx, 400, errors.New("ID não encontrado"))
			return
		}

		var consulta domain.Consulta

		err = ctx.ShouldBindJSON(&consulta)

		if err != nil {
			web.Failure(ctx, 400, err)
			return
		}

		isValid, err := validarConsulta(&consulta)

		if isValid {
			web.Failure(ctx, 400, err)
			return
		}

		response, err := ch.s.Update(id, consulta)

		if err != nil {
			web.Failure(ctx, 404, err)
			return
		}

		web.Success(ctx, 200, response)
	}
}

func (ch *consultaHandler) Patch() gin.HandlerFunc {
	type request struct {
		Descricao    string `json:"descricao"`
		DataConsulta string `json:"dataConsulta"`
		DentistaId   int    `json:"dentistaId"`
		PacienteId   int    `json:"pacienteId"`
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
			web.Failure(ctx, 400, err)
			return
		}

		updateConsulta := domain.Consulta{
			Descricao:    request.Descricao,
			DataConsulta: request.DataConsulta,
			DentistaId:   request.DentistaId,
			PacienteId:   request.PacienteId,
		}

		response, err := ch.s.Update(id, updateConsulta)

		if err != nil {
			web.Failure(ctx, 404, err)
			return
		}

		web.Success(ctx, 200, response)
	}
}

func (ch *consultaHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idParam := ctx.Param("id")

		id, err := strconv.Atoi(idParam)

		if err != nil {
			web.Failure(ctx, 400, errors.New("Id não encontrado"))
			return
		}

		err = ch.s.Delete(id)

		if err != nil {
			web.Failure(ctx, 404, err)
			return
		}

		ctx.JSON(204, "")
	}
}

func validarConsulta(c *domain.Consulta) (bool, error) {

	if c.Descricao != "" || c.DataConsulta != "" || c.DentistaId < 0 || c.PacienteId < 0 {
		return false, errors.New("Todos os campos são obrigatórios")
	}

	return true, nil
}
