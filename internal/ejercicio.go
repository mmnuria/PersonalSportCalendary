package PersonalSportCalendary

import "fmt"

var ListaEjercicios []Ejercicio

const (
	ClaseFlexibilidad  = "Flexibilidad"
	ClaseFuerza        = "Fuerza"
	ClaseCardio        = "Cardiovasculares"
	ClaseCalentamiento = "Calentamiento"
	ClaseEstiramiento  = "Estiramiento"
)

const (
	IntensidadBaja  = "Baja"
	IntensidadMedia = "Media"
	IntensidadAlta  = "Alta"
)

type Ejercicio struct {
	Nombre         string
	Descripcion    string
	MinsEstimados  uint
	ClaseEjercicio string
	Intensidad     string
	Materiales     [] string
}

func init() {
	ListaEjercicios = []Ejercicio{
		{"Calentamiento", "Estiramiento general de extremidades", 5, ClaseCalentamiento, IntensidadBaja, []string{}},
		{"Saltos", "Saltos en el lugar", 5, ClaseCardio, IntensidadMedia, []string{}},
		{"Flexiones", "Flexiones estándar", 10, ClaseFuerza, IntensidadAlta, []string{}},
		{"Estiramientos", "Estiramiento básico de brazos", 5, ClaseEstiramiento, IntensidadBaja, []string{}},
	}
}

func NewEjercicio(nombre string, descripcion string, minsEstimados uint, claseEjercicio string, intensidad string, materiales []string) (Ejercicio, error) {
	if minsEstimados <= 0 || minsEstimados > 1440 {
		return Ejercicio{}, fmt.Errorf("el tiempo debe estar entre 1 y 1440 minutos, recibido: %d", minsEstimados)
	}
	switch string(claseEjercicio) {
	case "Flexibilidad", "Fuerza", "Cardiovasculares", "Calentamiento", "Estiramiento":
		break
	default:
		return Ejercicio{}, fmt.Errorf("Ejercicio inválido: %v", claseEjercicio)
	}
	switch string(intensidad) {
	case "Baja", "Media", "Alta":
		break
	default:
		return Ejercicio{}, fmt.Errorf("intensidad inválida: %v", intensidad)
	}
	return Ejercicio{
		Nombre:         nombre,
		Descripcion:    descripcion,
		MinsEstimados:  minsEstimados,
		ClaseEjercicio: claseEjercicio,
		Intensidad:     intensidad,
		Materiales:     materiales,
	}, nil
}
