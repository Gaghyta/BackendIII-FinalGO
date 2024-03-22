package turnoStore

import "github.com/Gaghyta/BackendIIIFinalGO/internal/domains"

type StoreInterface interface {
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
