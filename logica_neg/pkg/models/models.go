package models

import "time"

// Definimos el molde de cliente para el handler.
type Cliente struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Usuario  string `json:"usuario"`
	Email    string `json:"email"`
}

// Definimos una instancia de clientes.
var Clientes = []Cliente{
	{"Tomas", "Di Carlo", "dica", "tdc@mail.com"},
	{"Bautista", "Cisilino", "cisa", "bc@mail.com"},
}

// Definimos un molde de resenas para el handler
type Resena struct {
	Titulo      string    `json:"titulo"`
	Descripcion string    `json:"descripcion"`
	Nota        int32     `json:"nota"`
	Fecha       time.Time `json:"fecha"`
}

var Resenas = []Resena{
	{"Centinela", "Gran lugar turistico", 9, time.Now()},
}
