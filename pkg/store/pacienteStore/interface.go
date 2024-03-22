package pacienteStore

import "github.com/Gaghyta/BackendIIIFinalGO/internal/domains"

type StoreInterface interface {
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
