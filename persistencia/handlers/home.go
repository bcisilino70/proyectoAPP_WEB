package handlers

import (
	"net/http"

	"proyectoAPP_WEB/persistencia/views"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	componente := views.HomePage()

	componente.Render(r.Context(), w)
}
