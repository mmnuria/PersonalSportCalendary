# Variables
pkg := ./...

# Tareas

# Tarea principal
all: install test build clean check help
	@echo "Ejecutando todas las tareas principales..."

# Instalar dependencias necesarias
install:
	@echo "Instalando dependencias..."
	go mod tidy

# Ejecutar pruebas unitarias
test:
	@echo "Ejecutando pruebas unitarias..."
	go test {{pkg}} -v

# Construir el binario
build:
	@echo "Construyendo el proyecto..."
	go build -o bin/app .

# Limpiar archivos generados
clean:
	@echo "Limpiando archivos generados..."
	rm -rf bin/

# Verificar sintaxis
check:
	@echo "Verificando sintaxis..."
	go fmt {{pkg}}

# Ayuda sobre comandos disponibles
help:
	@echo "Comandos disponibles:"
	@echo "  just all     - Ejecuta install, test, build, clean, check y help"
	@echo "  just install - Instala dependencias."
	@echo "  just test    - Ejecuta pruebas unitarias."
	@echo "  just build   - Construye el proyecto."
	@echo "  just clean   - Limpia archivos generados."
	@echo "  just check   - Verifica sintaxis"
	@echo "  just help    - Muestra esta ayuda."
