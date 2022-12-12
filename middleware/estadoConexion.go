package middleware

import (
	"net/http"

	"github.com/jfuentesd1801/crud_personas/conexion"
)

// EstadoConexion es la funcion encarda de verificar el estado de la conexion con la bd
func EstadoConexion(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !conexion.ChecarConexion() {
			http.Error(w, "Error, conexion con la base de datos perdida", http.StatusConflict)
			return
		}

		next(w, r)
	}
}
