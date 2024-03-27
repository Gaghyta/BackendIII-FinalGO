package Store

import (
	"database/sql"
	"errors"

	"fmt"
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
	o.OdontologoId = id

	// Retornar los datos del odontólogo recuperado
	return o, nil
}

func (s *OdontologoSqlStore) GetByID(id int) (domains.Odontologo, error) {
	var odontologoEncontrado domains.Odontologo

	query := fmt.Sprintf("SELECT * FROM odontologos WHERE odontologo_id = %d", id)
	row := s.db.QueryRow(query)
	err := row.Scan(&odontologoEncontrado.OdontologoId, &odontologoEncontrado.ApellidoOdontologo, &odontologoEncontrado.NombreOdontologo, &odontologoEncontrado.Matricula)
	if err != nil {
		return domains.Odontologo{}, err
	}
	return odontologoEncontrado, nil
}

func (s *OdontologoSqlStore) GetByMatricula(matricula string) (domains.Odontologo, error) {
	var odontologoEncontrado domains.Odontologo

	query := fmt.Sprintf("SELECT * FROM odontologos WHERE matricula = %s", matricula)
	row := s.db.QueryRow(query)

	err := row.Scan(&odontologoEncontrado.OdontologoId, &odontologoEncontrado.ApellidoOdontologo, &odontologoEncontrado.NombreOdontologo, &odontologoEncontrado.Matricula)
	if err != nil {
		return domains.Odontologo{}, err
	}

	return odontologoEncontrado, nil
}

func (s *OdontologoSqlStore) Create(o domains.Odontologo) (domains.Odontologo, error) {
	query := "INSERT INTO odontologos (apellido_odontologo, nombre_odontologo, matricula) VALUES (?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(o.ApellidoOdontologo, o.NombreOdontologo, o.Matricula)
	if err != nil {
		log.Fatal(err)
	}

	err = s.db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&o.OdontologoId)
	if err != nil {
		log.Fatal(err)
	}

	return o, nil
}

/* func (s *OdontologoSqlStore) Update(o domains.Odontologo) error {
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
} */

func (s *OdontologoSqlStore) Update(id int, odontologoInput domains.Odontologo) (domains.Odontologo, error) {

	_, err := s.db.Query("SELECT * FROM odontologos WHERE odontologo_id = ?", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("El Odontologo no existe en la base de datos:", err)
		} else {
			fmt.Println("Error al ejecutar la consulta:", err)
		}
		return domains.Odontologo{}, err
	}
	query := "UPDATE odontologos SET"
	if odontologoInput.OdontologoId > 0 {
		query += " odontologo_id = '" + fmt.Sprint(odontologoInput.OdontologoId) + "',"
	}
	if odontologoInput.NombreOdontologo != "" {
		query += " apellido_odontologo = '" + odontologoInput.NombreOdontologo + "',"
	}
	if odontologoInput.ApellidoOdontologo != "" {
		query += " nombre_odontologo = '" + odontologoInput.ApellidoOdontologo + "',"
	}
	if odontologoInput.Matricula != "" {
		query += " matricula = '" + odontologoInput.Matricula + "',"
	}

	query = query[0 : len(query)-1]
	query += " WHERE odontologo_id = ?"


	// actualizo el odontologo
	_, err = s.db.Exec(query, id)

	if err != nil {
		fmt.Println("Error al ejecutar la consulta:", err)
		return domains.Odontologo{}, err
	}

	// obtengo el recurso actualizado
	updatedOdontologo, _ := s.GetByID(id)

	return updatedOdontologo, nil
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

/* func (r *repository) Patch(c context.Context, id int, campos map[string]interface{}) (*domains.Odontologo, error) {
	odontologo, err := r.GetByID (id)
	if err != nil {
		return nil, err
	}
	for campo, valor := range campos {
		switch campo {
		case "apellido_odontologo":
			odontologo.ApellidoOdontologo = valor.(string)
		case "nombre_odontologo":
			odontologo.NombreOdontologo = valor.(string)
		case "matricula":
			odontologo.Matricula = valor.(string)
		}
	}
	var (QueryUpdateOdontologo = `UPDATE turnos_odontologos SET apellido_odontologo = ? ,nombre_odontologo = ? , matricula = ?
	FROM my_db.odontologos WHERE id = ?`)
	_, err = r.db.ExecContext(c, QueryUpdateOdontologo, odontologo.ApellidoOdontologo, odontologo.NombreOdontologo, odontologo.Matricula, odontologo.OdontologoId)
	if err != nil {
		return nil, err
	}
	return &odontologo, nil
} */

// **********************************************************************
// 2 IMPLEMENTACIÓN DE MÉTODOS DE LA INTERFACE DE PACIENTES

func (s *PacienteSqlStore) Read(id int) (domains.Paciente, error) {
	// Consulta para recuperar el paciente con el ID proporcionado
	query := "SELECT nombre_paciente, apellido_paciente, domicilio, dni, fecha_de_alta FROM pacientes WHERE paciente_id = ?"

	// Ejecutar la consulta y recuperar los datos
	var p domains.Paciente
	err := s.db.QueryRow(query, id).Scan(&p.NombrePaciente, &p.ApellidoPaciente, &p.DomicilioPaciente, &p.Dni, &p.FechaDeAlta)

	if err != nil {
		if err == sql.ErrNoRows {
			// El paciente con el ID proporcionado no fue encontrado
			return domains.Paciente{}, errors.New("paciente no encontrado")
		}
		// Ocurrió un error al ejecutar la consulta
		return domains.Paciente{}, err
	}
	p.PacienteID = id

	// Retornar los datos del paciente recuperado
	return p, nil
}

