package store

import (
	"database/sql"
	"errors"
	"log"

	odontologos "github.com/Gaghyta/BackendIIIFinalGO/internal/domains"
	//"github.com/Gaghyta/BackendIIIFinalGO/internal/odontologos"
)

type sqlStore struct {
	db *sql.DB
}

func NewSqlStore(db *sql.DB) StoreInterface {
	return &sqlStore{
		db: db,
	}
}

// IMPLEMENTACIÓN DE MÉTODOS DE LA INTERFACE

func (s *sqlStore) Read(id int) (odontologos.Odontologos, error) {
	// Consulta para recuperar el odontólogo con el ID proporcionado
	query := "SELECT apellido_odontologo, nombre_odontologo, matricula FROM odontologos WHERE id = ?"

	// Ejecutar la consulta y recuperar los datos
	var o odontologos.Odontologos
	err := s.db.QueryRow(query, id).Scan(&o.ApellidoOdontologo, &o.NombreOdontologo, &o.Matricula)
	if err != nil {
		if err == sql.ErrNoRows {
			// El odontólogo con el ID proporcionado no fue encontrado
			return odontologos.Odontologos{}, errors.New("odontólogo no encontrado")
		}
		// Ocurrió un error al ejecutar la consulta
		return odontologos.Odontologos{}, err
	}

	// Retornar los datos del odontólogo recuperado
	return o, nil
}

func (s *sqlStore) Create(p odontologos.Odontologos) error {
	query := "INSERT INTO odontologos (apellido_odontologo, nombre_odontologo, matricula) VALUES (?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(p.ApellidoOdontologo, p.NombreOdontologo, p.Matricula)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *sqlStore) Update(o odontologos.Odontologos) error {
	return errors.New("not implemented yet")
}

func (s *sqlStore) Delete(id int) error {
	return errors.New("not implemented yet")
}

func (s *sqlStore) Exists(dni string) bool {
	return true
}
