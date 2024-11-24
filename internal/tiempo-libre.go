package PersonalSportCalendary

import "fmt"

type TiempoLibre struct {
	Días   int //Días libres a la semana
	Tiempo int //Tiempo en minutos
}

func NewTiempoLibre(días int, tiempo int) TiempoLibre {
	if tiempo <= 0 || tiempo > 1440 {
		fmt.Printf("El tiempo debe estar entre 1 y 1440 minutos, recibido: %d", tiempo)
		return TiempoLibre{}
	}

	return TiempoLibre{
		Días:   días,
		Tiempo: tiempo,
	}
}
