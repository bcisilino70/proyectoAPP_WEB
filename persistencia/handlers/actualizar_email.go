package handlers

import (
	"log"
	"net/http"
	"strconv"

	db "proyectoAPP_WEB/persistencia/db/sqlc"
)

// ActualizarEmailHandler maneja el POST para cambiar el email del usuario
func ActualizarEmailHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Solo aceptamos POST
	if r.Method != http.MethodPost {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}

	// 2. Obtener el ID del cliente desde la cookie de sesion ("uid")
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

	// 4. Obtener el nuevo email del formulario (el 'name' del input)
	nuevoEmail := r.FormValue("email")
	if nuevoEmail == "" {
		http.Error(w, "El email no puede estar vacio", http.StatusBadRequest)
		return
	}

	// 5. Preparar los parametros para la consulta SQLC
	//    (La consulta 'UpdateCliente' espera Email y ID)
	params := db.UpdateClienteParams{
		Email: nuevoEmail,
		ID:    int32(clienteID),
	}

	// 6. Ejecutar la consulta SQLC para actualizar el cliente
	err = queries.UpdateCliente(r.Context(), params)
	if err != nil {
		log.Printf("Error al actualizar email: %v", err)
		// El error más común aquí sería un email duplicado (UNIQUE constraint)
		http.Error(w, "No se pudo actualizar el email", http.StatusInternalServerError)
		return
	}

	// 7. Si todo sale bien, redirigir al usuario a su panel
	http.Redirect(w, r, "/userpage", http.StatusSeeOther)
}
