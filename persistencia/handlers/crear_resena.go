package handlers

import (
	"log"
	"net/http"
	"strconv"

	db "proyectoAPP_WEB/persistencia/db/sqlc"
	"proyectoAPP_WEB/persistencia/views"
)

// CrearResenaHandler maneja el POST del formulario de nueva resena
func CrearResenaHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Solo aceptamos POST
	if r.Method != http.MethodPost {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}

	// 2. Obtener el ID del cliente desde la cookie de sesion
	cookie, err := r.Cookie("uid")
	if err != nil {
		http.Error(w, "No autenticado", http.StatusUnauthorized)
		return
	}
	// Convertir el ID de la cookie (string) a int32 para la BD
	clienteID, err := strconv.Atoi(cookie.Value)
	if err != nil {
		http.Error(w, "ID de usuario invalido", http.StatusBadRequest)
		return
	}

	// 3. Leer y parsear los datos del formulario
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Formulario invalido", http.StatusBadRequest)
		return
	}

	// 4. Convertir la nota (string) a int32
	nota, err := strconv.Atoi(r.FormValue("nota"))
	if err != nil {
		http.Error(w, "Nota invalida", http.StatusBadRequest)
		return
	}

	// 5. Preparar los parametros para la consulta SQLC
	params := db.CreateResenaParams{
		Titulo:      r.FormValue("titulo"),
		Descripcion: r.FormValue("descripcion"),
		Nota:        int32(nota),
		ClienteID:   int32(clienteID), // ID de la cookie
	}

	// 6. Ejecutar la consulta SQLC para crear la resena
	_, err = queries.CreateResena(r.Context(), params)
	if err != nil {
		log.Printf("Error al crear resena: %v", err)
		http.Error(w, "No se pudo crear la resena ", http.StatusBadRequest)
		return
	}

	// ---HTMX ---

	// 7. En lugar de redirigir, obtenemos la lista actualizada de este cliente
	misResenas, err := queries.ListResenas(r.Context(), int32(clienteID))
	if err != nil {
		http.Error(w, "Error al actualizar la lista", http.StatusInternalServerError)
		return
	}

	// 8. Renderizamos el componente MisResenas
	// HTMX toma este HTML y lo coloca en el hx-target (#mis-resenas)
	// reemplazando el contenido anterior podriamos hacer lo mismo para las resenas recientes
	// pero no se cumpliria ni veria la idea de usar htmx para solo renderizar una parte
	componente := views.MisResenas(misResenas)
	componente.Render(r.Context(), w)
}
