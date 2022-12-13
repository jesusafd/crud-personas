package rutas

import (
	"net/http"
	"strconv"

	"github.com/jfuentesd1801/crud_personas/conexion"
)

func EliminarPerosona(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	err = conexion.Eliminar(id)

	if err != nil {
		http.Error(w, "Error al eliminar el registro de la base de bd : "+err.Error(), http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
