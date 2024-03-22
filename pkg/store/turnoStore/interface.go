package turnoStore

import "github.com/Gaghyta/BackendIIIFinalGO/internal/domains"

type StoreInterface interface {
	// Read devuelve un paciente por su id
	Read(id int) (domains.Turno, error)
	// Create agrega un nuevo paciente
	Create(paciente domains.Turno) error
	// Update actualiza un paciente
	Update(paciente domains.Turno) error
	// Delete elimina un paciente
	Delete(id int) error
	// Exists verifica si un paciente existe
	Exists(dni string) bool
}
