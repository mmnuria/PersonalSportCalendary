package PersonalSportCalendary

type TiempoLibre struct {
	Dia           string   // Día de la semana
	HorarioInicio string   // Hora de inicio, formato "HH:MM"
	HorarioFin    string   // Hora de fin, formato "HH:MM"
	Restricciones []string // Restricciones opcionales
}
