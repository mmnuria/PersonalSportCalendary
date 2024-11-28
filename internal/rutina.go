package PersonalSportCalendary

import "fmt"

// Rutina representa una rutina de ejercicios que un usuario puede realizar.
// - Nombre: Nombre de la rutina.
// - Tiempo: Duración en minutos de la rutina.
// - Ejercicios: Instrucciones detalladas sobre los ejercicios que componen la rutina y sus repeticiones.
// - Materiales: Materiales necesarios para realizar la rutina entre: Peso, BandaElastica, Cuerda, Ester
type Rutina struct {
	Nombre         string
	TiempoDuracion uint
	Ejercicios     []Ejercicio
	Materiales     []string
}

// NewRutina crea una nueva instancia de Rutina.
func NewRutina(nombre string, tiempoDuracion uint, ejercicios []Ejercicio, materiales []string) (Rutina, error) {
	if tiempoDuracion > 1440 {
		return Rutina{}, fmt.Errorf("el tiempo debe estar entre 1 y 1440 minutos, recibido: %d", tiempoDuracion)
	}
	for _, m := range materiales {
		switch string(m) {
		case "Peso", "BandaElastica", "Cuerda", "Esterilla", "Pelota":
		default:
			return Rutina{}, fmt.Errorf("material inválido: %v", m)
		}
	}
	return Rutina{
		Nombre:         nombre,
		TiempoDuracion: tiempoDuracion,
		Ejercicios:     ejercicios,
		Materiales:     materiales,
	}, nil
}
