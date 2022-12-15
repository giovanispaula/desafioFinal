package main

import (
	"desafioII/cmd/server/handler"
	"desafioII/internal/consulta"
	"desafioII/internal/dentista"
	"desafioII/internal/paciente"
	"desafioII/pkg/store"
	"github.com/gin-gonic/gin"
)

func main() {

	sqlConsulta := store.NewConsultaStore()
	repoConsulta := consulta.NewRepository(sqlConsulta)
	serviceConsulta := consulta.NewService(repoConsulta)
	handlerConsulta := handler.NewConsultaHandler(serviceConsulta)

	sqlDentista := store.NewDentistaStore()
	repoDentista := dentista.NewRepository(sqlDentista)
	serviceDentista := dentista.NewService(repoDentista)
	handlerDentista := handler.NewDentistaHandler(serviceDentista)

	sqlPaciente := store.NewPacienteStore()
	repoPaciente := paciente.NewRepository(sqlPaciente)
	servicePaciente := paciente.NewService(repoPaciente)
	handlerPaciente := handler.NewPacienteHandler(servicePaciente)

	app := gin.New()

	app.Use(gin.Recovery(), gin.Logger())

	consultas := app.Group("/consulta")
	{
		consultas.GET(":id", handlerConsulta.GetById())
		consultas.POST("", handlerConsulta.Post())
		consultas.PUT(":id", handlerConsulta.Put())
		consultas.PATCH(":id", handlerConsulta.Patch())
		consultas.DELETE(":id", handlerConsulta.Delete())
	}

	dentistas := app.Group("/dentista")
	{
		dentistas.GET(":id", handlerDentista.GetById())
		dentistas.POST("", handlerDentista.Post())
		dentistas.PUT(":id", handlerDentista.Put())
		dentistas.PATCH(":id", handlerDentista.Patch())
		dentistas.DELETE(":id", handlerDentista.Delete())
	}

	pacientes := app.Group("/paciente")
	{
		pacientes.GET(":id", handlerPaciente.GetById())
		pacientes.POST("", handlerPaciente.Post())
		pacientes.PUT(":id", handlerPaciente.Put())
		pacientes.PATCH(":id", handlerPaciente.Patch())
		pacientes.DELETE(":id", handlerPaciente.Delete())
	}

	app.Run(":8080")
}
