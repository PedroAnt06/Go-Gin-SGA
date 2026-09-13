package main

import "errors"

type AlunoService struct {
	repositorio *AlunosRepositorio
}

func (a *AlunoService) ListarAlunos() []Aluno {
	return a.repositorio.listarAlunos()
}

func (a *AlunoService) CriarAluno(nome string, email string) (Aluno, error) {
	if nome == "" || email == "" {
		return Aluno{}, errors.New("preencha pfvr o nome e/ou email")
	}

	return a.repositorio.criarAluno(nome, email), nil
}
