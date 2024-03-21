package pacientes

type Pacientes struct {
	PacienteID        int    `json:"paciente_id"`
	NombrePaciente    string `json:"nombre_paciente"`
	ApellidoPaciente  string `json:"apellido_paciente"`
	DomicilioPaciente string `json:"domicilio_paciente"`
	DNI               string `json:"dni"`
	FechaDeAlta       string `json:"fecha_de_alta"`
}
