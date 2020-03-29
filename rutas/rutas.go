package rutas

import (
	"covid-19API/controladores"

	"github.com/gorilla/mux"
)

// Routes all routes
func Routes() *mux.Router {
	router := mux.NewRouter()

	// API VERSION
	s := router.PathPrefix("/v1").Subrouter()

	// -------------------Rutas Diferentes Casos---------------------------------------

	// Get casos-confirmados.
	s.HandleFunc("/casos-confirmados", controladores.GetCasosConfirmados).Methods("GET")

	// Get casos-muertos.
	s.HandleFunc("/casos-muertos", controladores.GetCasosMuertos).Methods("GET")

	// Get casos-recuperados.
	s.HandleFunc("/casos-recuperados", controladores.GetCasosRecuperados).Methods("GET")

	return router
}
