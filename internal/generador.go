package PersonalSportCalendary

import (
	"PersonalSportCalendary/logger"
	"errors"

	"go.uber.org/zap"
)

const MaxMinutosPorDia = 1440

func GuardarEjercicioSiValido(ejercicio Ejercicio, ejercicios *[]Ejercicio) {
	logger.InitLogger("test_config", "yaml", "../")

	logger.Logger.Info("Verificando si el ejercicio es válido", zap.String("ejercicio", ejercicio.Nombre))
	if len(ejercicio.Materiales) == 0 {
		*ejercicios = append(*ejercicios, ejercicio)
		logger.Logger.Info("Ejercicio agregado", zap.String("ejercicio", ejercicio.Nombre))
	} else {
		logger.Logger.Warn("Ejercicio no válido, tiene materiales asociados", zap.String("ejercicio", ejercicio.Nombre))
	}
}

func GenerarRutina(tiempoDisponible uint) (Rutina, error) {
	if tiempoDisponible <= 0 || tiempoDisponible > MaxMinutosPorDia {
		logger.Logger.Error("El tiempo disponible es inválido", zap.Uint("tiempoDisponible", tiempoDisponible))
		return Rutina{}, errors.New("el tiempo disponible debe estar entre 1 y 1440 minutos")
	}

	logger.Logger.Info("Generando rutina", zap.Uint("tiempoDisponible", tiempoDisponible))
	ejerciciosValidos := []Ejercicio{}
	tiempoRestante := tiempoDisponible

	for _, ejercicio := range ListaEjercicios {
		GuardarEjercicioSiValido(ejercicio, &ejerciciosValidos)
	}

	ejerciciosSeleccionados := []Ejercicio{}
	logger.Logger.Info("Seleccionando ejercicios para la rutina")

	for i := 0; tiempoRestante > 0 && i < len(ListaEjercicios); i++ {
		ejercicio := ListaEjercicios[i]
		logger.Logger.Debug("Seleccionando ejercicio", zap.String("ejercicio", ejercicio.Nombre), zap.Uint("tiempoRestante", tiempoRestante))
		tiempoRestante -= ejercicio.MinsEstimados
		ejerciciosSeleccionados = append(ejerciciosSeleccionados, ejercicio)
	}

	if tiempoRestante > 0 {
		logger.Logger.Error("No hay suficientes ejercicios para llenar el tiempo disponible", zap.Uint("tiempoRestante", tiempoRestante))
		return Rutina{}, errors.New("no hay suficientes ejercicios para llenar el tiempo disponible")
	}
	logger.Logger.Info("Rutina generada exitosamente", zap.Uint("tiempoDisponible", tiempoDisponible), zap.Int("totalEjercicios", len(ejerciciosSeleccionados)))
	return Rutina{
		Nombre:         "Rutina personalizada",
		TiempoDuracion: tiempoDisponible,
		Ejercicios:     ejerciciosSeleccionados,
		Materiales:     []string{},
	}, nil
}

func GenerarPlanSemanal(tiempoLibre TiempoLibre) (PlanSemanal, error) {
	logger.Logger.Info("Generando plan semanal", zap.Uint("diasLibres", tiempoLibre.DiasLibresSemanales), zap.Uint("tiempoLibreDiario", tiempoLibre.TiempoLibreDiario))
	var rutinas []Rutina
	for i := uint(0); i < tiempoLibre.DiasLibresSemanales; i++ {
		logger.Logger.Info("Generando rutina para el día", zap.Uint("dia", i+1))
		rutina, err := GenerarRutina(tiempoLibre.TiempoLibreDiario)
		if err != nil {
			logger.Logger.Error("Error al generar rutina", zap.Error(err))
			return PlanSemanal{}, err
		}
		rutinas = append(rutinas, rutina)
	}
	logger.Logger.Info("Plan semanal generado exitosamente", zap.Uint("diasLibres", tiempoLibre.DiasLibresSemanales))
	return NewPlanSemanal(tiempoLibre, rutinas)
}
