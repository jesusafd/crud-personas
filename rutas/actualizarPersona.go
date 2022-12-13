package rutas

import (
	"encoding/json"
	"net/http"

	"github.com/jfuentesd1801/crud_personas/conexion"
	"github.com/jfuentesd1801/crud_personas/modelos"
)

func ActualizarPersona(w http.ResponseWriter, r *http.Request) {
	var persona modelos.Persona
	// Decodificamos el cuerpo de la peticion en el objeto persona
	json.NewDecoder(r.Body).Decode(&persona)

	err := conexion.Actualizar(persona)

	if err != nil {
		http.Error(w, "Error al realizar la actualizacion del registro : "+err.Error(), http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
