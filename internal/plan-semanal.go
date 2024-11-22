package PersonalSportCalendary

type PlanSemanal struct {
	TiempoLibre TiempoLibre
	Rutinas     []Rutina
}

func NewPlanSemanal(tiempoLibre TiempoLibre, rutinas []Rutina) PlanSemanal {
	return PlanSemanal{
		TiempoLibre: tiempoLibre,
		Rutinas:     rutinas,
	}
}
