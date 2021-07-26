package main

import (
	"api-documentos/bd"
	"api-documentos/controladores"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	log.Println("El servidor se encuentra corriendo sin problema")
	rout()
}

func rout() {
	bd.NuevaConexionBD()
	defer bd.CerrarBD()

	gorillaRoute := mux.NewRouter().StrictSlash(true)

	gorillaRoute.HandleFunc("/grancompu/devoluciones", controladores.TodosDevolucionesEndpoint).Methods("GET")
	gorillaRoute.HandleFunc("/grancompu/cotizacion", controladores.InserCorizacionEndpoint).Methods("POST")

	gorillaRoute.HandleFunc("/mxgourmet/ordenCompra", controladores.InserOrdenEndpoint).Methods("POST")
	gorillaRoute.HandleFunc("/mxgourmet/ordenes", controladores.SelectOrdenesEndpoint).Methods("GET")

	//Config
	gorillaRoute.PathPrefix("/")

	http.Handle("/", gorillaRoute)

	handlerCORS := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
		/*MaxAge: 7200,
		OptionsPassthrough: true,
		Debug: true,*/
	}).Handler(gorillaRoute)

	//log.Fatal(http.ListenAndServe(":9090", handlerCORS))
	log.Fatal(http.ListenAndServe(":8083", handlerCORS))
}
