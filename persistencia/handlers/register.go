package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	db "proyectoAPP_WEB/persistencia/db/sqlc"
)

// queries es el acceso a la capa generada por sqlc.
var queries *db.Queries

// InitDB inicializa queries a partir de la conexión SQL.
func InitDB(sqlDB *sql.DB) {
	queries = db.New(sqlDB)
}

// RegisterHandler procesa POST /register
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Formulario invalido", http.StatusBadRequest)
		return
	}
	p := db.CreateClienteParams{
		Nombre:   r.FormValue("nombre"),
		Apellido: r.FormValue("apellido"),
		Usuario:  r.FormValue("usuario"),
		Pass:     r.FormValue("contrasena"),
		Email:    r.FormValue("mail"),
	}
	cli, err := queries.CreateCliente(r.Context(), p)
	if err != nil {
		log.Printf("error CreateCliente: %v", err)
		http.Error(w, "No se pudo registrar", http.StatusBadRequest)
		return
	}

	// Guardar ID en cookie de sesión simple
	http.SetCookie(w, &http.Cookie{
		Name:     "uid",
		Value:    fmt.Sprint(cli.ID),
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	http.Redirect(w, r, "/userpage", http.StatusSeeOther)
}

// LoginHandler procesa POST /login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Formulario invalido", http.StatusBadRequest)
		return
	}
	p := db.GetClienteUsuarioYPassParams{
		Usuario: r.FormValue("usuario"),
		Pass:    r.FormValue("contrasena"),
	}
	cli, err := queries.GetClienteUsuarioYPass(r.Context(), p)
	if err != nil {
		log.Printf("login invalido: %v", err)
		http.Error(w, "Usuario o contrasena incorrectos", http.StatusUnauthorized)
		return
	}

	// Guardar ID en cookie de seion simple
	http.SetCookie(w, &http.Cookie{
		Name:     "uid",
		Value:    fmt.Sprint(cli.ID),
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	// Redirigir a la pagina principal o de usuario
	http.Redirect(w, r, "/userpage", http.StatusSeeOther)
}
