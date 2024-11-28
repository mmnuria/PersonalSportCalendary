package PersonalSportCalendary

import "fmt"

type Rutina struct {
	Nombre     string      //Nombre de la rutina
	Tiempo     int         //Duración en minutos
	Ejercicios []Ejercicio //Instrucciones detalladas sobre los ejercicios que componen la rutina y sus repeticiones
	Materiales []string    //Materiales necesarios para realizar la rutina entre: Peso, BandaElastica, Cuerda, Esterilla, Pelota
}

func NewRutina(nombre string, tiempo int, ejercicios []Ejercicio, materiales []string) (Rutina, error) {
	//Comprobación de materiales válidos
	for _, m := range materiales {
		switch string(m) {
		case "Peso", "BandaElastica", "Cuerda", "Esterilla", "Pelota":
		default:
			return Rutina{}, fmt.Errorf("material inválido: %v", m)
		}
	}
	return Rutina{
		Nombre:     nombre,
		Tiempo:     tiempo,
		Ejercicios: ejercicios,
		Materiales: materiales,
	}, nil
}
