package PersonalSportCalendary

import "fmt"

type Ejercicio struct {
	Nombre         string //Nombre del ejercicio
	Descripcion    string //Descripción del ejercicio
	MinsEstimados  int    //Tiempo estimado en minutos para realizar el ejercicio
	ClaseEjercicio string //Clase de ejercicio a realizar entre: Flexibilidad, Fuerza, Cardiovasculares, Calentamiento, Estiramiento
	Intensidad     string //Intensidad del ejercicio entre: Baja, Media, Alta
}

func NewEjercicio(nombre string, descripcion string, minsEstimados int, claseEjercicio string, intensidad string) (Ejercicio, error) {
	//Comprobación de clase de ejercicio
	switch string(claseEjercicio) {
	case "Flexibilidad", "Fuerza", "Cardiovasculares", "Calentamiento", "Estiramiento":
		break
	default:
		return Ejercicio{}, fmt.Errorf("wjercicio inválido: %v", claseEjercicio)
	}
	//Comprobación de intensidad
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
