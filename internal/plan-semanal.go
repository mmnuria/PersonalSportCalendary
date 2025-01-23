package PersonalSportCalendary

import "fmt"

var ListaPlanesSemanales []PlanSemanal

type PlanSemanal []Rutina

func NewPlanSemanal(tiempoLibre TiempoLibre, rutinas []Rutina) (PlanSemanal, error) {
	if tiempoLibre.DiasLibresSemanales != uint(len(rutinas)) {
		return PlanSemanal{}, fmt.Errorf("el número de días libres (%d) no coincide con el número de rutinas (%d)", tiempoLibre.DiasLibresSemanales, len(rutinas))
	}
	for i, rutina := range rutinas {
		if rutina.TiempoDuracion != tiempoLibre.TiempoLibreDiario {
			return PlanSemanal{}, fmt.Errorf("la duración de la rutina %d (%d) no coincide con el tiempo libre diario (%d)", i, rutina.TiempoDuracion, tiempoLibre.TiempoLibreDiario)
		}
	}
	return PlanSemanal(rutinas), nil
}
