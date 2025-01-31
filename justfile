pkg := "./..."
folders := "./internal"

install:
	@echo "Instalando dependencias..."
	go mod tidy

check:
	@echo "Verificando sintaxis..."
	gofmt -e {{folders}}

test:
	@echo "Ejecutando pruebas unitarias..."
	go test {{pkg}} -v

help:
	@echo "Comandos disponibles:"
	@echo "  just install - Instala dependencias."
	@echo "  just check   - Verifica sintaxis"
	@echo "  just test   - Ejecuta pruebas unitarias"
