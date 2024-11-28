package PersonalSportCalendary

import "fmt"

// TiempoLibre representa la cantidad de tiempo libre de un usuario durante la semana.
// - DiasLibresSemanales: Número de días libres por semana (valor menor que 7).
// - TiempoLibreDiario: Tiempo libre diario en minutos (máximo 1440).
type TiempoLibre struct {
	DiasLibresSemanales uint
	TiempoLibreDiario   uint
}

// NewTiempoLibre crea una nueva instancia de TiempoLibre.
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
