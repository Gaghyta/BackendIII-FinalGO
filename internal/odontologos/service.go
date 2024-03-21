package odontologos

import (
	"github.com/Gaghyta/BackendIIIFinalGO/internal/domains"
)

type Service interface {
	GetByID(id int) (domains.Odontologo, error)
	Create(o domains.Odontologo) (domains.Odontologo, error)
	Delete(id int) error
	Update(id int, o domains.Odontologo) (domains.Odontologo, error)
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
	if uO.Name != "" {
		o.Name = uO.Name
	}
	if uO.ApellidoOdontologo != "" {
		o.ApellidoOdontologo = uO.ApellidoOdontologo
	}
	if uO.matricula != "" {
		o.matricula = uO.matricula
	}
	o, err = s.r.Update(id, o)
	if err != nil {
		return domains.Odontologo{}, err
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
