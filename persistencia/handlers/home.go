package handlers

import (
	"net/http"

	"proyectoAPP_WEB/persistencia/views"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	component := views.HomePage()

	component.Render(r.Context(), w)
}
