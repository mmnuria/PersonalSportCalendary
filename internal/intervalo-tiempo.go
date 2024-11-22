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

func NewMomento(dia Dia, hora string) Momento {
	return Momento{
		Dia:  dia,
		Hora: hora,
	}
}

type IntervaloTiempo struct {
	Inicio Momento // Momento de inicio del tiempo libre
	Fin    Momento // Momento de fin del tiempo libre
}

func NewIntervaloTiempo(inicio Momento, fin Momento) IntervaloTiempo {
	return IntervaloTiempo{
		Inicio: inicio,
		Fin:    fin,
	}
}
