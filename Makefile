# Makefile para Persistencia de Datos
# Automatización completa del flujo

# Variables de configuración
CONTAINER_NAME = app_postgres
DB_NAME = app_db
DB_USER = app_user
DB_PASSWORD = app_pass
DB_ADMIN = app_admin
ADMIN_PASSWORD = admin_pass
DB_PORT = 5432

.PHONY: help all setup run test clean destroy

# Ayuda - muestra todos los comandos
help:
	@echo "🚀 COMANDOS DISPONIBLES:"
	@echo "  make all      - Flujo completo: setup + run + test + destroy"
	@echo "  make setup    - Crear contenedor y configurar BD"
	@echo "  make run      - Ejecutar la aplicación"
	@echo "  make test     - Ejecutar tests"
	@echo "  make clean    - Limpiar archivos temporales"
	@echo "  make destroy  - Eliminar contenedor y recursos"
	@echo "  make status   - Ver estado del contenedor"

# Flujo completo automático
all:
	@echo "🎯 INICIANDO FLUJO COMPLETO AUTOMÁTICO..."
	@echo "=========================================="
	@echo ""
	$(MAKE) destroy
	@echo ""
	$(MAKE) setup
	@echo ""
	$(MAKE) run_logica
	@echo "✅ FLUJO COMPLETADO EXITOSAMENTE!"

# Configuración inicial
setup:
	@echo "🐳 CONFIGURANDO ENTORNO..."
	
	@# 1. Crear contenedor PostgreSQL
	@echo "1. Creando contenedor PostgreSQL..."
	docker run --name $(CONTAINER_NAME) \
		-e POSTGRES_DB=$(DB_NAME) \
		-e POSTGRES_USER=$(DB_USER) \
		-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
		-p $(DB_PORT):5432 \
		-d postgres:13
	
	@# 2. Esperar que PostgreSQL esté listo
	@echo "2. Esperando que PostgreSQL esté listo..."
	@sleep 10
	
	@# 3. Crear usuario admin y otorgar privilegios
	@echo "3. Configurando usuario administrador..."
	echo "CREATE USER $(DB_ADMIN) WITH PASSWORD '$(ADMIN_PASSWORD)';" | \
	docker exec -i $(CONTAINER_NAME) psql -U $(DB_USER) -d $(DB_NAME)
	echo "GRANT ALL PRIVILEGES ON DATABASE $(DB_NAME) TO $(DB_ADMIN);" | \
	docker exec -i $(CONTAINER_NAME) psql -U $(DB_USER) -d $(DB_NAME)
	
	@# 4. Ejecutar schema.sql
	@echo "4. Creando tablas con schema.sql..."
	docker exec -i $(CONTAINER_NAME) psql -U $(DB_USER) -d $(DB_NAME) < persistencia/db/schema/schema.sql
	
	@echo "✅ CONFIGURACIÓN COMPLETADA!"
	@echo "   Contenedor: $(CONTAINER_NAME)"
	@echo "   Base de datos: $(DB_NAME)"
	@echo "   Puerto: $(DB_PORT)"

run_logica: 
	@echo "🚀 EJECUTANDO APLICACIÓN NIVEL CAPA DE LOGICA DE NEGOCIOS..."
	@echo "=========================================="
	go run ./cmd/server/main.go
	@echo "✅ EJECUCIÓN COMPLETADA!"

# Ejecutar tests
test: setup
	@echo "🧪 EJECUTANDO TESTS..."
	@echo "=========================================="
	go test -v ./...
	@echo "✅ TESTS COMPLETADOS!"

# Ver estado del contenedor
status:
	@echo "📊 ESTADO DEL CONTENEDOR:"
	@docker ps -a --filter "name=$(CONTAINER_NAME)" --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"

# Limpiar archivos temporales
clean:
	@echo "🧹 LIMPIANDO ARCHIVOS TEMPORALES..."
	go clean
	rm -f *.log
	@echo "✅ LIMPIEZA COMPLETADA!"

# Destruir contenedor y recursos
destroy:
	@echo "🗑️  ELIMINANDO CONTENEDOR Y RECURSOS..."
	@-docker stop $(CONTAINER_NAME) 2>/dev/null || true
	@-docker rm $(CONTAINER_NAME) 2>/dev/null || true
	@echo "✅ CONTENEDOR ELIMINADO: $(CONTAINER_NAME)"

# Comando para desarrollo (solo setup, sin destroy)
dev: setup
	@echo "🔧 MODO DESARROLLO:"
	@echo "   Contenedor activo: $(CONTAINER_NAME)"
	@echo "   Puerto: $(DB_PORT)"
	@echo "   Ejecuta 'make destroy' cuando termines"

# Conectar a la base de datos
db-connect:
	@echo "🔗 CONECTANDO A LA BASE DE DATOS..."
	docker exec -it $(CONTAINER_NAME) psql -U $(DB_USER) -d $(DB_NAME)

# Ver datos de ejemplo
db-see-data:
	@echo "👀 VISUALIZANDO DATOS..."
	@echo "--- CLIENTES ---"
	docker exec $(CONTAINER_NAME) psql -U $(DB_USER) -d $(DB_NAME) -c "SELECT * FROM clientes;"
	@echo ""
	@echo "--- RESEÑAS ---"
	docker exec $(CONTAINER_NAME) psql -U $(DB_USER) -d $(DB_NAME) -c "SELECT * FROM resenas;"