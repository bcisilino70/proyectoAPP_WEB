package handlers

import (
	"log"
	"net/http"
	"strconv"

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
	username := "Usuario" + cookie.Value

	// 4. Obtener la lista de "Mis Reseñas" usando sqlc
	misResenas, err := queries.ListResenas(r.Context(), clienteID)
	if err != nil {
		log.Printf("Error al obtener ListResenas: %v", err)
		http.Error(w, "Error al cargar tus resenas", http.StatusInternalServerError)
		return
	}

	// Obtener las resenas recientes de todos (ej. las ultimas 10)
	resenasRecientes, err := queries.ListResenasRecientes(r.Context(), 10)
	if err != nil {
		log.Printf("Error al obtener ListResenasRecientes: %v", err)
		http.Error(w, "Error al cargar reseñas recientes", http.StatusInternalServerError)
		return
	}

	// 5. Renderizar el componente principal de la pagina de usuario
	//    PASANDOLE LOS DATOS OBTENIDOS
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Ahora pasamos los slices de resenas al componente Templ
	componente := views.UserPage(username, misResenas, resenasRecientes)
	err = componente.Render(r.Context(), w)

	if err != nil {
		log.Printf("Error al renderizar UserPage: %v", err)
		http.Error(w, "Error al renderizar la pagina", http.StatusInternalServerError)
	}
}
