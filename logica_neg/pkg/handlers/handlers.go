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
		if params.Email == "" {
			http.Error(w, `{"error":"Email requerido"}`, http.StatusBadRequest)
			return
		}

		// 1. Obtener el ID del cliente desde la URL
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, `{"error":"Se requiere el parámetro id"}`, http.StatusBadRequest)
			return
		}

		// 2. Convertir el ID a int32
		_, err := fmt.Sscanf(idStr, "%d", &params.ID)
		if err != nil {
			http.Error(w, `{"error":"ID inválido"}`, http.StatusBadRequest)
			return
		}

		// 3. Actualizar en la BD
		err = queries.UpdateCliente(r.Context(), params)
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
	- Obtiene el ID del cliente desde la URL (query parameter)
	- Una vez que el usuario se loguea, se coloca el ID en la URL (problema para después)
	- Usa el método DeleteCliente de sqlc para eliminar de la base de datos
	- Devuelve estado 204 No Content si se eliminó correctamente
*/
func DeleteClienteHandler(queries *sqlc.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// 1. Obtener el ID del cliente desde la URL
		idStr := r.URL.Query().Get("id")
		if idStr == "" {
			http.Error(w, `{"error":"Se requiere el parámetro id"}`, http.StatusBadRequest)
			return
		}

		// 2. Convertir el ID a int32
		var clienteID int32
		_, err := fmt.Sscanf(idStr, "%d", &clienteID)
		if err != nil {
			http.Error(w, `{"error":"ID inválido"}`, http.StatusBadRequest)
			return
		}

		// 3. Eliminar de la BD
		err = queries.DeleteCliente(r.Context(), clienteID)
		if err != nil {
			http.Error(w, `{"error":"Error al eliminar cliente"}`, http.StatusInternalServerError)
			return
		}

		// 4. Respuesta 204 No Content (sin body)
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
		w.Header().Set("Content-Type", "application/json")

		// 1. Decodificar JSON (solo titulo, descripcion, nota)
		var input struct {
			Titulo      string `json:"titulo"`
			Descripcion string `json:"descripcion"`
			Nota        int32  `json:"nota"`
		}

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, `{"error":"JSON inválido"}`, http.StatusBadRequest)
			return
		}

		// 2. Validación en un solo if
		if input.Titulo == "" || input.Descripcion == "" || input.Nota < 1 || input.Nota > 10 {
			http.Error(w, `{"error":"Título, descripción son requeridos y la nota debe estar entre 1 y 10"}`, http.StatusBadRequest)
			return
		}

		// 3. Obtener el clienteID de la URL
		clienteIDStr := r.URL.Query().Get("cliente_id")
		if clienteIDStr == "" {
			http.Error(w, `{"error":"cliente_id es requerido"}`, http.StatusBadRequest)
			return
		}

		var clienteID int32
		_, err := fmt.Sscanf(clienteIDStr, "%d", &clienteID)
		if err != nil || clienteID <= 0 {
			http.Error(w, `{"error":"cliente_id inválido"}`, http.StatusBadRequest)
			return
		}

		// 4. Crear la reseña en la BD (combinando JSON + URL)
		resena, err := queries.CreateResena(r.Context(), sqlc.CreateResenaParams{
			Titulo:      input.Titulo,
			Descripcion: input.Descripcion,
			Nota:        input.Nota,
			ClienteID:   clienteID, // De la URL
		})
		if err != nil {
			http.Error(w, `{"error":"Error al crear la reseña"}`, http.StatusInternalServerError)
			return
		}

		// 5. Respuesta 201 Created con la reseña creada
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resena)
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
