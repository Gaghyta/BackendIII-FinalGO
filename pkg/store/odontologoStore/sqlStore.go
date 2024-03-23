package ododontologoStore

import (
	"database/sql"
	"errors"
	"log"

	"github.com/Gaghyta/BackendIIIFinalGO/internal/domains"

	_ "github.com/go-sql-driver/mysql"
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

func (s *sqlStore) Read(id int) (domains.Odontologo, error) {
	// Consulta para recuperar el odontólogo con el ID proporcionado
	query := "SELECT apellido_odontologo, nombre_odontologo, matricula FROM odontologos WHERE odontologo_id = ?"

	// Ejecutar la consulta y recuperar los datos
	var o domains.Odontologo
	err := s.db.QueryRow(query, id).Scan(&o.ApellidoOdontologo, &o.NombreOdontologo, &o.Matricula)
	if err != nil {
		if err == sql.ErrNoRows {
			// El odontólogo con el ID proporcionado no fue encontrado
			return domains.Odontologo{}, errors.New("odontólogo no encontrado")
		}
		// Ocurrió un error al ejecutar la consulta
		return domains.Odontologo{}, err
	}

	// Retornar los datos del odontólogo recuperado
	return o, nil
}

func (s *sqlStore) Create(p domains.Odontologo) error {
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

func (s *sqlStore) Update(o domains.Odontologo) error {
	query := "UPDATE odontologos SET apellido_odontologo = ?, nombre_odontologo = ?, matricula = ? WHERE odontologo_id = ?"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(o.OdontologoId, o.ApellidoOdontologo, o.NombreOdontologo, o.Matricula)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *sqlStore) Delete(id int) error {
	query := "DELETE FROM odontologos WHERE odontologo_id = ?"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *sqlStore) Exists(dni string) bool {
	return true
}
