package PersonalSportCalendary

//relaciona un tiempo libre con una rutina
type PlanDiario struct {
	TiempoLibre TiempoLibre
	Rutina      []Rutina
}

func NewPlanDiario(tiempoLibre TiempoLibre, rutina []Rutina) PlanDiario {
	return PlanDiario{
		TiempoLibre: tiempoLibre,
		Rutina:      rutina,
	}
}
