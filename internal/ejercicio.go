package PersonalSportCalendary

import "fmt"

type Ejercicio struct {
	Nombre        string         //Nombre del ejercicio
	Descripcion   string         //Descripción del ejercicio
	MinsEstimados int            //Tiempo estimado en minutos para realizar el ejercicio
	TipoEjercicio ClaseEjercicio //Grupo muscular al que pertenece el ejercicio
	Intensidad    Intensidad     //Intensidad del ejercicio
}

func NewEjercicio(nombre string, descripcion string, minsEstimados int, tipoEjercicio ClaseEjercicio, intensidad Intensidad) Ejercicio {
	return Ejercicio{
		Nombre:        nombre,
		Descripcion:   descripcion,
		MinsEstimados: minsEstimados,
		TipoEjercicio: tipoEjercicio,
		Intensidad:    intensidad,
	}
}

type ClaseEjercicio string //Enum de tipos de ejercicios
const (
	Flexibilidad     = "Flexibilidad"
	Fuerza           = "Fuerza"
	Cardiovasculares = "Cardiovasculares"
	Calentamiento    = "Calentamiento"
	Estiramiento     = "Estiramiento"
)

func NewClaseEjercicio(claseEjercicio string) (ClaseEjercicio, error) {
	switch ClaseEjercicio(claseEjercicio) {
	case Flexibilidad, Fuerza, Cardiovasculares, Calentamiento, Estiramiento:
		return ClaseEjercicio(claseEjercicio), nil
	default:
		return "", fmt.Errorf("Material inválido: %v", claseEjercicio)
	}
}

type Intensidad string //Enum de intensidades
const (
	Baja  = "Baja"
	Media = "Media"
	Alta  = "Alta"
)

func NewIntensidad(intensidad string) (Intensidad, error) {
	switch Intensidad(intensidad) {
	case Baja, Media, Alta:
		return Intensidad(intensidad), nil
	default:
		return "", fmt.Errorf("Intensidad inválida: %v", intensidad)
	}
}
