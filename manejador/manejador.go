package manejador

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jfuentesd1801/crud_personas/middleware"
	"github.com/jfuentesd1801/crud_personas/rutas"
	"github.com/rs/cors"
)

func Manejador() {
	ruta := mux.NewRouter()

	ruta.HandleFunc("/persona", middleware.EstadoConexion(rutas.InsertarPersona)).Methods("POST")

	PORT := "8080"

	manejador := cors.AllowAll().Handler(ruta)
	log.Fatal((http.ListenAndServe(":"+PORT, manejador)))
}
