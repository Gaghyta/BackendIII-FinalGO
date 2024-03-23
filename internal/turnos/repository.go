package turnos

import (
	"errors"

	"github.com/Gaghyta/BackendIIIFinalGO/internal/domains"
	"github.com/Gaghyta/BackendIIIFinalGO/pkg/store/turnoStore"
)

type Repository interface {
	// GetByID busca un paciente por su id
	GetByID(id int) (domains.Turno, error)
	Create(o domains.Turno) (domains.Turno, error)
	// Update actualiza un paciente
	Update(id int, p domains.Turno) (domains.Turno, error)
	// Delete elimina un paciente
	Delete(id int) error
}

type repository struct {
	storage turnoStore.StoreInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage turnoStore.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetByID(id int) (domains.Turno, error) {
	objturno, err := r.storage.Read(id)
	if err != nil {
		return domains.Turno{}, errors.New("El odontólogo buscado no existe")
	}
	return objturno, nil

}

func (r *repository) Create(t domains.Turno) (domains.Turno, error) {
	if r.storage.Exists(t.FechaYHora, t.DentistaIDDentista) {
		return domains.Turno{}, errors.New("El odontólogo tiene un turno asignado en ese horario en nuestra base de datos. Por favor, revíselo.")
	}
	err := r.storage.Create(t)
	if err != nil {
		return domains.Turno{}, errors.New("Error guardando turno")
	}
	return t, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(id int, o domains.Turno) (domains.Turno, error) {
	if !r.storage.Exists(o.FechaYHora, o.DentistaIDDentista) {
		return domains.Turno{}, errors.New("El DNI ingresado ya existe")
	}
	err := r.storage.Update(o)
	if err != nil {
		return domains.Turno{}, errors.New("Error modificando el paciente")
	}
	return o, nil
}