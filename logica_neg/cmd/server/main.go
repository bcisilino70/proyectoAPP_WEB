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

	// ----- CONFIGURACION DE RUTAS ----- //
	http.HandleFunc("/clientes", handlers.ClientesHandler(queries))
	http.HandleFunc("/resenas", handlers.ResenasHandler(queries))
	// ---------------------------------- //
	/*
		ESTAN HECHOS EN HANDLER.GO DOS PUNTOS DEL ENUNCIADO, LOS GET PARA LISTAR TODOS LOS CLIENTES Y LAS RESENAS DE UN CLIENTE
		1. MAKE ALL -> MAKE DESTROY - MAKE RUN_DATOS ( capa de datos ) - MAKE RUN_SERVER ( capa de logica de negocio )
		2. ABRIR OTRA TERMINAL Y EJECUTAR hurl logica_neg/hurl/clientes.hurl o hurl logica_neg/hurl/resenas.hurl
		Archivos para mirar:
		- logica_neg/pkg/handlers/handler.go
		- logica_neg/hurl/clientes.hurl y logica_neg/hurl/resenas.hurl
		- logica_neg/pkg/handlers/models.go ( tiene los modelos que muestra la logica de negocio evitando mostrar datos sensibles como contraseñas o IDs)
	*/
	log.Println("Iniciando servidor en http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("No se pudo iniciar el servidor: %v", err)
	}
}
