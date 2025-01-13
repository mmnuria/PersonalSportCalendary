package PersonalSportCalendary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerarRutina(t *testing.T) {
	const TiempoValido = 10
	rutina, err := GenerarRutina(TiempoValido)
	assert.NoError(t, err)
	assert.Equal(t, uint(TiempoValido), rutina.TiempoDuracion)
	assert.Len(t, rutina.Ejercicios, 2)
}

func TestGenerarRutinaTiempoInvalido(t *testing.T) {
	const TiempoMinimo = 0
	const TiempoMaximo = MaxMinutosPorDia
	tiemposInvalidos := []uint{TiempoMinimo, TiempoMaximo + 1}

	for _, tiempo := range tiemposInvalidos {
		_, err := GenerarRutina(tiempo)
		assert.Error(t, err)
		assert.Equal(t, "el tiempo disponible debe estar entre 1 y 1440 minutos", err.Error())
	}
}
	assert.Error(t, err)
}

func TestGenerarPlanSemanal(t *testing.T) {
	tiempoLibre := TiempoLibre{DiasLibresSemanales: 3, TiempoLibreDiario: 10}
	plan, err := GenerarPlanSemanal(tiempoLibre)
	assert.NoError(t, err)
	assert.Len(t, plan, 3)
}
