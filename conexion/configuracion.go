package conexion

import (
	_ "embed"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v3"
)

// configuracionBD es la estructura usitlizada para
// extraer los datos necesarios para la conexion con
// la BD, del archivo configuracion.yaml
type ConfiguracionBD struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

// archivoConfiguracion es la varible con la cual se extraera los
// datos del archivo .yaml

//go:embed configuracion.yaml
var archivoConfiguracion []byte

// NuevaConfiguracionBD es la funcion encargada de leer los datos
// del archivo .yaml y crear un objeto de tipo ConfiguracionBD el
// cula permita manipualar estos datos
func NuevaConfiguracionBD() (ConfiguracionBD, error) {
	var configBD ConfiguracionBD

	err := yaml.Unmarshal(archivoConfiguracion, &configBD)
	if err != nil {
		log.Println("Error al extraer los datos de configuracion de la bd" + err.Error())
	}
	return configBD, err
}
