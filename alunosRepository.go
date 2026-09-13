package main

type Aluno struct {
	Matrícula int    `json:"matricula"`
	Nome      string `json:"nome"`
	Email     string `json:"email"`
}

type AlunosRepositorio struct {
	alunos []Aluno
}

func NovoAlunosRepositorio() *AlunosRepositorio {
	return &AlunosRepositorio{
		alunos: []Aluno{},
	}
}

func (r *AlunosRepositorio) listarAlunos() []Aluno {
	return r.alunos
}

func (r *AlunosRepositorio) criarAluno(matricula int, nome string, email string) Aluno {
	novoAluno := Aluno{
		Matrícula: matricula,
		Nome:      nome,
		Email:     email,
	}

	r.alunos = append(r.alunos, novoAluno)
	return novoAluno
}
