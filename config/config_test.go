package config

import (
	"PersonalSportCalendary/logger"
	"testing"

	"github.com/onsi/gomega"
	"github.com/spf13/viper"
)

func TestConfigInitialization(t *testing.T) {
	viper.SetConfigName("test_config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../")

	if err := viper.ReadInConfig(); err != nil {
		t.Fatalf("Error al leer archivo de configuración de prueba: %s", err)
	}
	level := GetLogLevel()
	logFile := GetLogFile()
	encoding := GetLogEncoding()

	logger.InitLogger(level, logFile, encoding)
	g := gomega.NewGomegaWithT(t)

	t.Run("Test configuration values", func(t *testing.T) {
		logLevel := GetLogLevel()
		logFile := GetLogFile()
		encoding := GetLogEncoding()

		g.Expect(logLevel).To(gomega.Equal("info"))
		g.Expect(logFile).To(gomega.Equal("app.log"))
		g.Expect(encoding).To(gomega.Equal("console"))

		g.Expect(logger.Logger).NotTo(gomega.BeNil())

		logger.Logger.Info("Configuración cargada correctamente")
	})
}
