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
	// Consulta para recuperar el paciente con el ID proporcionado
	query := "SELECT nombre, apellido, domicilio, dni, fecha_de_alta FROM pacientes WHERE idpaciente = ?"

	// Ejecutar la consulta y recuperar los datos
	var p domains.Paciente
	err := s.db.QueryRow(query, id).Scan(&p.NombrePaciente, &p.ApellidoPaciente, &p.DomicilioPaciente,
		&p.Dni, p.FechaDeAlta)
	if err != nil {
		if err == sql.ErrNoRows {
			// El paciente con el ID proporcionado no fue encontrado
			return domains.Paciente{}, errors.New("paciente no encontrado")
		}
		// Ocurrió un error al ejecutar la consulta
		return domains.Paciente{}, err
	}

	// Retornar los datos del paciente recuperado
	return p, nil
}

func (s *sqlStore) Create(p domains.Paciente) error {
	query := "INSERT INTO pacientes (nombre, apellido, domicilio, dni, fecha_de_alta) VALUES (?, ?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(p.NombrePaciente, p.ApellidoPaciente, p.DomicilioPaciente, p.Dni, p.FechaDeAlta)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *sqlStore) Update(p domains.Paciente) error {
	query := "UPDATE pacientes SET nombre = ?, apellido = ?, domicilio = ?, dni = ?, fecha_de_alta = ? WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(p.NombrePaciente, p.ApellidoPaciente, p.DomicilioPaciente, p.Dni, p.FechaDeAlta, p.PacienteID)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) Delete(id int) error {
	query := "DELETE FROM pacientes WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) Exists(dni string) bool {
	var exists bool
	var id int
	query := "SELECT idpaciente FROM pacientes WHERE dni = ?;"
	row := s.db.QueryRow(query, dni)
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	if id > 0 {
		exists = true
	}
	return exists
}
