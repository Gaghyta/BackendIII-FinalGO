package handler

import (
	"errors"
	"strconv"

	"github.com/Gaghyta/BackendIIIFinalGO/internal/domains"
	"github.com/Gaghyta/BackendIIIFinalGO/internal/turnos"

	"github.com/Gaghyta/BackendIIIFinalGO/pkg/web"
	"github.com/gin-gonic/gin"
)

type turnoHandler struct {
	ts turnos.Service
}

func NewTurnoHandler(t turnos.Service) *turnoHandler {
	return &turnoHandler{
		ts: t,
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
		t, err := h.ts.Patch(id, update)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, t)
	}
}
