package domain

type Turno struct {
	TurnosId           int    `json:"turnos_id"`
	FechaYHora         string `json:"fecha_y_hora"`
	Descripcion        string `json:"descripcion"`
	OdontologoId       int    `json:"odontologo"`
	PacienteIDPaciente int    `json:"paciente_id_paciente"`
}
