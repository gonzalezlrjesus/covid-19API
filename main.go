package main

import (
	"covid-19API/rutas"
	"fmt"
	"net/http"
	"os"

	// Gorm Postgres dialects
	"github.com/rs/cors"
)

func main() {

	// Routes
	routes := rutas.Routes()

	// Server port reading
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	// CORS configuration
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200", "*", "http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Origin", "X-Requested-With", "Content-Length", "Accept-Encoding", "Cache-Control", "Authorization"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	handler := c.Handler(routes)

	// Server listening
	err := http.ListenAndServe(":"+port, handler) //Launch the app.
	if err != nil {
		fmt.Print(err)
	}
}
