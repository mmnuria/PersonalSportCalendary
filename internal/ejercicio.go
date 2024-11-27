package PersonalSportCalendary

type TipoEjercicio int //Enum de tipos de ejercicios
const (
	Flexibilidad TipoEjercicio = iota
	Resistencia
	Fuerza
	Cardiovasculares
	Calentamiento
	Estiramiento
)

type Intensidad int //Enum de intensidades
const (
	Baja Intensidad = iota
	Media
	Alta
)

type Ejercicio struct {
	Nombre        string        //Nombre del ejercicio
	Descripcion   string        //Descripción del ejercicio
	MinsEstimados int           //Tiempo estimado en minutos para realizar el ejercicio
	TipoEjercicio TipoEjercicio //Grupo muscular al que pertenece el ejercicio
	Intensidad    Intensidad    //Intensidad del ejercicio
}

func NewEjercicio(nombre string, descripcion string, minsEstimados int, tipoEjercicio TipoEjercicio, intensidad Intensidad) Ejercicio {
	return Ejercicio{
		Nombre:        nombre,
		Descripcion:   descripcion,
		MinsEstimados: minsEstimados,
		TipoEjercicio: tipoEjercicio,
		Intensidad:    intensidad,
	}
}
