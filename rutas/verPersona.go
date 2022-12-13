package rutas

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jfuentesd1801/crud_personas/conexion"
	"github.com/jfuentesd1801/crud_personas/modelos"
)

func VerPersona(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	var persona modelos.Persona
	persona, err = conexion.Ver(id)

	if err != nil {
		http.Error(w, "Error al buscar registro en la base de datos : "+err.Error(), http.StatusConflict)
		return
	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(persona)
}
