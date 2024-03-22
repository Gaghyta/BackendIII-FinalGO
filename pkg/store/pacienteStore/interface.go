package pacienteStore

import "github.com/Gaghyta/BackendIIIFinalGO/internal/domain"

type StoreInterface interface {
	// Read devuelve un paciente por su id
	Read(id int) (domain.Paciente, error)
	// Create agrega un nuevo paciente
	Create(paciente domain.Paciente) error
	// Update actualiza un paciente
	Update(paciente domain.Paciente) error
	// Delete elimina un paciente
	Delete(id int) error
	// Exists verifica si un paciente existe
	Exists(dni string) bool
}
