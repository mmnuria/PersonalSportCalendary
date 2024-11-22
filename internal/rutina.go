package PersonalSportCalendary

type Rutina struct {
	Nombre     string      //Nombre de la rutina
	Tiempo     int         //Duración en minutos
	Ejercicios []Ejercicio //Instrucciones detalladas sobre los ejercicios que componen la rutina y sus repeticiones
}

func NewRutina(nombre string, tiempo int, ejercicios []Ejercicio, materiales []string) Rutina {
	return Rutina{
		Nombre:     nombre,
		Tiempo:     tiempo,
		Ejercicios: ejercicios,
	}
}
