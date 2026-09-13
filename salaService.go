package main

import "errors"

type SalaService struct {
	repositorio *SalasRepositorio
}

func NovoSalaService(repositorio *SalasRepositorio) *SalaService {
	return &SalaService{
		repositorio: repositorio,
	}
}

func (s *SalaService) ListarSalas() []Sala {
	return s.repositorio.listarSalas()
}

func (s *SalaService) CriarSala(id int, nome string, capacidade int, recursos []string) (Sala, error) {

	if id <= 0 || nome == "" || capacidade <= 0 || len(recursos) == 0 {
		return Sala{}, errors.New("Dados inválidos para criação da sala, revise-os")
	}

	if s.repositorio.buscarSalaPorID(id) != nil {
		return Sala{}, errors.New("Sala já existe")
	}

	return s.repositorio.criarSala(id, nome, capacidade, recursos), nil
}

func (s *SalaService) BuscarSalaPorID(id int) (*Sala, error) {
	if s.repositorio.buscarSalaPorID(id) == nil {
		return nil, errors.New("Sala não encontrada")
	}
	return s.repositorio.buscarSalaPorID(id), nil
}
