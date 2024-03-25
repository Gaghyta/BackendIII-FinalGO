package odontologoStore

import "github.com/Gaghyta/BackendIIIFinalGO/internal/domains"

type StoreOdontologoInterface interface {
	// Read devuelve un odontólogo por su id
	Read(id int) (domains.Odontologo, error)
	// Create agrega un nuevo odontólogo
	Create(odontologo domains.Odontologo) error
	// Update actualiza un odontólogo
	Update(odontologo domains.Odontologo) error
	// Delete elimina un odontólogo
	Delete(id int) error
	// Exists verifica si un odontólogo existe
	Exists(matricula string) bool
}

type StorePacienteInterface interface {
	// Read devuelve un paciente por su id
	Read(id int) (domains.Paciente, error)
	// Create agrega un nuevo paciente
	Create(paciente domains.Paciente) error
	// Update actualiza un paciente
	Update(paciente domains.Paciente) error
	// Delete elimina un paciente
	Delete(id int) error
	// Exists verifica si un paciente existe
	Exists(dni string) bool
}

type StoreTurnoInterface interface {
	// Read devuelve un turno por su id
	Read(id int) (domains.Turno, error)
	// Create agrega un nuevo turno
	Create(turno domains.Turno) error
	// Update actualiza un turno
	Update(turno domains.Turno) error
	// Delete elimina un turno
	Delete(id int) error
	// Exists verifica si un turno existe
	Exists(fecha_y_hora string, odontologo int) bool
}
