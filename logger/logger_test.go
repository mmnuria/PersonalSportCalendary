package logger

import (
	"PersonalSportCalendary/config"
	"bytes"
	"testing"

	"github.com/onsi/gomega"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestLoggerInitialization(t *testing.T) {
	config.InitConfig("test_config", "../", "yaml")

	g := gomega.NewGomegaWithT(t)

	level := config.GetLogLevel()
	logFile := config.GetLogFile()
	encoding := config.GetLogEncoding()
	InitLogger(level, logFile, encoding)

	g.Expect(Logger).NotTo(gomega.BeNil())

	var buf bytes.Buffer
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(&buf),
		zapcore.DebugLevel,
	)
	testLogger := zap.New(core)

	testLogger.Info("Mensaje de prueba")

	g.Expect(buf.String()).To(gomega.ContainSubstring("Mensaje de prueba"))
	Logger.Info("TestLoggerInitialization finalizado correctamente")
}
