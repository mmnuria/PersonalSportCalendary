package PersonalSportCalendary

type TiempoLibre struct {
	Días   int //Días libres a la semana
	Tiempo int //Tiempo en minutos
}

func NewTiempoLibre(días int, tiempo int) TiempoLibre {
	return TiempoLibre{
		Días:   días,
		Tiempo: tiempo,
	}
}
