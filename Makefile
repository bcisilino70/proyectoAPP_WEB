.PHONY: all up down logs clean db-data

all:
	@make down
	@make up
	@echo " Esperando 5 segundos a que la base de datos inicie."
	@sleep 5
	@make test
	@make db-data
	@echo ""
	@echo "========================================"
	@echo "Servidor disponible en: http://localhost:8080"
	@echo "========================================"
# Levanta los servicios con Docker Compose
up:
	@echo "🚀 Levantando entorno con Docker Compose..."
	docker compose up --build -d
	@echo "✅ Servidor corriendo en http://localhost:8080"

# Detiene y elimina los contenedores
down:
	@echo "🛑 Deteniendo servicios..."
	docker compose down

# Ver datos de las tablas
db-data:
	@echo "👀 VISUALIZANDO DATOS..."
	@echo "--- CLIENTES ---"
	docker exec app_postgres_db psql -U app_user -d app_db -c "SELECT * FROM cliente;"
	@echo ""
	@echo "--- RESEÑAS ---"
	docker exec app_postgres_db psql -U app_user -d app_db -c "SELECT * FROM resena;"

test:
	@echo "🧪 Ejecutando tests de integración..."
	hurl -v tests/cliente.hurl
	hurl -v tests/resena.hurl
