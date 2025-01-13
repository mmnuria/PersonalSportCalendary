pkg := ./...
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
	@echo "  just all     - Ejecuta install, test, build, clean, check y help"
	@echo "  just install - Instala dependencias."
	@echo "  just test    - Ejecuta pruebas unitarias."
	@echo "  just build   - Construye el proyecto."
	@echo "  just clean   - Limpia archivos generados."
	@echo "  just check   - Verifica sintaxis"
	@echo "  just help    - Muestra esta ayuda."
