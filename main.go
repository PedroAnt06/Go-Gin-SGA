package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	alunosRepositorio := NovoAlunosRepositorio()
	alunosService := NovoAlunoService(alunosRepositorio)

	salasRepositorio := NovoSalasRepositorio()
	salasService := NovoSalaService(salasRepositorio)

	turmasRepositorio := NovoTurmasRepositorio()
	turmasService := NovoTurmaService(turmasRepositorio, salasRepositorio)

	r := gin.New()

	// Uso dos Middlewares globais nativos e personalizados
	r.Use(gin.Recovery())

	// 4. Mapeamento de Rotas sob Grupo Versionado
	v1 := r.Group("/api/v1")
	{
		// Monitoramento da API
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":    "healthy",
				"timestamp": time.Now(),
				"version":   "1.0.0",
			})
		})

		// Domínio de Alunos
		v1.POST("/alunos", func(c *gin.Context) {
			var body Aluno
			if err := c.ShouldBindJSON(&body); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
				return
			}

			aluno, err := alunosService.CriarAluno(body.Matrícula, body.Nome, body.Email)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
				return
			}

			c.JSON(http.StatusCreated, aluno)
		})

		v1.GET("/alunos", func(c *gin.Context) {
			c.JSON(http.StatusOK, alunosService.ListarAlunos())
		})

		v1.GET("/alunos/:matricula", func(c *gin.Context) {
			matriculaParam := c.Param("matricula")
			matricula, err := strconv.Atoi(matriculaParam)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"erro": "Matrícula inválida"})
				return
			}
			aluno, err := alunosService.BuscarAlunoPorMatricula(matricula)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
				return
			}
			c.JSON(http.StatusOK, aluno)
		})

		// Domínio de Salas
		v1.POST("/salas", func(c *gin.Context) {
			var body Sala
			if err := c.ShouldBindJSON(&body); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
				return
			}

			sala, err := salasService.CriarSala(body.ID, body.Nome, body.Capacidade, body.Recursos)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
				return
			}

			c.JSON(http.StatusCreated, sala)
		})

		v1.GET("/salas", func(c *gin.Context) {
			c.JSON(http.StatusOK, salasService.ListarSalas())
		})

		v1.GET("/salas/:id", func(c *gin.Context) {
			idParam := c.Param("id")
			id, err := strconv.Atoi(idParam)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
				return
			}
			sala, err := salasService.BuscarSalaPorID(id)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
				return
			}
			c.JSON(http.StatusOK, sala)
		})

		// Domínio de Turmas (Classes)
		v1.POST("/turmas", func(c *gin.Context) {
			var body Turma
			if err := c.ShouldBindJSON(&body); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
				return
			}

			turma, err := turmasService.CriarTurma(body.ID, body.Nome, body.Disciplina, body.Professor)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
				return
			}

			c.JSON(http.StatusCreated, turma)
		})

		v1.GET("/turmas", func(c *gin.Context) {
			c.JSON(http.StatusOK, turmasService.ListarTurmas())
		})

		v1.GET("/turmas/:id", func(c *gin.Context) {
			idParam := c.Param("id")
			id, err := strconv.Atoi(idParam)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
				return
			}
			turma, err := turmasService.BuscarTurmaPorID(id)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"erro": err.Error()})
				return
			}
			c.JSON(http.StatusOK, turma)
		})

		v1.POST("/turmas/:id/alocar", func(c *gin.Context) {
			idParam := c.Param("id")
			id, err := strconv.Atoi(idParam)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
				return
			}

			var body Alocacao
			if err := c.ShouldBindJSON(&body); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
				return
			}

			turma, err := turmasService.AlocarSala(id, body.SalaID, body.DiaSemana, body.HorarioInicio, body.HorarioFim)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
				return
			}

			c.JSON(http.StatusOK, turma)
		})
	}

	r.Run(":8080")
}
