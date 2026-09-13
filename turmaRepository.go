package main

type Alocacao struct {
	SalaID        int    `json:"sala_id"`
	DiaSemana     string `json:"dia_semana"`
	HorarioInicio string `json:"horario_inicio"`
	HorarioFim    string `json:"horario_fim"`
}

type Turma struct {
	ID         int       `json:"id"`
	Nome       string    `json:"nome"`
	Disciplina string    `json:"disciplina"`
	Professor  string    `json:"professor"`
	AlunoIDs   []int     `json:"aluno_ids"`
	Alocacao   *Alocacao `json:"alocacao,omitempty"`
}

type TurmasRepositorio struct {
	turmas []Turma
}

func NovoTurmasRepositorio() *TurmasRepositorio {
	return &TurmasRepositorio{
		turmas: []Turma{},
	}
}

func (r *TurmasRepositorio) listarTurmas() []Turma {
	return r.turmas
}

func (r *TurmasRepositorio) criarTurma(id int, nome string, disciplina string, professor string) Turma {
	novaTurma := Turma{
		ID:         id,
		Nome:       nome,
		Disciplina: disciplina,
		Professor:  professor,
		AlunoIDs:   []int{},
	}
	r.turmas = append(r.turmas, novaTurma)
	return novaTurma
}

func (r *TurmasRepositorio) buscarTurmaPorID(id int) *Turma {
	for i := range r.turmas {
		if r.turmas[i].ID == id {
			return &r.turmas[i]
		}
	}
	return nil
}
