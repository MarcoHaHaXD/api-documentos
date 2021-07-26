package controladores

import (
	"api-documentos/bd"
	"api-documentos/modelos"
	"encoding/json"
	"net/http"
)

// /grancompu/devoluciones	GET
func TodosDevolucionesEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(bd.TodosDevoluciones())
}

//grancompu/cotizacion POST
func InserCorizacionEndpoint(w http.ResponseWriter, r *http.Request) {
	var devolucion modelos.Cotizacion
	json.NewDecoder(r.Body).Decode(&devolucion)
	json.NewEncoder(w).Encode(bd.InserCorizacion(devolucion))
}