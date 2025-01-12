package PersonalSportCalendary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerarRutina(t *testing.T) {
	rutina, err := GenerarRutina(10)
	assert.NoError(t, err)
	assert.Equal(t, uint(10), rutina.TiempoDuracion)
	assert.Len(t, rutina.Ejercicios, 2)
}

func TestGenerarRutinaInsuficienteTiempo(t *testing.T) {
	_, err := GenerarRutina(1000)
	assert.Error(t, err)
}

func TestGenerarPlanSemanal(t *testing.T) {
	tiempoLibre := TiempoLibre{DiasLibresSemanales: 3, TiempoLibreDiario: 10}
	plan, err := GenerarPlanSemanal(tiempoLibre)
	assert.NoError(t, err)
	assert.Len(t, plan, 3)
}
