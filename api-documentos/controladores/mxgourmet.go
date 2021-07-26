package controladores

import (
	"api-documentos/bd"
	"api-documentos/modelos"
	"encoding/json"
	"net/http"
)

//mxgourmet/ordenCompra POST
func InserOrdenEndpoint(w http.ResponseWriter, r *http.Request) {
	var orden modelos.OrdenCompra
	json.NewDecoder(r.Body).Decode(&orden)
	json.NewEncoder(w).Encode(bd.InserOrden(orden))
}

// /mxgourmet/ordenes	GET
func SelectOrdenesEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(bd.SelectOrdenes())
}