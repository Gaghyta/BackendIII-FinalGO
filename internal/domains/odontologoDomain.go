package domains

type Odontologo struct {
	OdontologoId       int    `json:"odontologo_id"`
	ApellidoOdontologo string `json:"apellido_odontologo"`
	NombreOdontologo   string `json:"nombre_odontologo"`
	Matricula          string `json:"matricula"`
}
