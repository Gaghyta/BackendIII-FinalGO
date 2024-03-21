package odontologos

import "github.com/Gaghyta/BackendIIIFinalGO/internal/domains/odontologos"

type Service interface {
	GetByID(id int) (odontologos.Odontologos, error)
	Create(o odontologos.Odontologos) (odontologos.Odontologos, error)
	Delete(id int) error
	Update(id int, o odontologos.Odontologos) (odontologos.Odontologos, error)
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetByID(id int) (odontologos.Odontologos, error) {
	o, err := s.r.GetByID(id)
	if err != nil {
		return odontologos.Odontologos{}, err
	}
	return o, nil
}

func (s *service) Create(o odontologos.Odontologos) (odontologos.Odontologos, error) {
	o, err := s.r.Create(o)
	if err != nil {
		return odontologos.Odontologos{}, err
	}
	return o, nil
}
func (s *service) Update(id int, uO odontologos.Odontologos) (odontologos.Odontologos, error) {
	o, err := s.r.GetByID(id)
	if err != nil {
		return odontologos.Odontologos{}, err
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
		return odontologos.Odontologos{}, err
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
