# Variables
pkg := "./..."
folders := "./internal"

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
	gofmt -e {{folders}}

# Ayuda sobre comandos disponibles
help:
	@echo "Comandos disponibles:"
	@echo "  make all     - Ejecuta install, test, build, clean, check y help"
	@echo "  make install - Instala dependencias."
	@echo "  make test    - Ejecuta pruebas unitarias."
	@echo "  make build   - Construye el proyecto."
	@echo "  make clean   - Limpia archivos generados."
	@echo "  make check   - Verifica sintaxis"
	@echo "  make help    - Muestra esta ayuda."
