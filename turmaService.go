package main

import "errors"

type TurmaService struct {
	repositorio *TurmasRepositorio
}

func NovoTurmaService(repositorio *TurmasRepositorio) *TurmaService {
	return &TurmaService{
		repositorio: repositorio,
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
