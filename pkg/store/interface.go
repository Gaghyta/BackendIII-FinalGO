package store

import "github.com/Gaghyta/BackendIIIFinalGO/internal/domains"

type StoreInterface interface {
	Read(id int) (domains.Odontologo, error)
	Create(odontologo domains.Odontologo) error
	Update(odontologo domains.Odontologo) error
	Delete(id int) error
	Exists(matricula string) bool
	PatchMatricula(matricula string, nuevaMatricula string) error
}
