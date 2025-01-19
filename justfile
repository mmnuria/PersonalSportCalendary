pkg := "./..."
folders := "./internal"

install:
	@echo "Instalando dependencias..."
	go mod tidy

check:
	@echo "Verificando sintaxis..."
	gofmt -e {{folders}}

help:
	@echo "Comandos disponibles:"
	@echo "  just install - Instala dependencias."
	@echo "  just check   - Verifica sintaxis"
