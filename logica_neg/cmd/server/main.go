package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq" // Driver de PostgreSQL
	// Asegúrate de que la ruta de importación coincida con el nombre de tu módulo en go.mod
	// seguido de la ruta al paquete sqlc.
	// Por ejemplo: "proyectoAPP_WEB/persistencia/db/sqlc"
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

	// 2. Configuración del Servidor HTTP
	// Aquí es donde configurarías tus handlers y el router HTTP,
	// pasando la instancia 'queries' a los handlers que la necesiten.
	// Por ejemplo:
	//
	// router := http.NewServeMux()
	// api := &API{db: queries} // Suponiendo que tienes una struct para tus handlers
	// router.HandleFunc("/entidad", api.ListEntidadesHandler)

	log.Println("Iniciando servidor en http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("No se pudo iniciar el servidor: %v", err)
	}
}
