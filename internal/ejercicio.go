package PersonalSportCalendary

type Tipo int //Enum de grupos musculares (obtenidos a partir de musclewiki.com)
const (
	Flexibilidad Tipo = iota
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
	Nombre        string     //Nombre del ejercicio
	Descripcion   string     //Descripción del ejercicio
	MinsEstimados int        //Tiempo estimado en minutos para realizar el ejercicio
	Tipo          Tipo       //Grupo muscular al que pertenece el ejercicio
	Intensidad    Intensidad //Intensidad del ejercicio
}
