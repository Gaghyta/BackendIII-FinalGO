package paciente

import (
	"github.com/Gaghyta/BackendIIIFinalGO/internal/domains"
)

type Service interface {
	GetByID(id int) (domains.Paciente, error)
	GetByDNI(dni string) (domains.Paciente, error)
	Create(o domains.Paciente) (domains.Paciente, error)
	Update(id int, o domains.Paciente) (domains.Paciente, error)
	Delete(id int) error
	Patch(id int, p domains.Paciente) (domains.Paciente, error)
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetByID(id int) (domains.Paciente, error) {
	o, err := s.r.GetByID(id)
	if err != nil {
		return domains.Paciente{}, err
	}
	return o, nil
}

func (s *service) GetByDNI(dni string) (domains.Paciente, error) {
	o, err := s.r.GetByDNI(dni)
	if err != nil {
		return domains.Paciente{}, err
	}
	return o, nil
}

func (s *service) Create(o domains.Paciente) (domains.Paciente, error) {
	o, err := s.r.Create(o)
	if err != nil {
		return domains.Paciente{}, err
	}
	return o, nil
}
func (s *service) Update(id int, uP domains.Paciente) (domains.Paciente, error) {
	o, err := s.r.GetByID(id)
	if err != nil {
		return domains.Paciente{}, err
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
		return domains.Paciente{}, err
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

func (s *service) Patch(id int, uP domains.Paciente) (domains.Paciente, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domains.Paciente{}, err
	}
	if uP.NombrePaciente != "" {
		p.NombrePaciente = uP.NombrePaciente
	}
	if uP.ApellidoPaciente != "" {
		p.ApellidoPaciente = uP.ApellidoPaciente
	}
	if uP.DomicilioPaciente != "" {
		p.DomicilioPaciente = uP.DomicilioPaciente
	}
	if uP.Dni != "" {
		p.Dni = uP.Dni
	}
	if uP.FechaDeAlta != "" {
		p.FechaDeAlta = uP.FechaDeAlta
	}
	p, err = s.r.Update(id, p)
	if err != nil {
		return domains.Paciente{}, err
	}
	return p, nil
}
