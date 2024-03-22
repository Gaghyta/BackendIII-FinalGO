package paciente

import (
	"errors"

	"github.com/Gaghyta/BackendIIIFinalGO/internal/domain"
	"github.com/Gaghyta/BackendIIIFinalGO/pkg/store/pacienteStore"
)

type Repository interface {
	// GetByID busca un paciente por su id
	GetByID(id int) (domain.Paciente, error)
	Create(o domain.Paciente) (domain.Paciente, error)
	// Update actualiza un paciente
	Update(id int, p domain.Paciente) (domain.Paciente, error)
	// Delete elimina un paciente
	Delete(id int) error
}

type repository struct {
	storage pacienteStore.StoreInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage pacienteStore.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetByID(id int) (domain.Paciente, error) {
	product, err := r.storage.Read(id)
	if err != nil {
		return domain.Paciente{}, errors.New("El odontólogo buscado no existe")
	}
	return product, nil

}

func (r *repository) Create(o domain.Paciente) (domain.Paciente, error) {
	if !r.storage.Exists(o.Dni) {
		return domain.Paciente{}, errors.New("El DNI ya existe en nuestra base de datos. Por favor, revíselo.")
	}
	err := r.storage.Create(o)
	if err != nil {
		return domain.Paciente{}, errors.New("Error guardando paciente")
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

func (r *repository) Update(id int, o domain.Paciente) (domain.Paciente, error) {
	if !r.storage.Exists(o.Dni) {
		return domain.Paciente{}, errors.New("El DNI ingresado ya existe")
	}
	err := r.storage.Update(o)
	if err != nil {
		return domain.Paciente{}, errors.New("Error modificando el paciente")
	}
	return o, nil
}
