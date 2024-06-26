package turnos

import (
	"errors"

	"github.com/Gaghyta/BackendIIIFinalGO/internal/domains"

	turnoStore "github.com/Gaghyta/BackendIIIFinalGO/pkg/store"
)

type Repository interface {
	// GetByID busca un paciente por su id
	GetByID(id int) (domains.Turno, error)
	// GetByDNI busca un turno por su DNI
	GetByDNI(dni string) (domains.Turno, error)
	// Create hace un post
	Create(t domains.Turno) (domains.Turno, error)
	// Update actualiza un paciente
	Update(id int, p domains.Turno) (domains.Turno, error)
	// Delete elimina un paciente
	Delete(id int) error
}

type repository struct {
	storage turnoStore.StoreTurnoInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage turnoStore.StoreTurnoInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetByID(id int) (domains.Turno, error) {
	objturno, err := r.storage.ReadById(id)
	if err != nil {
		return domains.Turno{}, errors.New("el turno buscado no existe")
	}
	return objturno, nil

}

func (r *repository) GetByDNI(dni string) (domains.Turno, error) {
	objturno, err := r.storage.GetByDNI(dni)
	if err != nil {
		return domains.Turno{}, errors.New("el turno buscado no existe")
	}
	return objturno, nil

}

func (r *repository) Create(t domains.Turno) (domains.Turno, error) {
	if r.storage.Exists(t.FechaYHora, t.DentistaIDDentista) {
		return domains.Turno{}, errors.New("el odontólogo tiene un turno asignado en ese horario en nuestra base de datos")
	}
	newT, err := r.storage.Create(t)
	if err != nil {
		return domains.Turno{}, errors.New("error guardando turno")
	}
	return newT, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(id int, o domains.Turno) (domains.Turno, error) {
	if r.storage.Exists(o.FechaYHora, o.DentistaIDDentista) {
		return domains.Turno{}, errors.New("ya existe un turno con esa hora y odontologo")
	}
	t, err := r.storage.Update(id, o)
	if err != nil {
		return domains.Turno{}, errors.New("error modificando el turno")
	}
	return t, nil
}
