package handler

import (
	"errors"
	"fmt"

	"strconv"

	"github.com/Gaghyta/BackendIIIFinalGO/internal/domains"
	"github.com/Gaghyta/BackendIIIFinalGO/internal/odontologos"

	"github.com/Gaghyta/BackendIIIFinalGO/pkg/web"
	"github.com/gin-gonic/gin"
)

type odontologoHandler struct {
	os odontologos.Service
}

// NewProductHandler crea un nuevo controller de productos
func NewOdontologoHandler(s odontologos.Service) *odontologoHandler {
	return &odontologoHandler{
		os: s,
	}
}

// Get obtiene un odontólogo por id
func (h *odontologoHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(1)
		idParam := c.Param("odontologo_id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("este id es inválido"))
			return
		}
		odontologo, err := h.os.GetByID(id)
		fmt.Println(2)
		if err != nil {
			web.Failure(c, 404, errors.New("el odontólogo que estás buscando no ha sido encontrado"))
			return
		}
		web.Success(c, 200, odontologo)
	}
}

// validateEmptyOdontologo valida que los campos no esten vacios
func validateEmptyOdontologo(odontologo *domains.Odontologo) (bool, error) {
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

		valid, err := validateEmptyOdontologo(&odontologo)
		if !valid {
			web.Failure(c, 400, err)
			return
		}

		o, err := h.os.Create(odontologo)
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
		idParam := c.Param("odontologo_id")
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
		valid, err := validateEmptyOdontologo(&odontologo)
		if !valid {
			web.Failure(c, 400, err)
			return
		}

		// Llamar al servicio para realizar la actualización en la base de datos
		odontologoActualizado, err := h.os.Update(id, odontologo)
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
		idParam := c.Param("odontologo_id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("ID inválido"))
			return
		}

		// Llamar al servicio para eliminar el odontólogo
		err = h.os.Delete(id)
		if err != nil {
			web.Failure(c, 500, err) // Cambiar el código de estado si es necesario
			return
		}

		web.Success(c, 201, "Odontólogo eliminado exitosamente")
	}
}

func (h *odontologoHandler) Patch() gin.HandlerFunc {

	// Obtener los datos del cuerpo de la solicitud
	type Request struct {
		NombreOdontologo   string `json:"nombre_odontologo,omitempty"`
		ApellidoOdontologo string `json:"apellido_odontologo,omitempty"`
		Matricula          string `json:"matricula,omitempty"`
	}

	return func(c *gin.Context) {
		var r Request
		idParam := c.Param("odontologo_id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid id"})
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			c.JSON(400, gin.H{"error": "invalid json"})
			return
		}

		update := domains.Odontologo{
			ApellidoOdontologo: r.ApellidoOdontologo,
			NombreOdontologo:   r.NombreOdontologo,
			Matricula:          r.Matricula,
		}
		o, err := h.os.Update(id, update)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, o)
	}
}
