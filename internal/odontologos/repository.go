package odontologos

import (
	"errors"
	"fmt"

	"github.com/Gaghyta/BackendIIIFinalGO/pkg/store"

	"github.com/Gaghyta/BackendIIIFinalGO/internal/domains"
)

type Repository interface {
	GetByID(id int) (domains.Odontologo, error)
	Create(p domains.Odontologo) (domains.Odontologo, error)
	Update(id int, p domains.Odontologo) (domains.Odontologo, error)
	Delete(id int) error
	//Patch(matricula string, nuevaMatricula string) (domains.Odontologo, error)
	GetByMatricula(matricula string) (domains.Odontologo, error)
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

func (r *repository) GetByMatricula(matricula string) (domains.Odontologo, error) {

	odontologo, err := r.storage.GetByMatricula(matricula)
	if err != nil {
		// Manejo de errores si ocurre algún problema al obtener el odontólogo por matrícula
		//return domains.Odontologo{}, fmt.Errorf("error al obtener odontólogo por matrícula %s: %s", matricula, err.Error())
		return domains.Odontologo{}, fmt.Errorf("error acá en GetByMatricula")
	}

	// Verificar si se encontró el odontólogo
	if odontologo.Matricula == "" {
		return domains.Odontologo{}, fmt.Errorf("odontólogo con matrícula %s no encontrado", matricula)
	}

	// Devolver el odontólogo encontrado
	return odontologo, nil
}

func (r *repository) Patch(matricula string) (domains.Odontologo, error) {
	odontologo, err := r.storage.GetByMatricula(matricula)
	if err != nil {
		// Manejo de error si ocurre algún problema al obtener el odontólogo por matrícula
		return domains.Odontologo{}, fmt.Errorf("HAYYYYY %s: %s", matricula, err.Error())
	}

	if odontologo.Matricula == "" {
		return domains.Odontologo{}, fmt.Errorf("odontólogo con matrícula %s no encontrado", matricula)
	}

	return odontologo, nil
}
