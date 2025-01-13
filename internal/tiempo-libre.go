package PersonalSportCalendary

import "fmt"

type TiempoLibre struct {
	DiasLibresSemanales uint
	TiempoLibreDiario   uint
}

func NewTiempoLibre(diasLibresSemanales uint, tiempoLibreDiario uint) (TiempoLibre, error) {
	if tiempoLibreDiario <= 0 || tiempoLibreDiario > 1440 {
		return TiempoLibre{}, fmt.Errorf("el tiempo debe estar entre 1 y 1440 minutos, recibido: %d", tiempoLibreDiario)
	}
	if diasLibresSemanales > 7 {
		return TiempoLibre{}, fmt.Errorf("no se puede superar los 7 días de la semana, recibido: %d", diasLibresSemanales)
	}

	return TiempoLibre{
		DiasLibresSemanales: diasLibresSemanales,
		TiempoLibreDiario:   tiempoLibreDiario,
	}, nil
}
