package odontologos

import (
	"errors"

	"github.com/Gaghyta/BackendIIIFinalGO/pkg/store"

	"github.com/Gaghyta/BackendIIIFinalGO/internal/domains"
)

type Repository interface {
	// GetByID busca un paciente por su id
	GetByID(id int) (domains.Odontologo, error)
	// Create agrega un nuevo paciente
	Create(p domains.Odontologo) (domains.Odontologo, error)
	// Update actualiza un paciente
	Update(id int, p domains.Odontologo) (domains.Odontologo, error)
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

func (r *repository) GetByID(id int) (domains.Odontologo, error) {
	odontologo, err := r.storage.Read(id)
	if err != nil {
		return domains.Odontologo{}, err
	}
	return odontologo, nil

}

func (r *repository) Create(o domains.Odontologo) (domains.Odontologo, error) {
	if !r.storage.Exists(o.Matricula) {
		return domains.Odontologo{}, errors.New("la matrícula de este odontologo ya existe en nuestra base de datos")
	}
	err := r.storage.Create(o)
	if err != nil {
		return domains.Odontologo{}, errors.New("error guardando odontólogo")
	}
	return o, nil
}

func (r *repository) Update(id int, o domains.Odontologo) (domains.Odontologo, error) {
	if !r.storage.Exists(o.Matricula) {
		return domains.Odontologo{}, errors.New("la matrícula ingresado existe")
	}
	err := r.storage.Update(o)
	if err != nil {
		return domains.Odontologo{}, errors.New("error modificando el odontólogo")
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
