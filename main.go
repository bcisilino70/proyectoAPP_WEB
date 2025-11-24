package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"proyectoAPP_WEB/persistencia/handlers"

	_ "github.com/lib/pq"
)

/*
# Variables de configuración
CONTAINER_NAME = app_postgres
DB_NAME = app_db
DB_USER = app_user
DB_PASSWORD = app_pass
DB_ADMIN = app_admin
ADMIN_PASSWORD = admin_pass
DB_PORT = 5432
*/

func main() {
	//Inicializar el servidor HTTP
	fs := http.FileServer(http.Dir("./persistencia/views"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Lógica para determinar el host de la DB
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost" // Fallback para ejecución local sin Docker
	}

	//Inicializar base de datos.
	connStr := fmt.Sprintf("host=%s port=5432 user=app_user password=app_pass dbname=app_db sslmode=disable", dbHost)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close()
	handlers.InitDB(db)
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/userpage", handlers.UserPageHandler)

	//Funcionalidades.
	http.HandleFunc("/crear-resena", handlers.CrearResenaHandler)
	http.HandleFunc("/resenas/{id}", handlers.EliminarResenaHandler)

	log.Println("Servidor iniciado en http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
