package PersonalSportCalendary

// PlanSemanal representa un plan semanal de rutinas y tiempo libre.
// - TiempoLibre: Tiempo libre del usuario durante la semana.
// - Rutinas: Rutinas de ejercicios que el usuario dispone para realizar durante la semana.
type PlanSemanal struct {
	TiempoLibre TiempoLibre
	Rutinas     []Rutina
}

// NewPlanSemanal crea una nueva instancia de PlanSemanal.
func NewPlanSemanal(tiempoLibre TiempoLibre, rutinas []Rutina) PlanSemanal {
	return PlanSemanal{
		TiempoLibre: tiempoLibre,
		Rutinas:     rutinas,
	}
}
