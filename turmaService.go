package main

import (
	"errors"
	"strings"
)

type TurmaService struct {
	repositorio       *TurmasRepositorio
	salasRepositorio  *SalasRepositorio
	alunosRepositorio *AlunosRepositorio
}

func NovoTurmaService(repositorio *TurmasRepositorio, salasRepositorio *SalasRepositorio, alunosRepositorio *AlunosRepositorio) *TurmaService {
	return &TurmaService{
		repositorio:       repositorio,
		salasRepositorio:  salasRepositorio,
		alunosRepositorio: alunosRepositorio,
	}
}

func (t *TurmaService) ListarTurmas() []Turma {
	return t.repositorio.listarTurmas()
}

func (t *TurmaService) CriarTurma(id int, nome string, disciplina string, professor string) (Turma, error) {
	if id <= 0 || nome == "" || disciplina == "" || professor == "" {
		return Turma{}, errors.New("Dados inválidos para criação da turma, revise-os")
	}

	if t.repositorio.buscarTurmaPorID(id) != nil {
		return Turma{}, errors.New("Turma já existe")
	}

	return t.repositorio.criarTurma(id, nome, disciplina, professor), nil
}

func (t *TurmaService) BuscarTurmaPorID(id int) (*Turma, error) {
	turma := t.repositorio.buscarTurmaPorID(id)
	if turma == nil {
		return nil, errors.New("Turma não encontrada")
	}
	return turma, nil
}

func (t *TurmaService) AlocarSala(turmaID int, salaID int, diaSemana string, horarioInicio string, horarioFim string) (*Turma, error) {
	turma := t.repositorio.buscarTurmaPorID(turmaID)
	if turma == nil {
		return nil, errors.New("Turma não encontrada")
	}

	sala := t.salasRepositorio.buscarSalaPorID(salaID)
	if sala == nil {
		return nil, errors.New("Sala não encontrada")
	}

	if sala.Capacidade < len(turma.AlunoIDs) {
		return nil, errors.New("Capacidade da sala é insuficiente para os alunos já matriculados")
	}

	for _, outra := range t.repositorio.listarTurmas() {
		if outra.ID == turmaID || outra.Alocacao == nil {
			continue
		}
		if outra.Alocacao.SalaID != salaID || !strings.EqualFold(outra.Alocacao.DiaSemana, diaSemana) {
			continue
		}
		sobrepoe, err := horariosSobrepoem(horarioInicio, horarioFim, outra.Alocacao.HorarioInicio, outra.Alocacao.HorarioFim)
		if err != nil {
			return nil, err
		}
		if sobrepoe {
			return nil, errors.New("Sala já possui outra turma alocada nesse dia e horário")
		}
	}

	for _, alunoID := range turma.AlunoIDs {
		for _, outra := range t.repositorio.listarTurmas() {
			if outra.ID == turmaID || outra.Alocacao == nil {
				continue
			}
			if !contemAluno(outra.AlunoIDs, alunoID) || !strings.EqualFold(outra.Alocacao.DiaSemana, diaSemana) {
				continue
			}
			sobrepoe, err := horariosSobrepoem(horarioInicio, horarioFim, outra.Alocacao.HorarioInicio, outra.Alocacao.HorarioFim)
			if err != nil {
				return nil, err
			}
			if sobrepoe {
				return nil, errors.New("Conflito de agenda: um aluno da turma já tem aula em outra turma nesse horário")
			}
		}
	}

	turma.Alocacao = &Alocacao{
		SalaID:        salaID,
		DiaSemana:     diaSemana,
		HorarioInicio: horarioInicio,
		HorarioFim:    horarioFim,
	}
	return turma, nil
}

func (t *TurmaService) MatricularAluno(turmaID int, alunoID int) (*Turma, error) {
	turma := t.repositorio.buscarTurmaPorID(turmaID)
	if turma == nil {
		return nil, errors.New("Turma não encontrada")
	}

	if t.alunosRepositorio.buscarAlunoPorMatricula(alunoID) == nil {
		return nil, errors.New("Aluno não encontrado")
	}

	if contemAluno(turma.AlunoIDs, alunoID) {
		return nil, errors.New("Aluno já matriculado nesta turma")
	}

	if turma.Alocacao != nil {
		sala := t.salasRepositorio.buscarSalaPorID(turma.Alocacao.SalaID)
		if sala == nil {
			return nil, errors.New("Sala não encontrada")
		}
		if len(turma.AlunoIDs)+1 > sala.Capacidade {
			return nil, errors.New("Capacidade da sala excedida")
		}
	}

	if turma.Alocacao != nil {
		for _, outra := range t.repositorio.listarTurmas() {
			if outra.ID == turmaID || outra.Alocacao == nil {
				continue
			}
			if !contemAluno(outra.AlunoIDs, alunoID) || !strings.EqualFold(outra.Alocacao.DiaSemana, turma.Alocacao.DiaSemana) {
				continue
			}
			sobrepoe, err := horariosSobrepoem(
				turma.Alocacao.HorarioInicio, turma.Alocacao.HorarioFim,
				outra.Alocacao.HorarioInicio, outra.Alocacao.HorarioFim,
			)
			if err != nil {
				return nil, err
			}
			if sobrepoe {
				return nil, errors.New("Conflito de agenda: aluno já tem aula em outra turma nesse horário")
			}
		}
	}

	turma.AlunoIDs = append(turma.AlunoIDs, alunoID)
	return turma, nil
}

func (t *TurmaService) ListarAlunosDaTurma(turmaID int) ([]Aluno, error) {
	turma := t.repositorio.buscarTurmaPorID(turmaID)
	if turma == nil {
		return nil, errors.New("Turma não encontrada")
	}

	alunos := make([]Aluno, 0, len(turma.AlunoIDs))
	for _, alunoID := range turma.AlunoIDs {
		if aluno := t.alunosRepositorio.buscarAlunoPorMatricula(alunoID); aluno != nil {
			alunos = append(alunos, *aluno)
		}
	}
	return alunos, nil
}

func contemAluno(alunoIDs []int, alunoID int) bool {
	for _, id := range alunoIDs {
		if id == alunoID {
			return true
		}
	}
	return false
}