func (s *PacienteSqlStore) GetByDNI(dni string) (domains.Paciente, error) {
	// Consulta para recuperar el paciente con el ID proporcionado
	query := "SELECT * FROM pacientes WHERE dni = ?"

	// Ejecutar la consulta y recuperar los datos
	var p domains.Paciente
	err := s.db.QueryRow(query, dni).Scan(&p.PacienteID, &p.NombrePaciente, &p.ApellidoPaciente, &p.DomicilioPaciente, &p.Dni, &p.FechaDeAlta)
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

func (s *PacienteSqlStore) Create(p domains.Paciente) (domains.Paciente, error) {
	query := "INSERT INTO pacientes (nombre_paciente, apellido_paciente, domicilio, dni, fecha_de_alta) VALUES (?, ?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(p.NombrePaciente, p.ApellidoPaciente, p.DomicilioPaciente, p.Dni, p.FechaDeAlta)
	if err != nil {
		log.Fatal(err)
	}


	err = s.db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&p.PacienteID)
	if err != nil {
		log.Fatal(err)
	}

	return p, nil
}

func (s *PacienteSqlStore) Update(id int, p domains.Paciente) (domains.Paciente, error) {
	query := "UPDATE pacientes SET nombre_paciente = ?, apellido_paciente = ?, domicilio = ?, dni = ?, fecha_de_alta = ? WHERE paciente_id = ?;"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		return domains.Paciente{},err
	}

	res, err := stmt.Exec(p.NombrePaciente, p.ApellidoPaciente, p.DomicilioPaciente, p.Dni, p.FechaDeAlta, id)
	if err != nil {
		return domains.Paciente{},err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return domains.Paciente{},err
	}

	updatedPaciente, _ := s.Read(id)

	return  updatedPaciente, nil
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

func (s *TurnoSqlStore) ReadById(id int) (domains.Turno, error) {
	// Consulta para recuperar el turno con el ID proporcionado
	query := "SELECT * FROM turnos WHERE turno_id = ?"

	// Ejecutar la consulta y recuperar los datos
	var t domains.Turno
	err := s.db.QueryRow(query, id).Scan(&t.TurnosId, &t.FechaYHora, &t.Descripcion, &t.DentistaIDDentista, &t.PacienteIDPaciente)
	if err != nil {
		if err == sql.ErrNoRows {
			// El paciente con el ID proporcionado no fue encontrado
			return domains.Turno{}, errors.New("turno no encontrado")
		}
		// Ocurrió un error al ejecutar la consulta
		return domains.Turno{}, err
	}

	// Retornar los datos del turno recuperado
	return t, nil
}

// Recupera de la base de datos los turnos del paciente según DNI
func (s *TurnoSqlStore) GetByDNI(dni string) (domains.Turno, error) {
	// Consulta para recuperar el Id del vpaciente con el DNI proporcionado
	var vpaciente domains.Paciente
	query := "SELECT paciente_id FROM pacientes WHERE dni = ?"
	row := s.db.QueryRow(query, dni)
	err := row.Scan(&vpaciente.PacienteID)
	if err != nil {
		return domains.Turno{}, errors.New("No existen pacientes con ese DNI")
	}
	fmt.Println(vpaciente.PacienteID)
	// Ejecutar la consulta y recuperar los datos
	var t domains.Turno
	query = "SELECT * FROM turnos WHERE paciente_id_paciente = ? LIMIT 1"
	err = s.db.QueryRow(query, vpaciente.PacienteID).Scan(&t.TurnosId, &t.FechaYHora, &t.Descripcion, &t.DentistaIDDentista, &t.PacienteIDPaciente)
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

func (s *TurnoSqlStore) Create(p domains.Turno) (domains.Turno, error) {
	fmt.Println(p)
	query := "INSERT INTO turnos (fecha_y_hora, descripcion, dentista_id_dentista, paciente_id_paciente) VALUES (?, ?, ?, ?);"
	stmt, err := s.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}


	_, err = stmt.Exec(p.FechaYHora, p.Descripcion, p.DentistaIDDentista, p.PacienteIDPaciente)
	if err != nil {
		log.Fatal(err)
	}
	err = s.db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&p.TurnosId)
	if err != nil {
		log.Fatal(err)
	}

	return p, nil
}

func (s *TurnoSqlStore) Update(id int, p domains.Turno) (domains.Turno, error) {

	query := "UPDATE turnos SET fecha_y_hora = ?, descripcion = ?, dentista_id_dentista = ?, paciente_id_paciente = ? WHERE turno_id = ?;"
	stmt, err := s.db.Prepare(query)

	fmt.Println(err)
	if err != nil {
		return domains.Turno{}, err
	}
	res, err := stmt.Exec(p.FechaYHora, p.Descripcion, p.DentistaIDDentista, p.PacienteIDPaciente,id)
	if err != nil {
		return domains.Turno{}, err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return domains.Turno{},err
	}
	p.TurnosId =id
	return p, nil
}

func (s *TurnoSqlStore) Delete(id int) error {
	query := "DELETE FROM turnos WHERE turno_id = ?;"
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
