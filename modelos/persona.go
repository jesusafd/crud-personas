package modelos

// Persona es la clase modelo para la base de datos
type Persona struct {
	ID_persona int
	Nombre     string `json:"nombre"`
	Apellido   string `json:"apellido"`
	Edad       int    `json:"edad"`
	Direccion  string `json:"direccion"`
	Email      string `json:"email"`
}
