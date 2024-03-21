package turnos

type Turno struct {
	TurnosId           int    `json:"turnos_id"`
	FechaYHora         string `json:"fecha_y_hora"`
	Descripcion        string `json:"descripcion"`
	DentistaIDDentista int    `json:"dentista_id_dentista"`
	PacienteIDPaciente int    `json:"paciente_id_paciente"`
}
