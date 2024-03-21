package store

import "github.com/Gaghyta/BackendIIIFinalGO/internal/pacientes"

type StoreInterface interface {
	// Read devuelve un paciente por su id
	Read(id int) (pacientes.Pacientes, error)
	// Create agrega un nuevo paciente
	Create(paciente pacientes.Pacientes) error
	// Update actualiza un paciente
	Update(paciente pacientes.Pacientes) error
	// Delete elimina un paciente
	Delete(id int) error
	// Exists verifica si un paciente existe
	Exists(dni string) bool
}
