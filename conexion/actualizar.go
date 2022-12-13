package conexion

import (
	"context"
	"log"
	"time"

	"github.com/jfuentesd1801/crud_personas/modelos"
)

func Actualizar(persona modelos.Persona) error {
	// Al colocar un contexto con WithTimeout evitamos que la app se quede colgada, debido a que se
	// cancela al pasar los 5 segundos
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	consulta := "UPDATE persona SET nombre=?,apellido=?,edad=?,direccion=?,email=? WHERE id_persona=?"
	_, err := ConexionBD.ExecContext(ctx, consulta, persona.Nombre, persona.Apellido, persona.Edad, persona.Direccion, persona.Email, persona.ID_persona)
	if err != nil {
		log.Println("Error al actualizar los datos en la base de datos : " + err.Error())
	}
	return err
}
