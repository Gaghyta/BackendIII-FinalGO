package odontologoStore

import "github.com/Gaghyta/BackendIIIFinalGO/internal/domain"

type StoreInterface interface {
	// Read devuelve un paciente por su id
	Read(id int) (domain.Odontologo, error)
	// Create agrega un nuevo paciente
	Create(paciente domain.Odontologo) error
	// Update actualiza un paciente
	Update(paciente domain.Odontologo) error
	// Delete elimina un paciente
	Delete(id int) error
	// Exists verifica si un paciente existe
	Exists(dni string) bool
}
