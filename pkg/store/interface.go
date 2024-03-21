package store

import (
	"github.com/Gaghyta/BackendIIIFinalGO/internal/odontologos"
)

type StoreInterface interface {
	// Read devuelve un odontologo por su id
	Read(id int) (odontologos.Odontologo, error)
	// Create agrega un nuevo odontologos
	Create(odontologo odontologos.Odontologo) error
	// Update actualiza un paciente
	Update(odontologo odontologos.Odontologo) error
	// Delete elimina un paciente
	Delete(id int) error
	// Exists verifica si un paciente existe
	Exists(matricula string) bool
}
