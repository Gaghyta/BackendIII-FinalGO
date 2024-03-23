package odontologos

import (
	"errors"

	"github.com/Gaghyta/BackendIIIFinalGO/internal/domains"

	"fmt"
)

type Service interface {
	GetByID(id int) (domains.Odontologo, error)
	Create(o domains.Odontologo) (domains.Odontologo, error)
	Delete(id int) error
	Update(id int, o domains.Odontologo) (domains.Odontologo, error)
	Patch(matricula string, nuevaMatricula string) (domains.Odontologo, error)
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetByID(id int) (domains.Odontologo, error) {
	o, err := s.r.GetByID(id)
	if err != nil {
		return domains.Odontologo{}, err
	}
	return o, nil
}

func (s *service) Create(o domains.Odontologo) (domains.Odontologo, error) {
	o, err := s.r.Create(o)
	if err != nil {
		return domains.Odontologo{}, err
	}
	return o, nil
}
func (s *service) Update(id int, uO domains.Odontologo) (domains.Odontologo, error) {
	o, err := s.r.GetByID(id)

	if err != nil {
		return domains.Odontologo{}, err
	}
	if uO.ApellidoOdontologo != "" {
		o.ApellidoOdontologo = uO.ApellidoOdontologo
	}
	if uO.NombreOdontologo != "" {
		o.NombreOdontologo = uO.NombreOdontologo
	}
	if uO.Matricula != "" {
		o.Matricula = uO.Matricula
	}
	o, err = s.r.Update(id, o)
	if err != nil {
		return domains.Odontologo{}, err
	}
	return o, nil
}

func (s *service) Delete(id int) error {

	err := s.r.Delete(id)
	if id <= 0 {
		return errors.New("ID inválido")
	}

	if err != nil {
		return err
	}
	return nil
}

func (s *service) Patch(matricula string, nuevaMatricula string) (domains.Odontologo, error) {

	odontologo, err := s.r.GetByMatricula(matricula)
	if err != nil {
		return domains.Odontologo{}, fmt.Errorf("error en patch service al obtener odontólogo")
	}

	// Actualizar la matrícula del odontólogo
	odontologo.Matricula = nuevaMatricula

	if odontologo.Matricula == "" {
		return domains.Odontologo{}, fmt.Errorf("odontólogo con matrícula no encontrada")
	}

	// Actualizar el odontólogo en la base de datos
	_, err = s.r.Update(odontologo.OdontologoId, odontologo)
	if err != nil {
		return domains.Odontologo{}, fmt.Errorf("error al actualizar")
	}

	return odontologo, nil
}
