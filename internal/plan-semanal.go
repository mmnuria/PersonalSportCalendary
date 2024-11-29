package PersonalSportCalendary

import "fmt"

// PlanSemanal representa el plan semanal de rutinas de un usuario con un determinado tiempo libre.
// Se comprueba que haya una rutina por día y que coincida la duración de cada rutina con el tiempo libre diario.
type PlanSemanal []Rutina

// NewPlanSemanal crea una nueva instancia de PlanSemanal.
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
