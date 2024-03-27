package odontologos

import (
	"errors"

	odontologoStore "github.com/Gaghyta/BackendIIIFinalGO/pkg/store"

	"github.com/Gaghyta/BackendIIIFinalGO/internal/domains"
)

type Repository interface {
	GetByID(id int) (domains.Odontologo, error)
	GetByMatricula(matricula string) (domains.Odontologo, error)
	Create(o domains.Odontologo) (domains.Odontologo, error)
	Update(id int, uO domains.Odontologo) (domains.Odontologo, error)
	Delete(id int) error
}

type repository struct {
	storage odontologoStore.StoreOdontologoInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage odontologoStore.StoreOdontologoInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetByID(id int) (domains.Odontologo, error) {
	odontologo, err := r.storage.Read(id)
	if err != nil {
		return domains.Odontologo{}, err
	}
	return odontologo, nil
}

func (r *repository) GetByMatricula(matricula string) (domains.Odontologo, error) {
	odontologo, err := r.storage.GetByMatricula(matricula)
	if err != nil {
		return domains.Odontologo{}, err
	}
	return odontologo, nil
}

func (r *repository) Create(o domains.Odontologo) (domains.Odontologo, error) {
	if r.storage.Exists(o.Matricula) {
		return domains.Odontologo{}, errors.New("la matrícula de este odontologo ya existe en nuestra base de datos")
	}
	nO, err := r.storage.Create(o)
	if err != nil {
		return domains.Odontologo{}, errors.New("error guardando odontólogo")
	}

	return nO, nil
}

func (r *repository) Update(id int, uO domains.Odontologo) (domains.Odontologo, error) {
	if r.storage.Exists(uO.Matricula) {
		return domains.Odontologo{}, errors.New("la matrícula ingresada existe")
	}
	modificado, err := r.storage.Update(id, uO)
	if err != nil {
		return domains.Odontologo{}, errors.New("error modificando el odontólogo")
	}
	return modificado, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
