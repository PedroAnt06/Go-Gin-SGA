package main

import "time"

func parseHorario(hhmm string) (int, error) {
	t, err := time.Parse("15:04", hhmm)
	if err != nil {
		return 0, err
	}
	return t.Hour()*60 + t.Minute(), nil
}

func horariosSobrepoem(inicioA, fimA, inicioB, fimB string) (bool, error) {
	iA, err := parseHorario(inicioA)
	if err != nil {
		return false, err
	}
	fA, err := parseHorario(fimA)
	if err != nil {
		return false, err
	}
	iB, err := parseHorario(inicioB)
	if err != nil {
		return false, err
	}
	fB, err := parseHorario(fimB)
	if err != nil {
		return false, err
	}

	return iA < fB && fA > iB, nil
}
