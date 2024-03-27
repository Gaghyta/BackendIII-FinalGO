package domains

type Paciente struct {
	PacienteID        int    `json:"paciente_id"`
	NombrePaciente    string `json:"nombre_paciente"`
	ApellidoPaciente  string `json:"apellido_paciente"`
	DomicilioPaciente string `json:"domicilio"`
	Dni               string `json:"dni"`
	FechaDeAlta       string `json:"fecha_de_alta"`
}
