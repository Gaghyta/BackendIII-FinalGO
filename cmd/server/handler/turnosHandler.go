package handler

import (
	"errors"
	"strconv"

	"github.com/Gaghyta/BackendIIIFinalGO/internal/domains"
	"github.com/Gaghyta/BackendIIIFinalGO/internal/odontologos"
	paciente "github.com/Gaghyta/BackendIIIFinalGO/internal/pacientes"

	//"github.com/Gaghyta/BackendIIIFinalGO/internal/pacientes"
	"github.com/Gaghyta/BackendIIIFinalGO/internal/turnos"

	//Store "github.com/Gaghyta/BackendIIIFinalGO/pkg/store"
	"github.com/Gaghyta/BackendIIIFinalGO/pkg/web"
	"github.com/gin-gonic/gin"
)

type turnoHandler struct {
	ts turnos.Service
	ps paciente.Service
	os odontologos.Service
}

func NewTurnoHandler(t turnos.Service, p paciente.Service, o odontologos.Service) *turnoHandler {
	return &turnoHandler{
		ts: t,
		ps: p,
		os: o,
	}
}
// GetById obtiene un turno por id
func (h *turnoHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("turnos_id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Este id es inválido"))
			return
		}
		turno, err := h.ts.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("El turno que estás buscando no ha sido encontrado"))
		}
		web.Success(c, 200, turno)
	}
}

// GetByDNI obtiene un turno por dni
func (h *turnoHandler) GetByDNI() gin.HandlerFunc {
	return func(c *gin.Context) {
		dniParam := c.Query("dni")
		if dniParam == "" {
			web.Failure(c, 400, errors.New("Este dni es inválido"))
			return
		}
		turno, err := h.ts.GetByDNI(dniParam)
		if err != nil {
			web.Failure(c, 404, errors.New("El turno que estás buscando no ha sido encontrado"))
		}
		web.Success(c, 200, turno)
	}
}

// validateEmptys valida que los campos no estén vacíos
func validateEmptyTurno(turno *domains.Turno) (bool, error) {
	if turno.FechaYHora == "" || turno.Descripcion == "" {
		//CONVERTIR  ||
		//turno.DentistaIDDentista == "" ||
		//turno.PacienteIDPaciente == "" {

		return false, errors.New("Los campos no pueden estar vacíos")
	}
	return true, nil
}

func (h *turnoHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var turno domains.Turno
		err := c.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(c, 400, errors.New("JSON inválido"))
			return
		}
		valido, err := validateEmptyTurno(&turno)
		if !valido {
			web.Failure(c, 400, err)
			return
		}
		t, err := h.ts.Create(turno)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, t)
	}
}

func (h *turnoHandler) PostWithDniAndMatricula() gin.HandlerFunc {
	return func(c *gin.Context) {
		//genero la estructura que espero recibir en el POST
		type Turno_dni_mat struct {
			TurnosId          int    `json:"turnos_id"`
			FechaYHora        string `json:"fecha_y_hora"`
			Descripcion       string `json:"descripcion"`
			MatriculaDentista string `json:"matricula_dentista"`
			DNIPaciente       string `json:"dni_paciente"`
		}
		// Creo variable asociada a la estructura
		var turno_dni_mat Turno_dni_mat

		// Cargo la estructura con los datos del POST
		err := c.ShouldBindJSON(&turno_dni_mat)
		if err != nil {
			web.Failure(c, 400, errors.New("JSON inválido"))
			return
		}

		// Me conecto con Service de pacientes, y recupero el idPaciente
		paciente, err := h.ps.GetByDNI(turno_dni_mat.DNIPaciente)
		if err != nil {
			web.Failure(c, 400, errors.New("No existe ningún paciente con ese DNI"))
			return
		}

		// Me conecto con Service de odontologos, y recupero el idOdontologo
		odontologo, err := h.os.GetByMatricula(turno_dni_mat.MatriculaDentista)
		if err != nil {
			web.Failure(c, 400, errors.New("No existe ningún odontólogo con esa matrícula"))
			return
		}

		// Creo variable con estructura Turno
		var turno domains.Turno

		// Cargo esta variable con los datos recién obtenidos
		turno.FechaYHora = turno_dni_mat.FechaYHora
		turno.Descripcion = turno_dni_mat.Descripcion
		turno.DentistaIDDentista = odontologo.OdontologoId
		turno.PacienteIDPaciente = paciente.PacienteID

		// Valido todos los datos
		valido, err := validateEmptyTurno(&turno)
		if !valido {
			web.Failure(c, 400, err)
			return
		}

		// Creo el turno
		t, err := h.ts.Create(turno)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}

		web.Success(c, 201, t)
	}
}

// put modifica un turno
func (h *turnoHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el Id del turno de la URL
		idParam := c.Param("turnos_id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Id inválido"))
			return
		}
		// Obtener los datos actualizados del turno del cuerpo de la solicitud
		var turno domains.Turno
		err = c.ShouldBindJSON(&turno)
		if err != nil {
			web.Failure(c, 400, errors.New("JSON inválido"))
			return
		}
		// validar los datos del turno actualizado
		valido, err := validateEmptyTurno(&turno)
		if !valido {
			web.Failure(c, 400, err)
			return
		}
		// Llamar al servicio para realizar la actualización en la base de datos
		odontologoActualizado, err := h.ts.Update(id, turno)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}

		web.Success(c, 201, odontologoActualizado)
	}
}


func (h *turnoHandler) DeleteByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el Id del turno de la URL
		idParam := c.Param("turnos_id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID inválido"))
			return
		}

		// llamar al servicio para eliminar un turno
		err = h.ts.Delete(id)
		if err != nil {
			web.Failure(c, 500, err) // Cambiar el código de estado si es necesario
			return
		}

		web.Success(c, 201, "Turno eliminado exitosamente")
	}
}

// Patch actualiza uno o varios campos
func (h *turnoHandler) Patch() gin.HandlerFunc {
	type Request struct {
		FechaYHora         string `json:"fecha_y_hora,omitempty"`
		Descripcion        string `json:"descripcion,omitempty"`
		DentistaIDDentista string `json:"dentista_id_dentista,omitempty"`
		PacienteIDPaciente string `json:"paciente_id_paciente,omitempty"`
	}
	return func(c *gin.Context) {
		var r Request
		idParam := c.Param("turnos_id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			c.JSON(400, gin.H{"error": "invalid json"})
			return
		}
		update := domains.Turno{
			FechaYHora:  r.FechaYHora,
			Descripcion: r.Descripcion,
			//CONVERTIR VER
			//DentistaIDDentista: r.DentistaIDDentista,
			//PacienteIDPaciente: r.PacienteIDPaciente,
		}
		t, err := h.ts.Update(id, update)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, t)
	}
}
