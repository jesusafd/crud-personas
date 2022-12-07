package conexion

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

// ConexionBD es el objeto conexion con la base de datos
var ConexionBD *sqlx.DB

// NewConexion Funcion encargada de crear la conexion con la
// base de datos
func NuevaConexion() error {
	ConfigBD, err := NuevaConfiguracionBD()
	if err != nil {
		// En caso de haber un error la funcion NuevaConfiguracionBD
		// lo notificara en la consola
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var datosBD string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", ConfigBD.User, ConfigBD.Password, ConfigBD.Host, ConfigBD.Port, ConfigBD.Name)

	var bd *sqlx.DB
	// Creamos la conexion con la BD
	bd, err = sqlx.ConnectContext(ctx, "mysql", datosBD)
	if err != nil {
		log.Println("Error al crear la conexion con la BD\n" + err.Error())
		return err
	}
	// Seteamos el numero maximo de conexiones
	bd.SetMaxOpenConns(5)
	// La creacion de la conexion fue exitosa
	ConexionBD = bd
	return nil
}

// ChecarConexion permite comprobar el estado de la conexion
func ChecarConexion() bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Hacemos un ping a la bd para saber el estado de esta
	err := ConexionBD.PingContext(ctx)
	if err != nil {
		log.Println("Conexion con la base de datos fallida" + err.Error())
		return false
	}
	return true
}
