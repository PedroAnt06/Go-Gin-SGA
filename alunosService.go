package main

import "errors"

type AlunoService struct {
	repositorio *AlunosRepositorio
}

func NovoAlunoService(repositorio *AlunosRepositorio) *AlunoService {
	return &AlunoService{
		repositorio: repositorio,
	}
}

func (a *AlunoService) ListarAlunos() []Aluno {
	return a.repositorio.listarAlunos()
}

func (a *AlunoService) CriarAluno(matricula int, nome string, email string) (Aluno, error) {
	if matricula == 0 || nome == "" || email == "" {
		return Aluno{}, errors.New("preencha pfv todos os campos")
	}

	if a.repositorio.buscarAlunoPorMatricula(matricula) != nil {
		return Aluno{}, errors.New("Aluno já existe")
	}

	return a.repositorio.criarAluno(matricula, nome, email), nil
}
