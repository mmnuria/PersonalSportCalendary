# Variables
pkg := "./..."
folders := "./internal"

all: install test build clean check help
	@echo "Ejecutando todas las tareas principales..."

install:
	@echo "Instalando dependencias..."
	go mod tidy

test:
	@echo "Ejecutando pruebas unitarias..."
	go test {{pkg}} -v

build:
	@echo "Construyendo el proyecto..."
	go build -o bin/app .

clean:
	@echo "Limpiando archivos generados..."
	rm -rf bin/

check:
	@echo "Verificando sintaxis..."
	gofmt -e {{folders}}

help:
	@echo "Comandos disponibles:"
	@echo "  make all     - Ejecuta install, test, build, clean, check y help"
	@echo "  make install - Instala dependencias."
	@echo "  make test    - Ejecuta pruebas unitarias."
	@echo "  make build   - Construye el proyecto."
	@echo "  make clean   - Limpia archivos generados."
	@echo "  make check   - Verifica sintaxis"
	@echo "  make help    - Muestra esta ayuda."
