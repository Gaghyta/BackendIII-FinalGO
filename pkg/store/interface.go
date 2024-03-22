package store

import "github.com/Gaghyta/BackendIIIFinalGO/internal/domains"

type StoreInterface interface {
	Read(id int) (domains.Odontologo, error)
	Create(odontologo domains.Odontologo) error
	Update(odontologo domains.Odontologo) error
	Delete(id int) error
	Exists(matricula string) bool
	Patch(matricula string, nuevaMatricula string) (domains.Odontologo, error)
	GetByMatricula(matricula string) (domains.Odontologo, error)
}
