package conexion

import (
	"context"
	"log"
	"time"
)

func Eliminar(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	consulta := "DELETE FROM persona WHERE id_persona = ?"
	_, err := ConexionBD.QueryContext(ctx, consulta, id)
	if err != nil {
		log.Println("Error al realizar el delete : " + err.Error())
	}
	return err
}
