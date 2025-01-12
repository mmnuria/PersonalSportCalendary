package PersonalSportCalendary

import "fmt"

// ListaEjecicios guarda el conjunto de ejercicios que un usuario puede realizar.
var ListaEjercicios []Ejercicio

// Ejercicio representa un ejercicio que un usuario puede realizar.
// - Nombre: Nombre del ejercicio.
// - Descripcion: Descripción detallada del ejercicio.
// - MinsEstimados: Tiempo estimado en minutos para realizar el ejercicio (máximo 1440, un día).
// - ClaseEjercicio: Clase de ejercicio a realizar entre: Flexibilidad, Fuerza, Cardiovasculares, Calentamiento, Estiramiento.
// - Intensidad: Intensidad del ejercicio entre: Baja, Media, Alta.
// - Materiales: Especifica una lista de materiales necesarios para realizar ese ejercicio, puede no necesitar ninguno y estar vacía
type Ejercicio struct {
	Nombre         string
	Descripcion    string
	MinsEstimados  uint
	ClaseEjercicio string
	Intensidad     string
	Materiales     []string
}

// init es una función especial que se ejecuta al inicio del paquete para inicializar las variables.
func init() {
	// Inicializamos ListaEjercicios con algunos ejercicios de ejemplo
	ListaEjercicios = []Ejercicio{
		{"Saltos", "Saltos en el lugar", 5, "Cardiovasculares", "Media", []string{}},
		{"Flexiones", "Flexiones estándar", 10, "Fuerza", "Alta", []string{}},
		{"Estiramientos", "Estiramiento básico de brazos", 5, "Estiramiento", "Baja", []string{}},
	}
}

// NewEjercicio crea una nueva instancia de Ejercicio.
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
