/* package store

import (
	"database/sql"
	"errors"
	"log"

	"github.com/Gaghyta/BackendIIIFinalGO/internal/odontologos"
	"github.com/Gaghyta/BackendIIIFinalGO/internal/pacientes"
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
func (s *sqlStore) Read(p odontologos.Odontologos) error {
	return odontologos.Odontologos{}, nil
}

func (s *sqlStore) Create(id int) (odontologos.Odontologos, error) {
	query := "INSERT INTO odontologos (apellido_odontologo, , nombre_odontologo, matricula) VALUES (?, ?, ?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(odontologos.ApellidoOdontologo, odontologos.NombreOdontologo, odontologos.Matricula)
	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (s *sqlStore) Update(p pacientes.Pacientes) error {
	return errors.New("not implemented yet")
}

func (s *sqlStore) Delete(id int) error {
	return errors.New("not implemented yet")
}

func (s *sqlStore) Exists(dni string) bool {
	return bool(true)
}
*/

package pacienteStore

import (
	"database/sql"
	"errors"
	"log"


	"github.com/Gaghyta/BackendIIIFinalGO/internal/domains"

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
func (s *sqlStore) Read(id int) (domains.Paciente, error) {
	// Consulta para recuperar el odontólogo con el ID proporcionado
	query := "SELECT apellido_odontologo, nombre_odontologo, matricula FROM odontologos WHERE id = ?"

	// Ejecutar la consulta y recuperar los datos
	var o domains.Paciente
	err := s.db.QueryRow(query, id).Scan(&o.ApellidoPaciente, &o.NombrePaciente, &o.Dni)
	if err != nil {
		if err == sql.ErrNoRows {
			// El odontólogo con el ID proporcionado no fue encontrado
			return domains.Paciente{}, errors.New("paciente no encontrado")
		}
		// Ocurrió un error al ejecutar la consulta
		return domains.Paciente{}, err
	}

	// Retornar los datos del paciente recuperado
	return o, nil
}

func (s *sqlStore) Create(p domains.Paciente) error {
	query := "INSERT INTO pacientes (nombre, apellido, domicilio, dni, fecha_de_alta) VALUES (?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(p.ApellidoPaciente, p.NombrePaciente, p.DomicilioPaciente, p.Dni, p.FechaDeAlta)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *sqlStore) Update(p domains.Paciente) error {
	return errors.New("not implemented yet")
}

func (s *sqlStore) Delete(id int) error {
	return errors.New("not implemented yet")
}

func (s *sqlStore) Exists(dni string) bool {
	return true
}