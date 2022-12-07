package main

import (
	"log"

	"github.com/jfuentesd1801/crud_personas/conexion"
)

func main() {
	// Cerramos la conexion a la bd al finalizar la ejecucion del main
	log.Println("Entre")
	err := conexion.NuevaConexion()
	if err != nil {
		log.Println("No se pudo establecer conxion con la bd")
	}
	defer conexion.ConexionBD.Close()

	log.Println("Conexion exitosa")
}
