package conexion

import (
	"context"
	"log"
	"time"

	"github.com/jfuentesd1801/crud_personas/modelos"
)

// Insertar es la funcion encargada de insertar las personas en la base de datos
func Insertar(persona modelos.Persona) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var consulta string = "INSERT INTO persona (nombre,apellido,edad,direccion,email) VALUES (?,?,?,?,?)"
	// Ejecutamos la consulta en la bd, en caso de ocurrir un error se devolvera como resultado de la funcion insertar
	_, err := ConexionBD.ExecContext(
		ctx,
		consulta,
		persona.Nombre,
		persona.Apellido,
		persona.Edad,
		persona.Direccion,
		persona.Email)

	if err != nil {
		log.Println("Error al insertar la persona en la base de datos" + err.Error())
	}
	// Se devuelve el error sin importar si existe o es igual a nil
	return err
}
