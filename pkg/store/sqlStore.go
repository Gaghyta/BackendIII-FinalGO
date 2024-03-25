package odontologoStore

import (
	"database/sql"
	"errors"

	//"fmt"
	"log"

	"github.com/Gaghyta/BackendIIIFinalGO/internal/domains"

	_ "github.com/go-sql-driver/mysql"
)

type OdontologoSqlStore struct {
	db *sql.DB
}

func NewOdontologoSqlStore(db *sql.DB) StoreOdontologoInterface {
	return &OdontologoSqlStore{
		db: db,
	}
}

type PacienteSqlStore struct {
	db *sql.DB
}

func NewPacienteSqlStore(db *sql.DB) StorePacienteInterface {
	return &PacienteSqlStore{
		db: db,
	}
}

type TurnoSqlStore struct {
	db *sql.DB
}

func NewTurnoSqlStore(db *sql.DB) StoreTurnoInterface {
	return &TurnoSqlStore{
		db: db,
	}
}

// **********************************************************************
// 3 IMPLEMENTACIÓN DE MÉTODOS DE LA INTERFACE DE ODONTOLOGOS

func (s *OdontologoSqlStore) Read(id int) (domains.Odontologo, error) {
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

func (s *OdontologoSqlStore) Create(p domains.Odontologo) error {
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

func (s *OdontologoSqlStore) Update(o domains.Odontologo) error {
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

func (s *OdontologoSqlStore) Delete(id int) error {
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

func (s *OdontologoSqlStore) Exists(matricula string) bool {
	var exists bool
	var id int
	query := "SELECT odontologo_id FROM odontologos WHERE matricula = ?;"
	row := s.db.QueryRow(query, matricula)
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	if id > 0 {
		exists = true
	}
	return exists
}

// **********************************************************************
// 2 IMPLEMENTACIÓN DE MÉTODOS DE LA INTERFACE DE PACIENTES

func (s *PacienteSqlStore) Read(id int) (domains.Paciente, error) {
	// Consulta para recuperar el paciente con el ID proporcionado
	query := "SELECT nombre, apellido, domicilio, dni, fecha_de_alta FROM pacientes WHERE paciente_id = ?"

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

func (s *PacienteSqlStore) Create(p domains.Paciente) error {
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

func (s *PacienteSqlStore) Update(p domains.Paciente) error {
	query := "UPDATE pacientes SET nombre = ?, apellido = ?, domicilio = ?, dni = ?, fecha_de_alta = ? WHERE paciente_id = ?;"
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

func (s *PacienteSqlStore) Delete(id int) error {
	query := "DELETE FROM pacientes WHERE paciente_id = ?;"
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

func (s *PacienteSqlStore) Exists(dni string) bool {
	var exists bool
	var id int
	query := "SELECT paciente_id FROM pacientes WHERE dni = ?;"
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

// **********************************************************************
// 3 IMPLEMENTACIÓN DE MÉTODOS DE LA INTERFACE DE TURNOS

func (s *TurnoSqlStore) Read(id int) (domains.Turno, error) {
	// Consulta para recuperar el turno con el ID proporcionado
	query := "SELECT dentista_id_dentista FROM turnos WHERE turnos_id = ?"

	// Ejecutar la consulta y recuperar los datos
	var t domains.Turno
	err := s.db.QueryRow(query, id).Scan(&t.TurnosId)
	if err != nil {
		if err == sql.ErrNoRows {
			// El paciente con el ID proporcionado no fue encontrado
			return domains.Turno{}, errors.New("turno no encontrado")
		}
		// Ocurrió un error al ejecutar la consulta
		return domains.Turno{}, err
	}

	// Retornar los datos del paciente recuperado
	return t, nil
}

func (s *TurnoSqlStore) Create(p domains.Turno) error {
	query := "INSERT INTO turnos (fecha_y_hora, descripcion, dentista_id_dentista, paciente_id_paciente) VALUES (?, ?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	var t domains.Turno

	_, err = stmt.Exec(t.FechaYHora, t.Descripcion, t.DentistaIDDentista, t.PacienteIDPaciente)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *TurnoSqlStore) Update(p domains.Turno) error {
	query := "UPDATE turnos SET fecha_y_hora = ?, descripcion = ?, dentista_id_dentista = ?, paciente_id_paciente = ? WHERE id = ?;"
	stmt, err := s.db.Prepare(query)
	var t domains.Turno
	if err != nil {
		return err
	}
	res, err := stmt.Exec(t.FechaYHora, t.Descripcion, t.DentistaIDDentista, t.PacienteIDPaciente)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *TurnoSqlStore) Delete(id int) error {
	query := "DELETE FROM turnos WHERE id = ?;"
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

func (s *TurnoSqlStore) Exists(fecha_y_hora string, odontologo int) bool {
	var exists bool
	var id int
	query := "SELECT turnos_id FROM pacientes WHERE fecha_y_hora = ?;"
	row := s.db.QueryRow(query, fecha_y_hora)
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	if id > 0 {
		exists = true
	}
	return exists
}
