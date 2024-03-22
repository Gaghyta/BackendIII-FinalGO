package turnos

import (
	"github.com/Gaghyta/BackendIIIFinalGO/internal/domains"
)

type Service interface {
	GetByID(id int) (domains.Turno, error)
	Create(t domains.Turno) (domains.Turno, error)
	Delete(id int) error
	Update(id int, t domains.Turno) (domains.Turno, error)
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetByID(id int) (domains.Turno, error) {
	t, err := s.r.GetByID(id)
	if err != nil {
		return domains.Turno{}, err
	}
	return t, nil
}

func (s *service) Create(t domains.Turno) (domains.Turno, error) {
	t, err := s.r.Create(t)
	if err != nil {
		return domains.Turno{}, err
	}
	return t, nil
}
func (s *service) Update(id int, ut domains.Turno) (domains.Turno, error) {
	t, err := s.r.GetByID(id)
	if err != nil {
		return domains.Turno{}, err
	}
	if ut.FechaYHora != "" {
		t.FechaYHora = ut.FechaYHora
	}
	if ut.DentistaIDDentista != 0 {
		t.DentistaIDDentista = ut.DentistaIDDentista
	}
	if ut.PacienteIDPaciente != 0 {
		t.PacienteIDPaciente = ut.PacienteIDPaciente
	}
	t, err = s.r.Update(id, t)
	if err != nil {
		return domains.Turno{}, err
	}
	return t, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
