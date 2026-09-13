package main

type Sala struct {
	ID         int      `json:"id"`
	Nome       string   `json:"nome"`
	Capacidade int      `json:"capacidade"`
	Recursos   []string `json:"recursos"`
}

type SalasRepositorio struct {
	salas []Sala
}

func NovoSalasRepositorio() *SalasRepositorio {
	return &SalasRepositorio{
		salas: []Sala{},
	}
}

func (r *SalasRepositorio) listarSalas() []Sala {
	return r.salas
}

func (r *SalasRepositorio) criarSala(id int, nome string, capacidade int, recursos []string) Sala {
	novaSala := Sala{
		ID:         id,
		Nome:       nome,
		Capacidade: capacidade,
		Recursos:   recursos,
	}
	r.salas = append(r.salas, novaSala)
	return novaSala
}

func (r *SalasRepositorio) buscarSalaPorID(id int) *Sala {
	for _, sala := range r.salas {
		if sala.ID == id {
			return &sala
		}
	}
	return nil
}
