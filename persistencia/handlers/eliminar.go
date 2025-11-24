package handlers

import (
	"log"
	"net/http"
	"strconv"

	// Ya no necesitamos importar "views" porque no vamos a renderizar nada
	db "proyectoAPP_WEB/persistencia/db/sqlc"
)

func EliminarResenaHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Validar Método (aunque el router de Go 1.22+ ya lo hace si pusiste "DELETE ...")
	// Lo dejamos por seguridad o costumbre.
	if r.Method != http.MethodDelete {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// 2. Obtener el ID de la resena desde la URL (Go 1.22 feature)
	idStr := r.PathValue("id")
	resenaID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de reseña inválido", http.StatusBadRequest)
		return
	}

	// 3. Obtener el ID del cliente desde la cookie (Autenticación)
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

	// 4. Preparar parámetros
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

	// 6. RESPUESTA HTMX "VACÍA"
	// Al enviar 200 OK sin cuerpo, HTMX toma ese "nada" y reemplaza
	// el target (la tarjeta de la reseña) por "nada", eliminándola visualmente.
	w.WriteHeader(http.StatusOK)
}
