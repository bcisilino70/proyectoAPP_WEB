package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "proyectoAPP_WEB/logica_neg/pkg/models"
	sqlc "proyectoAPP_WEB/persistencia/db/sqlc"
)

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

// --- CREAR CLIENTE usando BD con el metodo CREATE de sqlc --- //
/*
	- Recibe datos en formato JSON desde el cuerpo de la petición
	- Valida que los campos requeridos no estén vacíos
	- Usa el método CreateCliente de sqlc para guardar en la base de datos
	- Devuelve el nuevo cliente creado como JSON con estado 201
*/
func CrearClienteHandler(queries *sqlc.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// 1. Decodificar JSON
		var params sqlc.CreateClienteParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, `{"error":"JSON inválido"}`, http.StatusBadRequest)
			return
		}

		// 2. Validaciones
		if params.Nombre == "" || params.Apellido == "" ||
			params.Usuario == "" || params.Pass == "" || params.Email == "" {
			http.Error(w, `{"error":"Todos los campos son requeridos"}`, http.StatusBadRequest)
			return
		}

		// 3. Crear en la BD
		cliente, err := queries.CreateCliente(r.Context(), params)
		if err != nil {
			http.Error(w, `{"error":"Error al crear cliente"}`, http.StatusInternalServerError)
			return
		}

		// 4. Respuesta 201 Created con el cliente creado
		w.WriteHeader(http.StatusCreated) // 201
		json.NewEncoder(w).Encode(cliente)
	}
}

// --- UPDATE CLIENTE usando BD con el metodo UPDATE de sqlc --- //
/*
	- Recibe datos en formato JSON con el nuevo email, usuario y contraseña
	- Valida que los campos requeridos no estén vacíos
	- Usa el método UpdateCliente de sqlc para actualizar en la base de datos
	- Devuelve estado 200 OK si se actualizó correctamente
*/
func UpdateClienteHandler(queries *sqlc.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// 1. Decodificar JSON
		var params sqlc.UpdateClienteParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, `{"error":"JSON inválido"}`, http.StatusBadRequest)
			return
		}

		// 2. Validaciones
		if params.Email == "" || params.Usuario == "" || params.Pass == "" {
			http.Error(w, `{"error":"Email, usuario y contraseña son requeridos"}`, http.StatusBadRequest)
			return
		}

		// 3. Actualizar en la BD
		err := queries.UpdateCliente(r.Context(), params)
		if err != nil {
			http.Error(w, `{"error":"Error al actualizar cliente"}`, http.StatusInternalServerError)
			return
		}

		// 4. Respuesta 200 OK con mensaje
		w.WriteHeader(http.StatusOK) // 200
		w.Write([]byte(`{"message":"Cliente actualizado correctamente"}`))
	}
}

// --- DELETE CLIENTE usando BD con el metodo DELETE de sqlc --- //
/*
	- Obtiene el usuario del query parameter o del body JSON
	- Usa el método DeleteCliente de sqlc para eliminar de la base de datos
	- Devuelve estado 204 No Content si se eliminó correctamente
*/
func DeleteClienteHandler(queries *sqlc.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. Obtener usuario del query param

		// RE-HACER ESTA PARTE. EL DELETE NECESITA EL ID COMO PARAMETRO
		/*
			usuario := r.URL.Query().Get("usuario")
			if usuario == "" {
				http.Error(w, `{"error":"Se requiere el parámetro usuario"}`, http.StatusBadRequest)
				return
			}
		*/

		// 2. Eliminar de la BD
		err := queries.DeleteCliente(r.Context(), usuario)
		if err != nil {
			http.Error(w, `{"error":"Error al eliminar cliente"}`, http.StatusInternalServerError)
			return
		}

		// 3. Respuesta 204 No Content (sin body)
		w.WriteHeader(http.StatusNoContent) // 204
	}
}

// ---------------------------------------- RESENAS -------------------------------------------------------- //

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

// --- CREAR RESENA usando BD con el metodo CREATE de sqlc --- //
/*
	- Recibe datos en formato JSON desde el cuerpo de la petición
	- Valida que los campos requeridos no estén vacíos (título, descripción, nota válida)
	- Usa el método CreateResena de sqlc para guardar en la base de datos
	- Devuelve la nueva reseña creada como JSON con estado 201
*/

func CrearResenaHandler(queries *sqlc.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementación para crear una resena
	}
}

// --- UPDATE RESENA usando BD con el metodo UPDATE de sqlc --- //
func UpdateResenaHandler(queries *sqlc.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementación para actualizar una reseña
	}
}

// --- DELETE RESENA usando BD con el metodo DELETE de sqlc --- //
func DeleteResenaHandler(queries *sqlc.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementación para eliminar una reseña
	}
}
