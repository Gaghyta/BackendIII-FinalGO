package odontologos

import (
	"errors"

	"github.com/Gaghyta/BackendIIIFinalGO/internal/domains/odontologos"
	"github.com/Gaghyta/BackendIIIFinalGO/pkg/store"
)

type Repository interface {
	// GetByID busca un paciente por su id
	GetByID(id int) (odontologos.Odontologos, error)
	// Create agrega un nuevo paciente
	Create(p odontologos.Odontologos) (odontologos.Odontologos, error)
	// Update actualiza un paciente
	Update(id int, p odontologos.Odontologos) (odontologos.Odontologos, error)
	// Delete elimina un paciente
	Delete(id int) error
}

type repository struct {
	storage store.StoreInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage store.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetByID(id int) (odontologos.Odontologos, error) {
	product, err := r.storage.Read(id)
	if err != nil {
		return odontologos.Odontologos{}, errors.New("El odontólogo buscado no existe")
	}
	return product, nil

}

func (r *repository) Create(o odontologos.Odontologos) (odontologos.Odontologos, error) {
	if !r.storage.Exists(o.dni) {
		return odontologos.Odontologos{}, errors.New("El dni de este odontólogo ya existe en nuestra base de datos. Por favor, revíselo.")
	}
	err := r.storage.Create(o)
	if err != nil {
		return odontologos.Odontologos{}, errors.New("Error guardando odontólogo")
	}
	return o, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(id int, o odontologos.Odontologos) (odontologos.Odontologos, error) {
	if !r.storage.Exists(o.dni) {
		return odontologos.Odontologos{}, errors.New("El dni ingresado existe")
	}
	err := r.storage.Update(o)
	if err != nil {
		return odontologos.Odontologos{}, errors.New("Error modificando el odontólogo")
	}
	return o, nil
}
