package PersonalSportCalendary

import "errors"

const MaxMinutosPorDia = 1440

func EsEjercicioValido(ejercicio Ejercicio, tiempoRestante uint) bool {
	return ejercicio.MinsEstimados <= tiempoRestante && len(ejercicio.Materiales) == 0
}

func filtrarEjerciciosValidos(ejercicios []Ejercicio, tiempoRestante uint) []Ejercicio {
	var ejerciciosValidos []Ejercicio
	for _, ejercicio := range ejercicios {
		if EsEjercicioValido(ejercicio, tiempoRestante) {
			ejerciciosValidos = append(ejerciciosValidos, ejercicio)
		}
	}
	return ejerciciosValidos
}

func GenerarRutina(tiempoDisponible uint) (Rutina, error) {
	if tiempoDisponible <= 0 || tiempoDisponible > MaxMinutosPorDia {
		return Rutina{}, errors.New("el tiempo disponible debe estar entre 1 y 1440 minutos")
	}

	ejerciciosSeleccionados := []Ejercicio{}
	tiempoRestante := tiempoDisponible

	ejerciciosFiltrados := filtrarEjerciciosValidos(ListaEjercicios, tiempoRestante)

	for _, ejercicio := range ejerciciosFiltrados {
		if ejercicio.MinsEstimados <= tiempoRestante {
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
