package rutas

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jfuentesd1801/crud_personas/conexion"
	"github.com/jfuentesd1801/crud_personas/modelos"
)

func InsertarPersona(w http.ResponseWriter, r *http.Request) {
	var persona modelos.Persona
	// Decodificamos los datos del request en el objeto de tipo persona
	err := json.NewDecoder(r.Body).Decode(&persona)

	if err != nil {
		log.Println("Error al decodificar el cuerpo de la peticion http")
		http.Error(w, "Error al decodificar el cuerpo de la peticion http", http.StatusConflict)
		return
	}
	// Se verifica que todos los campos sean validos
	if persona.Nombre == "" || persona.Apellido == "" {
		http.Error(w, "El nombre y el apellido son necesarios para realizar el registro", http.StatusBadRequest)
		return
	}

	err = conexion.Insertar(persona)

	if err != nil {
		log.Println("Error al realizar el registro en la base de datos : " + err.Error())
		http.Error(w, "Error al realizar el registro en la base de datos : "+err.Error(), http.StatusConflict)
	}

	w.WriteHeader(http.StatusAccepted)

}
