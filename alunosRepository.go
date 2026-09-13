package main

type Aluno struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

type AlunosRepositorio struct {
	alunos    []Aluno
	proximoID int
}

func NovoAlunosRepositorio() *AlunosRepositorio {
	return &AlunosRepositorio{
		alunos:    []Aluno{},
		proximoID: 1,
	}
}

func (r *AlunosRepositorio) listarAlunos() []Aluno {
	return r.alunos
}

func (r *AlunosRepositorio) criarAluno(nome string, email string) Aluno {
	novoAluno := Aluno{
		ID:    r.proximoID,
		Nome:  nome,
		Email: email,
	}

	r.alunos = append(r.alunos, novoAluno)
	r.proximoID++
	return novoAluno
}
