package handlers

import (
	"log"
	"net/http"
	"strconv"

	db "proyectoAPP_WEB/persistencia/db/sqlc"
)

// CrearResenaHandler maneja el POST del formulario de nueva reseña
func CrearResenaHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Solo aceptamos POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// 2. Obtener el ID del cliente desde la cookie de sesión
	cookie, err := r.Cookie("uid")
	if err != nil {
		http.Error(w, "No autenticado", http.StatusUnauthorized)
		return
	}
	// Convertir el ID de la cookie (string) a int32 para la BD
	clienteID, err := strconv.Atoi(cookie.Value)
	if err != nil {
		http.Error(w, "ID de usuario inválido", http.StatusBadRequest)
		return
	}

	// 3. Leer y parsear los datos del formulario
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Formulario inválido", http.StatusBadRequest)
		return
	}

	// 4. Convertir la nota (string) a int32
	nota, err := strconv.Atoi(r.FormValue("nota"))
	if err != nil {
		http.Error(w, "Nota inválida", http.StatusBadRequest)
		return
	}

	// 5. Preparar los parámetros para la consulta SQLC
	// (Usamos los 'name' de los inputs del formulario)
	params := db.CreateResenaParams{
		Titulo:      r.FormValue("titulo"),      //
		Descripcion: r.FormValue("descripcion"), //
		Nota:        int32(nota),                //
		ClienteID:   int32(clienteID),           // ID de la cookie
	}

	// 6. Ejecutar la consulta SQLC para crear la reseña
	//    (queries es la variable global definida en register.go)
	_, err = queries.CreateResena(r.Context(), params) //
	if err != nil {
		log.Printf("Error al crear reseña: %v", err)
		http.Error(w, "No se pudo crear la reseña (título duplicado u otro error)", http.StatusBadRequest)
		return
	}

	// 7. Redirigir al usuario a su panel si todo salio bien
	http.Redirect(w, r, "/userpage", http.StatusSeeOther)
}
