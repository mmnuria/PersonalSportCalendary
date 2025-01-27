package PersonalSportCalendary

import (
	"testing"

	. "github.com/onsi/gomega"
)

const TiempoValido = 10
const DiasLibres = 3

func TestGenerarRutina(t *testing.T) {
	g := NewGomegaWithT(t)

	rutina, err := GenerarRutina(TiempoValido)
	g.Expect(err).To(BeNil())
	g.Expect(rutina.TiempoDuracion).To(Equal(uint(TiempoValido)))

	totalTiempoEjercicios := uint(0)
	for _, ejercicio := range rutina.Ejercicios {
		totalTiempoEjercicios += ejercicio.MinsEstimados
	}
	g.Expect(totalTiempoEjercicios).To(Equal(uint(TiempoValido)))
}

func TestGenerarRutinaTiempoInvalido(t *testing.T) {
	g := NewGomegaWithT(t)

	const TiempoMinimo = 0
	tiemposInvalidos := []uint{TiempoMinimo, MaxMinutosPorDia + 1}

	for _, tiempo := range tiemposInvalidos {
		_, err := GenerarRutina(tiempo)
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("el tiempo disponible debe estar entre 1 y 1440 minutos"))
	}
}

func TestGenerarRutinaErrorPorFaltaDeEjercicios(t *testing.T) {
	g := NewGomegaWithT(t)

	const TiempoValido = 1000

	_, err := GenerarRutina(TiempoValido)
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("no hay suficientes ejercicios para llenar el tiempo disponible"))
}

func TestGenerarPlanSemanal(t *testing.T) {
	g := NewGomegaWithT(t)

	tiempoLibre := TiempoLibre{DiasLibresSemanales: DiasLibres, TiempoLibreDiario: TiempoValido}
	plan, err := GenerarPlanSemanal(tiempoLibre)
	g.Expect(err).To(BeNil())
	g.Expect(plan).To(HaveLen(DiasLibres))
}

func TestGenerarPlanSemanalDuracion(t *testing.T) {
	g := NewGomegaWithT(t)

	tiempoLibre := TiempoLibre{DiasLibresSemanales: DiasLibres, TiempoLibreDiario: TiempoValido}
	plan, err := GenerarPlanSemanal(tiempoLibre)
	g.Expect(err).To(BeNil())

	for _, rutina := range plan {
		g.Expect(rutina.TiempoDuracion).To(Equal(uint(TiempoValido)))
	}
}

func TestGenerarPlanSemanalEjercicios(t *testing.T) {
	g := NewGomegaWithT(t)

	tiempoLibre := TiempoLibre{DiasLibresSemanales: DiasLibres, TiempoLibreDiario: TiempoValido}
	plan, err := GenerarPlanSemanal(tiempoLibre)
	g.Expect(err).To(BeNil())

	for _, rutina := range plan {
		totalTiempoEjercicios := uint(0)
		for _, ejercicio := range rutina.Ejercicios {
			totalTiempoEjercicios += ejercicio.MinsEstimados
		}
		g.Expect(totalTiempoEjercicios).To(Equal(uint(TiempoValido)))
	}
}
