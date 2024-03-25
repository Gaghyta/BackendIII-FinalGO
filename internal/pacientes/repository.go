package paciente

import (
	"errors"

	"github.com/Gaghyta/BackendIIIFinalGO/internal/domains"
	pacienteStore "github.com/Gaghyta/BackendIIIFinalGO/pkg/store"
)

type Repository interface {
	// GetByID busca un paciente por su id
	GetByID(id int) (domains.Paciente, error)
	// Create actualiza un paciente
	Create(o domains.Paciente) (domains.Paciente, error)
	// Update actualiza un paciente
	Update(id int, p domains.Paciente) (domains.Paciente, error)
	// Delete elimina un paciente
	Delete(id int) error
}

type repository struct {
	storage pacienteStore.StorePacienteInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage pacienteStore.StorePacienteInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetByID(id int) (domains.Paciente, error) {
	product, err := r.storage.Read(id)
	if err != nil {
		return domains.Paciente{}, errors.New("el paciente buscado no existe")
	}
	return product, nil

}

func (r *repository) Create(o domains.Paciente) (domains.Paciente, error) {
	if !r.storage.Exists(o.Dni) {
		return domains.Paciente{}, errors.New("el DNI ya existe en nuestra base de datos")
	}
	err := r.storage.Create(o)
	if err != nil {
		return domains.Paciente{}, errors.New("error guardando paciente")
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

func (r *repository) Update(id int, o domains.Paciente) (domains.Paciente, error) {
	if r.storage.Exists(o.Dni) {
		return domains.Paciente{}, errors.New("el DNI ingresado ya existe")
	}
	err := r.storage.Update(o)
	if err != nil {
		return domains.Paciente{}, errors.New("error modificando el paciente")
	}
	return o, nil
}
