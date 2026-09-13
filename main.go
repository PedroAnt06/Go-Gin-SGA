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

		// Domínio de Turmas (Classes)
		//v1.POST("/turmas", turmaHandler.CriarTurma)
		//v1.GET("/turmas", turmaHandler.ListarTurmas)
		//v1.POST("/turmas/:id/alocar", turmaHandler.AlocarSala)
	}

	r.Run(":8080")
}
