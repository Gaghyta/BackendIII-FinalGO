package handler

import (
	"errors"
	"strconv"

	"github.com/Gaghyta/BackendIIIFinalGO/internal/domains"
	"github.com/Gaghyta/BackendIIIFinalGO/internal/odontologos"

	"github.com/Gaghyta/BackendIIIFinalGO/pkg/web"
	"github.com/gin-gonic/gin"
)

type odontologoHandler struct {
	s odontologos.Service
}

// NewProductHandler crea un nuevo controller de productos
func NewOdontologoHandler(s odontologos.Service) *odontologoHandler {
	return &odontologoHandler{
		s: s,
	}
}

// Get obtiene un odontólogo por id
func (h *odontologoHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("este id es inválido"))
			return
		}
		odontologo, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("el odontólogo que estás buscando no ha sido encontrado"))
			return
		}
		web.Success(c, 200, odontologo)
	}
}

// validateEmptys valida que los campos no esten vacios
func validateEmptys(odontologo *domains.Odontologo) (bool, error) {
	if odontologo.NombreOdontologo == "" || odontologo.ApellidoOdontologo == "" || odontologo.Matricula == "" {
		return false, errors.New("fields can't be empty")
	}

	return true, nil
}

// Post crea un nuevo odontólogo
func (h *odontologoHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var odontologo domains.Odontologo
		err := c.ShouldBindJSON(&odontologo)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptys(&odontologo)
		if !valid {
			web.Failure(c, 400, err)
			return
		}

		o, err := h.s.Create(odontologo)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, o)
	}
}

// Put modifica un odontólogo
func (h *odontologoHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Obtener el ID del odontólogo de la URL
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid ID"))
			return
		}

		// Obtener los datos actualizados del odontólogo del cuerpo de la solicitud
		var odontologo domains.Odontologo
		err = c.ShouldBindJSON(&odontologo)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid JSON"))
			return
		}

		// Validar los datos del odontólogo actualizado
		valid, err := validateEmptys(&odontologo)
		if !valid {
			web.Failure(c, 400, err)
			return
		}

		// Llamar al servicio para realizar la actualización en la base de datos
		odontologoActualizado, err := h.s.Update(id, odontologo)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}

		web.Success(c, 201, odontologoActualizado)

	}
}

func (h *odontologoHandler) DeleteByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el ID del odontólogo de la URL
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID inválido"))
			return
		}

		// Llamar al servicio para eliminar el odontólogo
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(c, 500, err) // Cambiar el código de estado si es necesario
			return
		}

		web.Success(c, 201, "Odontólogo eliminado exitosamente")
	}
}

/* func (h *odontologoHandler) Patch() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener los datos del cuerpo de la solicitud
		var patchData struct {
			NuevaMatricula string `json:"nueva_matricula"`
		}
		if err := c.ShouldBindJSON(&patchData); err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("datos JSON inválidos"))
			return
		}

		// Obtener la matrícula actual y la nueva matrícula del cuerpo de la solicitud
		matricula := c.Param("matricula")
		nuevaMatricula := patchData.NuevaMatricula

		// Llamar al servicio para actualizar la matrícula
		odontologo, err := h.s.Patch(matricula, nuevaMatricula)
		if err != nil {
			web.Failure(c, http.StatusInternalServerError, err)
			return
		}

		// Si no hay errores, responder con el objeto odontólogo actualizado
		web.Success(c, http.StatusOK, odontologo)
	}
} */
