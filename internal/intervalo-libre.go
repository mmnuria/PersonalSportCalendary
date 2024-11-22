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

type Momento struct {
	Dia  Dia    // Día de la semana (enum)
	Hora string // Hora en formato "HH:MM"
}

func NewIntervaloHorario(dia Dia, hora string) Momento {
	return Momento{
		Dia:  dia,
		Hora: hora,
	}
}

type IntervaloLibre struct {
	Inicio Momento // Momento de inicio del tiempo libre
	Fin    Momento // Momento de fin del tiempo libre
}

func NewIntervaloLibre(inicio Momento, fin Momento) IntervaloLibre {
	return IntervaloLibre{
		Inicio: inicio,
		Fin:    fin,
	}
}
