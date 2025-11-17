package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"proyectoAPP_WEB/persistencia/handlers"
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
	//Inicializar base de datos.
	connStr := "host=localhost port=5432 user=app_user password=app_pass dbname=app_db sslmode=disable"
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
	http.HandleFunc("/eliminar-resena", handlers.EliminarResenaHandler)
	http.HandleFunc("/actualizar-email", handlers.ActualizarEmailHandler)

	log.Println("Servidor iniciado en http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
