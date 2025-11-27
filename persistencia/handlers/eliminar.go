package handlers

import (
	"log"
	"net/http"
	"strconv"

	db "proyectoAPP_WEB/persistencia/db/sqlc"
)

func EliminarResenaHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Validar metodo
	if r.Method != http.MethodDelete {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// 2. Obtener el ID de la resena desde la URL
	idStr := r.PathValue("id")
	resenaID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de reseña invalido", http.StatusBadRequest)
		return
	}

	// 3. Obtener el ID del cliente desde la cookie
	cookie, err := r.Cookie("uid")
	if err != nil {
		http.Error(w, "No autenticado", http.StatusUnauthorized)
		return
	}
	clienteID, err := strconv.Atoi(cookie.Value)
	if err != nil {
		http.Error(w, "ID de usuario invalido", http.StatusBadRequest)
		return
	}

	// 4. Preparar parametros
	params := db.DeleteResenaParams{
		ID:        int32(resenaID),
		ClienteID: int32(clienteID),
	}

	// 5. Ejecutar borrado en BD
	err = queries.DeleteResena(r.Context(), params)
	if err != nil {
		log.Printf("Error al eliminar resena: %v", err)
		http.Error(w, "No se pudo eliminar la resena", http.StatusInternalServerError)
		return
	}

	// 6. Respuesta htmx vacia
	w.WriteHeader(http.StatusOK)
}
