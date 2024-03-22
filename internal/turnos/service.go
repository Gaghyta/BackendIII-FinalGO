package turnos

import (
	"github.com/Gaghyta/BackendIIIFinalGO/internal/domains"
)

type Service interface {
	GetByID(id int) (domains.Turno, error)
	Create(o domains.Turno) (domains.Turno, error)
	Delete(id int) error
	Update(id int, o domains.Turno) (domains.Turno, error)
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetByID(id int) (domains.Turno, error) {
	o, err := s.r.GetByID(id)
	if err != nil {
		return domains.Turno{}, err
	}
	return o, nil
}

func (s *service) Create(o domains.Turno) (domains.Turno, error) {
	o, err := s.r.Create(o)
	if err != nil {
		return domains.Turno{}, err
	}
	return o, nil
}
func (s *service) Update(id int, uP domains.Turno) (domains.Turno, error) {
	o, err := s.r.GetByID(id)
	if err != nil {
		return domains.Turno{}, err
	}
	if uP.NombrePaciente != "" {
		o.NombrePaciente = uP.NombrePaciente
	}
	if uP.ApellidoPaciente != "" {
		o.ApellidoPaciente = uP.ApellidoPaciente
	}
	if uP.Dni != "" {
		o.Dni = uP.Dni
	}
	o, err = s.r.Update(id, o)
	if err != nil {
		return domains.Turno{}, err
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
