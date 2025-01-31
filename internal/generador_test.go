package PersonalSportCalendary

import (
	"testing"

	"PersonalSportCalendary/logger"

	. "github.com/onsi/gomega"
)

const TiempoValido = 10

func TestGenerarRutina(t *testing.T) {
	g := NewGomegaWithT(t)

	logger.InitLogger("test_config", "yaml", "../")

	logger.Logger.Infow("Ejecutando TestGenerarRutina", "TiempoValido", TiempoValido)

	rutina, err := GenerarRutina(TiempoValido)
	g.Expect(err).To(BeNil())
	g.Expect(rutina.TiempoDuracion).To(Equal(uint(TiempoValido)))

	totalTiempoEjercicios := uint(0)
	for _, ejercicio := range rutina.Ejercicios {
		totalTiempoEjercicios += ejercicio.MinsEstimados
	}
	logger.Logger.Infow("Rutina generada correctamente", "TotalEjercicios", len(rutina.Ejercicios), "TotalTiempoEjercicios", totalTiempoEjercicios)
	g.Expect(totalTiempoEjercicios).To(Equal(uint(TiempoValido)))
}

func TestGenerarRutinaTiempoInvalido(t *testing.T) {
	g := NewGomegaWithT(t)

	const TiempoMinimo = 0
	tiemposInvalidos := []uint{TiempoMinimo, MaxMinutosPorDia + 1}

	for _, tiempo := range tiemposInvalidos {
		logger.Logger.Warnw("Probando tiempo inválido", "TiempoIngresado", tiempo)
		_, err := GenerarRutina(tiempo)
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("el tiempo disponible debe estar entre 1 y 1440 minutos"))
	}
}

func TestGenerarRutinaErrorPorFaltaDeEjercicios(t *testing.T) {
	g := NewGomegaWithT(t)
	logger.Logger.Infow("Iniciando prueba de generación de rutina por falta de ejercicios", "TiempoValido", TiempoValido)
	const TiempoValido = 1000

	_, err := GenerarRutina(TiempoValido)
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("no hay suficientes ejercicios para llenar el tiempo disponible"))

	logger.Logger.Infow("TestGenerarRutinaErrorPorFaltaDeEjercicios completado", "TiempoValido", TiempoValido)
}

func TestGenerarPlanSemanal(t *testing.T) {
	g := NewGomegaWithT(t)

	const DiasLibres = 3
	tiempoLibre := TiempoLibre{DiasLibresSemanales: DiasLibres, TiempoLibreDiario: TiempoValido}
	logger.Logger.Infow("Iniciando generación del plan semanal", "DiasLibres", DiasLibres, "TiempoLibre", tiempoLibre)
	plan, err := GenerarPlanSemanal(tiempoLibre)
	g.Expect(err).To(BeNil())
	g.Expect(plan).To(HaveLen(DiasLibres))
	logger.Logger.Infow("Plan semanal generado correctamente", "DiasLibres", DiasLibres)
}

func TestGenerarPlanSemanalDuracion(t *testing.T) {
	g := NewGomegaWithT(t)

	const DiasLibres = 3
	tiempoLibre := TiempoLibre{DiasLibresSemanales: DiasLibres, TiempoLibreDiario: TiempoValido}
	logger.Logger.Infow("Iniciando generación del plan semanal para verificar la duración", "DiasLibres", DiasLibres, "TiempoLibre", tiempoLibre)
	plan, err := GenerarPlanSemanal(tiempoLibre)
	g.Expect(err).To(BeNil())

	for _, rutina := range plan {
		logger.Logger.Infow("Verificando duración de rutina", "RutinaDuracion", rutina.TiempoDuracion)
		g.Expect(rutina.TiempoDuracion).To(Equal(uint(TiempoValido)))
	}

	logger.Logger.Infow("Duración del plan semanal verificada correctamente", "DiasLibres", DiasLibres)
}

func TestGenerarPlanSemanalEjercicios(t *testing.T) {
	g := NewGomegaWithT(t)

	const DiasLibres = 3
	tiempoLibre := TiempoLibre{DiasLibresSemanales: DiasLibres, TiempoLibreDiario: TiempoValido}
	logger.Logger.Infow("Iniciando generación del plan semanal para verificar los ejercicios", "DiasLibres", DiasLibres, "TiempoLibre", tiempoLibre)
	plan, err := GenerarPlanSemanal(tiempoLibre)
	g.Expect(err).To(BeNil())

	for _, rutina := range plan {
		totalTiempoEjercicios := uint(0)
		for _, ejercicio := range rutina.Ejercicios {
			totalTiempoEjercicios += ejercicio.MinsEstimados
		}
		logger.Logger.Infow("Verificando tiempo total de ejercicios de rutina", "TotalTiempoEjercicios", totalTiempoEjercicios)
		g.Expect(totalTiempoEjercicios).To(Equal(uint(TiempoValido)))
	}
	logger.Logger.Infow("Ejercicios del plan semanal verificados correctamente", "DiasLibres", DiasLibres)
}
