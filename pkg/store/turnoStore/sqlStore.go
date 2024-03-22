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

	"github.com/Gaghyta/BackendIIIFinalGO/internal/domain"
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
func (s *sqlStore) Read(id int) (domain.Turno, error) {
	// Consulta para recuperar el odontólogo con el ID proporcionado
	query := "SELECT fecha_y_hora, descripcion, idDentista, idPaciente FROM turnos WHERE id = ?"

	// Ejecutar la consulta y recuperar los datos
	var t domain.Turno
	err := s.db.QueryRow(query, id).Scan(&t.FechaYHora, &t.Descripcion, &t.OdontologoId, t.PacienteIDPaciente)
	if err != nil {
		if err == sql.ErrNoRows {
			// El odontólogo con el ID proporcionado no fue encontrado
			return domain.Turno{}, errors.New("turno no encontrado")
		}
		// Ocurrió un error al ejecutar la consulta
		return domain.Turno{}, err
	}

	// Retornar los datos del odontólogo recuperado
	return t, nil
}

func (s *sqlStore) Create(t domain.Turno) error {
	query := "INSERT INTO turno (fecha_y_hora, descripcion, idPaciente, idodontologo) VALUES (?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(t.FechaYHora, t.Descripcion, t.PacienteIDPaciente, t.OdontologoId)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *sqlStore) Update(t domain.Turno) error {
	return errors.New("not implemented yet")
}

func (s *sqlStore) Delete(id int) error {
	return errors.New("not implemented yet")
}

func (s *sqlStore) Exists(dni string) bool {
	return true
}
