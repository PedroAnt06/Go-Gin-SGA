package main

type AlunoService struct {
	repositorio *AlunosRepositorio
}

func (a *AlunoService) ListarAlunos() []Aluno {
	return a.repositorio.listarAlunos()
}

func (a *AlunoService) CriarAluno(nome string, email string) Aluno {
	if nome == "" || email == "" {
		return Aluno{}
	}

	return a.repositorio.criarAluno(nome, email)
}
