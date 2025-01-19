package PersonalSportCalendary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const TiempoValido = 10

func TestGenerarRutina(t *testing.T) {
	rutina, err := GenerarRutina(TiempoValido)
	assert.NoError(t, err)
	assert.Equal(t, uint(TiempoValido), rutina.TiempoDuracion)

	totalTiempoEjercicios := uint(0)
	for _, ejercicio := range rutina.Ejercicios {
		totalTiempoEjercicios += ejercicio.MinsEstimados
	}
	assert.Equal(t, totalTiempoEjercicios, uint(TiempoValido))
}

func TestGenerarRutinaTiempoInvalido(t *testing.T) {
	const TiempoMinimo = 0
	tiemposInvalidos := []uint{TiempoMinimo, MaxMinutosPorDia + 1}

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
	const DiasLibres = 3
	tiempoLibre := TiempoLibre{DiasLibresSemanales: DiasLibres, TiempoLibreDiario: TiempoValido}
	plan, err := GenerarPlanSemanal(tiempoLibre)
	assert.NoError(t, err)
	assert.Len(t, plan, DiasLibres)
}

func TestGenerarPlanSemanalDuracion(t *testing.T) {
	const DiasLibres = 3
	tiempoLibre := TiempoLibre{DiasLibresSemanales: DiasLibres, TiempoLibreDiario: TiempoValido}
	plan, err := GenerarPlanSemanal(tiempoLibre)
	assert.NoError(t, err)

	for _, rutina := range plan {
		assert.Equal(t, uint(TiempoValido), rutina.TiempoDuracion)
	}
}

func TestGenerarPlanSemanalEjercicios(t *testing.T) {
	const DiasLibres = 3
	tiempoLibre := TiempoLibre{DiasLibresSemanales: DiasLibres, TiempoLibreDiario: TiempoValido}
	plan, err := GenerarPlanSemanal(tiempoLibre)
	assert.NoError(t, err)

	for _, rutina := range plan {
		totalTiempoEjercicios := uint(0)
		for _, ejercicio := range rutina.Ejercicios {
			totalTiempoEjercicios += ejercicio.MinsEstimados
		}
		assert.Equal(t, totalTiempoEjercicios, uint(TiempoValido))
	}
}
