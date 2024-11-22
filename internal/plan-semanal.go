package PersonalSportCalendary

type PlanIntervalo struct {
	IntervaloTiempo IntervaloTiempo
	Rutina          Rutina
}

func NewPlanIntervalo(intervaloTiempo IntervaloTiempo, rutina Rutina) PlanIntervalo {
	return PlanIntervalo{
		IntervaloTiempo: intervaloTiempo,
		Rutina:          rutina,
	}
}

// relaciona un tiempo libre con una rutina
type PlanSemanal struct {
	PlanIntervalos []PlanIntervalo
}

func NewPlanSemanal(planIntervalo []PlanIntervalo) PlanSemanal {
	return PlanSemanal{
		PlanIntervalos: planIntervalo,
	}
}
