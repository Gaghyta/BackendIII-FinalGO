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

package turnoStore

import (
	"database/sql"
	"errors"
	"log"

	//"github.com/Gaghyta/BackendIIIFinalGO/internal/odontologos"

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
func (s *sqlStore) Read(id int) (domains.Turno, error) {
	// Consulta para recuperar el odontólogo con el ID proporcionado
	query := "SELECT fecha_y_hora, descripcion, idDentista, idPaciente FROM turnos WHERE id = ?"

	// Ejecutar la consulta y recuperar los datos
	var t domains.Turno
	err := s.db.QueryRow(query, id).Scan(&t.FechaYHora, &t.Descripcion, &t.DentistaIDDentista, t.PacienteIDPaciente)
	if err != nil {
		if err == sql.ErrNoRows {
			// El odontólogo con el ID proporcionado no fue encontrado
			return domains.Turno{}, errors.New("turno no encontrado")
		}
		// Ocurrió un error al ejecutar la consulta
		return domains.Turno{}, err
	}

	// Retornar los datos del odontólogo recuperado
	return t, nil
}

func (s *sqlStore) Create(t domains.Turno) error {
	query := "INSERT INTO turno (fecha_y_hora, descripcion, idPaciente, idodontologo) VALUES (?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(t.FechaYHora, t.Descripcion, t.PacienteIDPaciente, t.DentistaIDDentista)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *sqlStore) Update(t domains.Turno) error {
	return errors.New("not implemented yet")
}

func (s *sqlStore) Delete(id int) error {
	return errors.New("not implemented yet")
}

func (s *sqlStore) Exists(fecha_y_hora string, odontologo int) bool {
	query := "SELECT fecha_y_hora, idDentista FROM turnos WHERE fecha_y_hora = ? AND id_odontologo = ?"

	// Ejecutar la consulta y recuperar los datos
	var t domains.Turno
	err := s.db.QueryRow(query, fecha_y_hora, odontologo).Scan(&t.FechaYHora, &t.DentistaIDDentista)

	if err == sql.ErrNoRows {
		// El turno con fecha y odontologo proporcionado no fue encontrado
		return true
	}
	return false
}
