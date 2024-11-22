package PersonalSportCalendary

type Material int //enum de materiales

const (
	Barbell = iota
	Dumbbells
	Bodyweight
	Machine
	MedicineBall
	Kettlebells
	Stretches
	Cables
	Band
	Plate
	TRX
	Yoga
	BosuBall
	Vitruvian
	Cardio
	SmithMachine
	Recovery
)

type Rutina struct {
	Nombre     string      //Nombre de la rutina
	Tiempo     int         //Duración en minutos
	Ejercicios []Ejercicio //Instrucciones detalladas sobre los ejercicios que componen la rutina y sus repeticiones
	Materiales []Material  //Materiales necesarios para realizar la rutina
}

func NewRutina(nombre string, tiempo int, ejercicios []Ejercicio, materiales []Material) Rutina {
	return Rutina{
		Nombre:     nombre,
		Tiempo:     tiempo,
		Ejercicios: ejercicios,
		Materiales: materiales,
	}
}
