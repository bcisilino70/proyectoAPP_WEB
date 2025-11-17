package handlers

import (
	"log" // <-- AÑADIDO
	"net/http"
	"strconv" // <-- AÑADIDO

	// db "proyectoAPP_WEB/persistencia/db/sqlc" // <-- AÑADIDO (Ver nota abajo)
	"proyectoAPP_WEB/persistencia/views"
)

func UserPageHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Servimos la user_page.templ
	cookie, err := r.Cookie("uid")
	if err != nil {
		http.Error(w, "No autenticado", http.StatusUnauthorized)
		return
	}

	// 2. Convertir el ID de la cookie a int32 para usarlo en la BD
	clienteID_64, err := strconv.ParseInt(cookie.Value, 10, 32)
	if err != nil {
		http.Error(w, "ID de usuario inválido", http.StatusBadRequest)
		return
	}
	clienteID := int32(clienteID_64)

	// 3. Crear un "username" simple
	username := "Usuario" + cookie.Value // Opcional: podrías hacer un GetCliente para buscar su nombre real

	// 4. === INICIO DE LA NUEVA LÓGICA (CUMPLE EL PUNTO 2) ===

	// 4a. Obtener la lista de "Mis Reseñas" usando sqlc
	//     (queries es la variable global de tu paquete 'handlers')
	misResenas, err := queries.ListResenas(r.Context(), clienteID)
	if err != nil {
		log.Printf("Error al obtener ListResenas: %v", err)
		http.Error(w, "Error al cargar tus reseñas", http.StatusInternalServerError)
		return
	}

	// 4b. Obtener las reseñas recientes de todos (ej. las últimas 10)
	//     (10 es el parámetro 'LIMIT $1' de tu consulta sqlc)
	resenasRecientes, err := queries.ListResenasRecientes(r.Context(), 10)
	if err != nil {
		log.Printf("Error al obtener ListResenasRecientes: %v", err)
		http.Error(w, "Error al cargar reseñas recientes", http.StatusInternalServerError)
		return
	}
	// === FIN DE LA NUEVA LÓGICA ===

	// 5. Renderizar el componente principal de la página de usuario
	//    PASÁNDOLE LOS DATOS OBTENIDOS
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// --- MODIFICADO ---
	// Ahora pasamos los slices de reseñas al componente Templ
	componente := views.UserPage(username, misResenas, resenasRecientes)
	err = componente.Render(r.Context(), w)

	if err != nil {
		log.Printf("Error al renderizar UserPage: %v", err)
		http.Error(w, "Error al renderizar la página", http.StatusInternalServerError)
	}
}
