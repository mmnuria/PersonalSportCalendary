package PersonalSportCalendary

import "fmt"

var ListaRutinas []Rutina

type Rutina struct {
	Nombre         string
	TiempoDuracion uint
	Ejercicios     []Ejercicio
	Materiales     []string
}

func NewRutina(nombre string, tiempoDuracion uint, ejercicios []Ejercicio, materiales []string) (Rutina, error) {
	if tiempoDuracion <= 0 || tiempoDuracion > 1440 {
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
