package odontologo

import (
	"github.com/Gaghyta/BackendIIIFinalGO/internal/domain"
)

type Service interface {
	GetByID(id int) (domain.Odontologo, error)
	Create(o domain.Odontologo) (domain.Odontologo, error)
	Delete(id int) error
	Update(id int, o domain.Odontologo) (domain.Odontologo, error)
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetByID(id int) (domain.Odontologo, error) {
	o, err := s.r.GetByID(id)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return o, nil
}

func (s *service) Create(o domain.Odontologo) (domain.Odontologo, error) {
	o, err := s.r.Create(o)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return o, nil
}
func (s *service) Update(id int, uO domain.Odontologo) (domain.Odontologo, error) {
	o, err := s.r.GetByID(id)
	if err != nil {
		return domain.Odontologo{}, err
	}
	if uO.NombreOdontologo != "" {
		o.NombreOdontologo = uO.NombreOdontologo
	}
	if uO.ApellidoOdontologo != "" {
		o.ApellidoOdontologo = uO.ApellidoOdontologo
	}
	if uO.Matricula != "" {
		o.Matricula = uO.Matricula
	}
	o, err = s.r.Update(id, o)
	if err != nil {
		return domain.Odontologo{}, err
	}
	return o, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
