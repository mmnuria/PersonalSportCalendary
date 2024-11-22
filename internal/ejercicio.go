package PersonalSportCalendary

type GrupoMuscular int //Enum de grupos musculares (obtenidos a partir de musclewiki.com)
const (
	Arms GrupoMuscular = iota
	Forearms
	Biceps
	Triceps
	FrontShoulders
	RearShoulders
	Traps
	Chest
	Lats
	Obliques
	Abdominals
	Glutes
	Quads
	Harmstrings
	Calves
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
	GrupoMuscular GrupoMuscular //Grupo muscular al que pertenece el ejercicio
	Intensidad    Intensidad    //Intensidad del ejercicio
}
