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

func TestGenerarRutinaErrorPorFaltaDeEjercicios(t *testing.T) {
	const TiempoValido = 1000

	_, err := GenerarRutina(TiempoValido)
	assert.Error(t, err)
	assert.Equal(t, "no hay suficientes ejercicios para llenar el tiempo disponible", err.Error())
}

func TestGenerarPlanSemanal(t *testing.T) {
	const TiempoValido = 10
	const DiasLibres = 3
	tiempoLibre := TiempoLibre{DiasLibresSemanales: DiasLibres, TiempoLibreDiario: TiempoValido}
	plan, err := GenerarPlanSemanal(tiempoLibre)
	assert.NoError(t, err)
	assert.Len(t, plan, DiasLibres)
}
