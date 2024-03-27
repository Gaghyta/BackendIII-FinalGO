package handler

import (
	"errors"
	"strconv"

	"github.com/Gaghyta/BackendIIIFinalGO/internal/domains"
	paciente "github.com/Gaghyta/BackendIIIFinalGO/internal/pacientes"
	"github.com/Gaghyta/BackendIIIFinalGO/pkg/web"
	"github.com/gin-gonic/gin"
)

type pacienteHandler struct {
	ps paciente.Service
}

func NewPacienteHandler(p paciente.Service) *pacienteHandler {
	return &pacienteHandler{
		ps: p,
	}
}

// GetById obtiene un paciente por id
func (h *pacienteHandler) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("paciente_id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("este Id es inválido"))
			return
		}
		paciente, err := h.ps.GetByID(id)
		if err != nil {
			web.Failure(ctx, 404, errors.New("el paciente que estás buscando no ha sido encontrado"))
		}
		web.Success(ctx, 200, paciente)
	}
}

// validateEmptys valida que los campos no estén vacíos
func validateEmptyPaciente(paciente *domains.Paciente) (bool, error) {
	if paciente.NombrePaciente == "" ||
		paciente.ApellidoPaciente == "" ||
		paciente.Dni == "" ||
		paciente.FechaDeAlta == "" {
		return false, errors.New("los campos no pueden estar vacíos")
	}
	return true, nil
}

func (h *pacienteHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paciente domains.Paciente
		err := ctx.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failure(ctx, 400, errors.New("JSON inválido"))
			return
		}
		valido, err := validateEmptyPaciente(&paciente)
		if !valido {
			web.Failure(ctx, 400, err)
			return
		}
		p, err := h.ps.Create(paciente)
		if err != nil {
			web.Failure(ctx, 400, err)
			return
		}
		web.Success(ctx, 201, p)
	}
}

// put modifica un paciente
func (h *pacienteHandler) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Obtener el Id del paciente de la URL
		idParam := ctx.Param("paciente_id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("id inválido"))
			return
		}
		// Obtener los datos actualizados del paciente del cuerpo de la solicitud
		var paciente domains.Paciente
		err = ctx.ShouldBindJSON(&paciente)
		if err != nil {
			web.Failure(ctx, 400, errors.New("JSON inválido"))
			return
		}
		// validar los datos del paciente actualizado
		valido, err := validateEmptyPaciente(&paciente)
		if !valido {
			web.Failure(ctx, 400, err)
			return
		}
		// Llamar al servicio para realizar la actualización en la base de datos
		pacienteActualizado, err := h.ps.Update(id, paciente)
		if err != nil {
			web.Failure(ctx, 400, err)
			return
		}

		web.Success(ctx, 201, pacienteActualizado)
	}
}

func (h *pacienteHandler) DeleteByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Obtener el Id del paciente de la URL
		idParam := ctx.Param("paciente_id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, 400, errors.New("ID inválido"))
			return
		}

		// llamar al servicio para eliminar un paciente
		err = h.ps.Delete(id)
		if err != nil {
			web.Failure(ctx, 500, err) // Cambiar el código de estado si es necesario
			return
		}

		web.Success(ctx, 201, "Paciente eliminado exitosamente")
	}
}

// Patch actualiza uno o varios campos
func (h *pacienteHandler) Patch() gin.HandlerFunc {
	type Request struct {
		NombrePaciente    string `json:"nombre_paciente,omitempty"`
		ApellidoPaciente  string `json:"apellido_paciente,omitempty"`
		DomicilioPaciente string `json:"domicilio_paciente,omitempty"`
		Dni               string `json:"dni,omitempty"`
		FechaDeAlta       string `json:"fecha_de_alta,omitempty"`
	}
	return func(ctx *gin.Context) {
		var r Request
		idParam := ctx.Param("paciente_id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		if err := ctx.ShouldBindJSON(&r); err != nil {
			ctx.JSON(400, gin.H{"error": "invalid json"})
			return
		}
		update := domains.Paciente{
			NombrePaciente:    r.NombrePaciente,
			ApellidoPaciente:  r.ApellidoPaciente,
			DomicilioPaciente: r.ApellidoPaciente,
			Dni:               r.Dni,
			FechaDeAlta:       r.FechaDeAlta,
		}
		p, err := h.ps.Patch(id, update)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, p)
	}
}
