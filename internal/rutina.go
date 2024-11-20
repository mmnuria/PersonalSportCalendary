package PersonalSportCalendary

type Rutina struct {
	Nombre     string   //Nombre de la rutina
	Tiempo     int      //Duración en minutos
	Ejercicios []string //Instrucciones detalladas sobre los ejercicios que componen la rutina y sus repeticiones
	Materiales []string //Material necesario para realizar cada ejercicio de la rutina
}

func NewRutina(nombre string, tiempo int, ejercicios []string, materiales []string) Rutina {
	return Rutina{
		Nombre:     nombre,
		Tiempo:     tiempo,
		Ejercicios: ejercicios,
		Materiales: materiales,
	}
}
