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
	@echo "  make install - Instala dependencias."
	@echo "  make check   - Verifica sintaxis"