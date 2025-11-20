package handlers

import (
	"log"
	"net/http"
	"strconv"

	db "proyectoAPP_WEB/persistencia/db/sqlc"
)

// EliminarResenaHandler maneja el POST para borrar una resena
func EliminarResenaHandler(w http.ResponseWriter, r *http.Request) {
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

	// 4. Obtener el ID de la resena desde el input oculto
	resenaID, err := strconv.Atoi(r.FormValue("resena_id"))
	if err != nil {
		http.Error(w, "ID de resena invalido", http.StatusBadRequest)
		return
	}

	// 5. Preparar los parámetros para la consulta SQLC
	//    (La consulta 'DeleteResena' espera el ID de la resena y el cliente_id
	//    para asegurarse de que solo el dueno pueda borrarla)
	params := db.DeleteResenaParams{
		ID:        int32(resenaID),
		ClienteID: int32(clienteID),
	}

	// 6. Ejecutar la consulta SQLC para borrar la resena
	err = queries.DeleteResena(r.Context(), params)
	if err != nil {
		log.Printf("Error al eliminar resena: %v", err)
		http.Error(w, "No se pudo eliminar la resena", http.StatusInternalServerError)
		return
	}

	// 7. Si todo sale bien, redirigir al usuario a su panel
	http.Redirect(w, r, "/userpage", http.StatusSeeOther)
}
