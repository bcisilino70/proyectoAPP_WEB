package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"proyectoAPP_WEB/logica_neg/pkg/models"
	sqlc "proyectoAPP_WEB/persistencia/db/sqlc"
)

// --- GET CLIENTES es una prueba --- //
func ClienteHandlerPrueba(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Clientes) //Dato: cuando importas variables de otros paquetes como Clientes debe estar en mayuscula.
}

// --- GET RESENAS es una prueba --- //
func ResenasHandlerPrueba(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Resenas)
}

// --- GET CLIENTES usando BD con el metodo LIST de sqlc --- //
/*
	- funcion que tiene como parametro una instancia de la base de datos hecha en el main.go y devuelve un handlerfunc
	- Notar que usa el queries del main para hacer la consulta a la base de datos.
	- guarda en Clientes el resultado del ListCliente que genera sqlc generate
	- y luego lo codifica a json con los datos que nosostros queremos mostrar.
	Responde a este enunciado: 	GET /<entidades>: Debe usar el método List... de sqlc para obtener todos los registros y devolverlos
								como un array JSON. Muestra todos los clientes.
*/
func ClientesHandler(queries *sqlc.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		Clientes, err := queries.ListCliente(r.Context())
		if err != nil {
			http.Error(w, "Error al obtener los clientes", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(Clientes)
	}
}

// --- GET RESENAS usando BD con el metodo LIST de sqlc --- //
/*
	- Revisando queries.sql.go el metodo ListResenas tiene dos parametros context y clienteID
	- Por lo que hay que obtener de algun lado el clienteID para invocar el metodo ListResenas.
	- Se obtiene de la URL entendiendo esta estructura de URL /resenas?cliente_id=1
	Responde a este enunciado:	GET /<entidades>/{id}: Debe obtener el ID de la URL, buscar el registro con Get... y devolverlo. Si no
								existe, debe devolver un 404. Muestra todas las resenas de un usuario.
*/
func ResenasHandler(queries *sqlc.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		//Obtener el clienteID de la URL
		clienteIDStr := r.URL.Query().Get("cliente_id")
		if clienteIDStr == "" {
			http.Error(w, "cliente_id es requerido", http.StatusBadRequest)
			return
		}
		//Convertir el clienteID a int32
		var clienteID int32
		_, err := fmt.Sscanf(clienteIDStr, "%d", &clienteID)
		if err != nil {
			http.Error(w, "cliente_id invalido", http.StatusBadRequest)
			return
		}

		Resenas, err := queries.ListResenas(r.Context(), clienteID)
		if err != nil {
			http.Error(w, "Error al obtener las reseñas", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(Resenas)
	}

}
