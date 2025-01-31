package config

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfig(configName, configPath, configType string) {
	viper.SetConfigName(configName) // Usar el nombre del archivo de configuración
	viper.SetConfigType(configType) // Tipo de archivo (por ejemplo, "yaml")
	viper.AddConfigPath(configPath) // Ruta donde buscar los archivos de configuración

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error al leer archivo de configuración: %v", err)
	}
}

func GetLogLevel() string {
	return viper.GetString("log.level")
}

func GetLogFile() string {
	return viper.GetString("log.logfile")
}

func GetLogEncoding() string {
	return viper.GetString("log.encoding")
}
