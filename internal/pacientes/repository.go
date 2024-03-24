package paciente

import (
	"errors"

	"github.com/Gaghyta/BackendIIIFinalGO/internal/domains"
	"github.com/Gaghyta/BackendIIIFinalGO/pkg/store/pacienteStore"
)

type Repository interface {
	// GetByID busca un paciente por su id
	GetByID(id int) (domains.Paciente, error)
	Create(o domains.Paciente) (domains.Paciente, error)
	// Update actualiza un paciente
	Update(id int, p domains.Paciente) (domains.Paciente, error)
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

func (r *repository) GetByID(id int) (domains.Paciente, error) {
	product, err := r.storage.Read(id)
	if err != nil {
		return domains.Paciente{}, errors.New("El paciente buscado no existe")
	}
	return product, nil
	
}

func (r *repository) Create(o domains.Paciente) (domains.Paciente, error) {
	if !r.storage.Exists(o.Dni) {
		return domains.Paciente{}, errors.New("El DNI ya existe en nuestra base de datos. Por favor, revíselo.")
	}
	err := r.storage.Create(o)
	if err != nil {
		return domains.Paciente{}, errors.New("Error guardando paciente")
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
		return domains.Paciente{}, errors.New("El DNI ingresado ya existe")
	}
	err := r.storage.Update(o)
	if err != nil {
		return domains.Paciente{}, errors.New("Error modificando el paciente")
	}
	return o, nil
}
