package odontologoStore

import "github.com/Gaghyta/BackendIIIFinalGO/internal/domains"

type StoreInterface interface {
	// Read devuelve un odontologo por su id
	Read(id int) (domains.Odontologo, error)
	// Create agrega un nuevo odontologo
	Create(odontologo domains.Odontologo) error
	// Update actualiza un odontologo
	Update(odontologo domains.Odontologo) error
	// Delete elimina un odontologo
	Delete(id int) error
	// Exists verifica si un odontologo existe
	Exists(matricula string) bool
}
