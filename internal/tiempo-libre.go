package PersonalSportCalendary

type Dia int //Enum de días (para poder referenciarlos por su nombre)
const (
	Lunes Dia = iota
	Martes
	Miércoles
	Jueves
	Viernes
	Sábado
	Domingo
)

type IntervaloHorario struct { //Intervalos de tiempo libre (formato de los string "HH:MM")
	HorarioInicio string
	HorarioFin    string
}

type TiempoLibre struct {
	Dias     []Dia                      //Días como enum
	Horarios map[Dia][]IntervaloHorario //Mapeamos los horarios de disponibilidad a su día correspondiente
}
