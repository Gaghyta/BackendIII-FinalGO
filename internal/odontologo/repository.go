package odontologo

import (
	"errors"

	"github.com/Gaghyta/BackendIIIFinalGO/internal/domain"

	"github.com/Gaghyta/BackendIIIFinalGO/pkg/store/odontologoStore"
)

type Repository interface {
	// GetByID busca un paciente por su id
	GetByID(id int) (domain.Odontologo, error)
	Create(o domain.Odontologo) (domain.Odontologo, error)
	// Update actualiza un paciente
	Update(id int, p domain.Odontologo) (domain.Odontologo, error)
	// Delete elimina un paciente
	Delete(id int) error
}

type repository struct {
	storage odontologoStore.StoreInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage odontologoStore.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetByID(id int) (domain.Odontologo, error) {
	od, err := r.storage.Read(id)
	if err != nil {
		return domain.Odontologo{}, errors.New("El odontólogo buscado no existe")
	}
	return od, nil

}

func (r *repository) Create(o domain.Odontologo) (domain.Odontologo, error) {
	if !r.storage.Exists(o.Matricula) {
		return domain.Odontologo{}, errors.New("La matrícula de este odontólogo ya existe en nuestra base de datos. Por favor, revíselo.")
	}
	err := r.storage.Create(o)
	if err != nil {
		return domain.Odontologo{}, errors.New("Error guardando odontólogo")
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

func (r *repository) Update(id int, o domain.Odontologo) (domain.Odontologo, error) {
	if !r.storage.Exists(o.Matricula) {
		return domain.Odontologo{}, errors.New("La matrícula ingresada ya existe")
	}
	err := r.storage.Update(o)
	if err != nil {
		return domain.Odontologo{}, errors.New("Error modificando el odontólogo")
	}
	return o, nil
}
