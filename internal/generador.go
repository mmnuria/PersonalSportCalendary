package PersonalSportCalendary

import "errors"

// GenerarRutina genera una rutina de ejercicios basada en el tiempo disponible diario.
func GenerarRutina(tiempoDisponible uint) (Rutina, error) {
	if tiempoDisponible <= 0 || tiempoDisponible > 1440 {
		return Rutina{}, errors.New("el tiempo disponible debe estar entre 1 y 1440 minutos")
	}

	var ejerciciosSeleccionados []Ejercicio
	tiempoRestante := tiempoDisponible

	for _, ejercicio := range ListaEjercicios {
		if ejercicio.MinsEstimados <= tiempoRestante && len(ejercicio.Materiales) == 0 {
			ejerciciosSeleccionados = append(ejerciciosSeleccionados, ejercicio)
			tiempoRestante -= ejercicio.MinsEstimados
		}
	}

	if tiempoRestante > 0 {
		return Rutina{}, errors.New("no hay suficientes ejercicios para llenar el tiempo disponible")
	}

	return Rutina{
		Nombre:         "Rutina personalizada",
		TiempoDuracion: tiempoDisponible,
		Ejercicios:     ejerciciosSeleccionados,
		Materiales:     []string{},
	}, nil
}

// GenerarPlanSemanal crea un plan semanal basado en el tiempo libre del usuario.
func GenerarPlanSemanal(tiempoLibre TiempoLibre) (PlanSemanal, error) {
	var rutinas []Rutina
	for i := uint(0); i < tiempoLibre.DiasLibresSemanales; i++ {
		rutina, err := GenerarRutina(tiempoLibre.TiempoLibreDiario)
		if err != nil {
			return PlanSemanal{}, err
		}
		rutinas = append(rutinas, rutina)
	}

	return NewPlanSemanal(tiempoLibre, rutinas)
}
