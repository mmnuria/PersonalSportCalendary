package PersonalSportCalendary

import "fmt"

// ListaEjecicios guarda el conjunto de ejercicios que un usuario puede realizar.
type ListaEjercicios []Ejercicio

// Ejercicio representa un ejercicio que un usuario puede realizar.
// - Nombre: Nombre del ejercicio.
// - Descripcion: Descripción detallada del ejercicio.
// - MinsEstimados: Tiempo estimado en minutos para realizar el ejercicio (máximo 1440, un día).
// - ClaseEjercicio: Clase de ejercicio a realizar entre: Flexibilidad, Fuerza, Cardiovasculares, Calentamiento, Estiramiento.
// - Intensidad: Intensidad del ejercicio entre: Baja, Media, Alta.
type Ejercicio struct {
	Nombre         string
	Descripcion    string
	MinsEstimados  uint
	ClaseEjercicio string
	Intensidad     string
}

// NewEjercicio crea una nueva instancia de Ejercicio.
func NewEjercicio(nombre string, descripcion string, minsEstimados uint, claseEjercicio string, intensidad string) (Ejercicio, error) {
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
	}, nil
}
