package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq" // Driver de PostgreSQL

	"proyectoAPP_WEB/logica_neg/pkg/handlers"

	sqlc "proyectoAPP_WEB/persistencia/db/sqlc"
)

func main() {
	// Reemplaza esta cadena con los datos de tu conexión a PostgreSQL
	connStr := "host=localhost port=5432 user=app_user password=app_pass dbname=app_db sslmode=disable"

	// 1. Conexión con la Base de Datos
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}
	defer db.Close()

	// Ping a la base de datos para ver si funciona
	if err := db.Ping(); err != nil {
		log.Fatalf("No se pudo hacer ping a la base de datos: %v", err)
	}

	log.Println("Conexión a la base de datos establecida exitosamente.")

	// Crea una instancia del repositorio generado por sqlc
	queries := sqlc.New(db)

	// Algo para que queries no tire error de variable no usada
	_ = queries

	// ----- CONFIGURACION DE RUTAS (ESTILO REST) ----- //

	// Ruta raíz para verificar que el servidor está funcionando
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"API de Clientes y Reseñas funcionando","version":"1.0","endpoints":["/clientes","/resenas"]}`))
	})

	http.HandleFunc("/clientes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet: // Listar
			handlers.ClientesHandler(queries)(w, r)
		case http.MethodPost: // Crear
			handlers.CrearClienteHandler(queries)(w, r)
		case http.MethodPut: // Actualizar
			handlers.UpdateClienteHandler(queries)(w, r)
		case http.MethodDelete: // Eliminar
			handlers.DeleteClienteHandler(queries)(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/resenas", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet: // Listar
			handlers.ResenasHandler(queries)(w, r)
		case http.MethodPost: // Crear
			handlers.CrearResenaHandler(queries)(w, r)
		case http.MethodPut: // Actualizar
			handlers.UpdateResenaHandler(queries)(w, r)
		case http.MethodDelete: // Eliminar
			handlers.DeleteResenaHandler(queries)(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	log.Println("Iniciando servidor en http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("No se pudo iniciar el servidor: %v", err)
	}
}
