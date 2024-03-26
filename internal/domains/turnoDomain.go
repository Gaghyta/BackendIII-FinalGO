package domains

type Turno struct {
	TurnosId           int    `json:"turno_id"`
	FechaYHora         string `json:"fecha_y_hora"`
	Descripcion        string `json:"descripcion"`
	DentistaIDDentista int    `json:"dentista_id_dentista"`
	PacienteIDPaciente int    `json:"paciente_id_paciente"`
}
