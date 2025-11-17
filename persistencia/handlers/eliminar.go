package handlers

import (
	"log"
	"net/http"
	"strconv"

	db "proyectoAPP_WEB/persistencia/db/sqlc"
)

// EliminarResenaHandler maneja el POST para borrar una reseña
func EliminarResenaHandler(w http.ResponseWriter, r *http.Request) {
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

	// 4. Obtener el ID de la reseña desde el input oculto
	resenaID, err := strconv.Atoi(r.FormValue("resena_id"))
	if err != nil {
		http.Error(w, "ID de reseña inválido", http.StatusBadRequest)
		return
	}

	// 5. Preparar los parámetros para la consulta SQLC
	//    (La consulta 'DeleteResena' espera el ID de la reseña y el cliente_id
	//    para asegurarse de que solo el dueño pueda borrarla)
	params := db.DeleteResenaParams{
		ID:        int32(resenaID),
		ClienteID: int32(clienteID),
	}

	// 6. Ejecutar la consulta SQLC para borrar la reseña
	err = queries.DeleteResena(r.Context(), params)
	if err != nil {
		log.Printf("Error al eliminar reseña: %v", err)
		http.Error(w, "No se pudo eliminar la reseña", http.StatusInternalServerError)
		return
	}

	// 7. Si todo sale bien, redirigir al usuario a su panel
	http.Redirect(w, r, "/userpage", http.StatusSeeOther)
}
