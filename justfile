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

# Verificar sintaxis y formateo del código
check:
	@echo "Verificando sintaxis y formato..."
	go fmt {{pkg}}
	go vet {{pkg}}

# Ayuda sobre comandos disponibles
help:
	@echo "Comandos disponibles:"
	@echo "  just all     - Ejecuta tidy, fmt, vet, lint, test y build."
	@echo "  just install - Instala dependencias."
	@echo "  just test    - Ejecuta pruebas unitarias."
	@echo "  just build   - Construye el proyecto."
	@echo "  just clean   - Limpia archivos generados."
	@echo "  just check   - Verifica sintaxis y formateo del código."
	@echo "  just help    - Muestra esta ayuda."
