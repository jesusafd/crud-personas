package conexion

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/jfuentesd1801/crud_personas/modelos"
)

// Ver es la funcion que permite visualizar un usuario por su id
// Esta funcion devuelve un error en caso de no encontrar el registro
// o en caso de que no se pueda ejecutar la consulta, en caso de que el
// el error se nil, el registro se encontro
func Ver(id int) (modelos.Persona, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var persona modelos.Persona
	consulta := "SELECT * FROM persona WHERE id_persona = ?"
	row, err := ConexionBD.QueryContext(ctx, consulta, id)
	if err != nil {
		log.Println("Error al realizar la consulta en la bd : " + err.Error())
		return persona, err
	}

	// En caso de que la sentencia SELECT fuera exitosa es necesario saber si esta msima encontro algun registro que coincidiera
	// para lo cual usamos row.Next, que nos dice si hay una fila que leer
	if row.Next() {
		// extraemos los datos en el objeto
		err = row.Scan(&persona.ID_persona, &persona.Nombre, &persona.Apellido, &persona.Edad, &persona.Direccion, &persona.Email)
		if err != nil {
			log.Println("Error al leer datos de la columna de la bd : " + err.Error())
			return persona, err
		}
		// si el registro se encontro y no hubo promblemas deolvemos el registro persona y un error nil
		return persona, nil
	} else {
		// En caso de que no haya niguna fila que leer no creamos un error que indique esto
		log.Println("No se encontro ningun registro.")
		return persona, errors.New("no se encontro ningun usuario con el id indicado")
	}
}
