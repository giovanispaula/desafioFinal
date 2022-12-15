package handler

import (
	"desafioII/internal/dentista"
	"desafioII/internal/domain"
	"desafioII/pkg/web"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

type dentistaHandler struct {
	s dentista.Service
}

func NewDentistaHandler(s dentista.Service) *dentistaHandler {
	return &dentistaHandler{s: s}
}

func (dh *dentistaHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dentista domain.Dentista

		err := ctx.ShouldBindJSON(&dentista)

		if err != nil {
			web.Failure(ctx, 400, errors.New("Parâmetros inválidos"))
			return
		}

		isValid, err := validarDentista(&dentista)

		if isValid {
			web.Failure(ctx, 400, err)
			return
		}

		response, err := dh.s.Post(dentista)

		if err != nil {
			web.Failure(ctx, 400, err)
			return
		}

		web.Success(ctx, 201, response)
	}
}

func (dh *dentistaHandler) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		id, err := strconv.Atoi(idParam)

		if err != nil {
			web.Failure(ctx, 400, errors.New("ID não encontrado"))
			return
		}

		response, err := dh.s.GetById(id)

		if err != nil {
			web.Failure(ctx, 404, errors.New("ID não encontrado"))
			return
		}

		web.Success(ctx, 200, response)
	}
}

func (dh *dentistaHandler) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idParam := ctx.Param("id")

		id, err := strconv.Atoi(idParam)

		if err != nil {
			web.Failure(ctx, 400, errors.New("ID não encontrado"))
			return
		}

		var dentista domain.Dentista

		err = ctx.ShouldBindJSON(&dentista)

		if err != nil {
			web.Failure(ctx, 400, errors.New("Parâmetros inválidos"))
			return
		}

		isValid, err := validarDentista(&dentista)

		if isValid {
			web.Failure(ctx, 400, err)
			return
		}

		response, err := dh.s.Update(id, dentista)

		if err != nil {
			web.Failure(ctx, 404, err)
			return
		}

		web.Success(ctx, 200, response)
	}
}

func (dh *dentistaHandler) Patch() gin.HandlerFunc {
	type request struct {
		Nome      string `json:"nome,omitempty"`
		Sobrenome string `json:"sobrenome,omitempty"`
		Matricula string `json:"matricula,omitempty"`
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

		updateDentista := domain.Dentista{
			Nome:      request.Nome,
			Sobrenome: request.Sobrenome,
			Matricula: request.Matricula,
		}

		response, err := dh.s.Update(id, updateDentista)

		if err != nil {
			web.Failure(ctx, 404, err)
			return
		}

		web.Success(ctx, 200, response)
	}
}

func (dh *dentistaHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idParam := ctx.Param("id")

		id, err := strconv.Atoi(idParam)

		if err != nil {
			web.Failure(ctx, 400, errors.New("ID não encontrado"))
			return
		}

		err = dh.s.Delete(id)

		if err != nil {
			web.Failure(ctx, 404, err)
			return
		}

		ctx.JSON(204, "")
	}
}

func validarDentista(d *domain.Dentista) (bool, error) {

	if d.Nome != "" || d.Sobrenome != "" || d.Matricula != "" {
		return false, errors.New("Todos os campos devem ser preenchidos")
	}

	return true, nil
}
