package handler

import (
	"errors"
	"fmt"
	"net/http"
	"os"
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

// GetByID godoc
// @Summary Get odontologo
// @Description obtiene un odontólogo por id
// @Tags domain.odontologo
// @Param id path int true "id del odontologo"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /odontologos/:id [get]

func (h *odontologoHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {

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
		return false, errors.New("los campos no pueden estar vacíos")
	}

	return true, nil
}

// Post godoc
// @Summary Post odontologo
// @Description Create a new odontologo
// @Tags domains.odontologo
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /odontologos [post]

func (h *odontologoHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token no encontrado"})
			return
		}
		if token != os.Getenv("TOKEN") {
			c.JSON(400, gin.H{"error": "token inválido"})
			return
		}

		var odontologo domains.Odontologo
		err := c.ShouldBindJSON(&odontologo)
		if err != nil {
			web.Failure(c, 400, errors.New("json inválido"))
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

// Put godoc
// @Summary Put odontologo
// @Description modifica un odontólogo por su id
// @Tags domain.odontologo
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /odontologos/:id [put]

func (h *odontologoHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token no encontrado"})
			return
		}
		if token != os.Getenv("TOKEN") {
			c.JSON(400, gin.H{"error": "token inválido"})
			return
		}
		// Obtener el ID del odontólogo de la URL
		idParam := c.Param("odontologo_id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("id inválido"))
			return
		}

		// Obtener los datos actualizados del odontólogo del cuerpo de la solicitud
		var odontologo domains.Odontologo
		err = c.ShouldBindJSON(&odontologo)
		if err != nil {
			web.Failure(c, 400, errors.New("json inválido"))
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

// DeleteById godoc
// @Summary DeletById odontologo
// @Description Elima un odontólogo por su id
// @Tags domain.odontologo
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /odontologos/:id [delete]
func (h *odontologoHandler) DeleteByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token no encontrado"})
			return
		}
		if token != os.Getenv("TOKEN") {
			c.JSON(400, gin.H{"error": "token inválido"})
			return
		}
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

// Patch godoc
// @Summary Patch odontologo
// @Description Mofica datos en los campos de odontologo mediante su id
// @Tags domain.odontologo
// @Param id path int true "id del odontologo"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /odontologos/:id [patch]

func (h *odontologoHandler) Patch() gin.HandlerFunc {

	// Obtener los datos del cuerpo de la solicitud
	type Request struct {
		NombreOdontologo   string `json:"nombre_odontologo,omitempty"`
		ApellidoOdontologo string `json:"apellido_odontologo,omitempty"`
		Matricula          string `json:"matricula,omitempty"`
	}

	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token no encontrado"})
			return
		}
		if token != os.Getenv("TOKEN") {
			c.JSON(400, gin.H{"error": "token inválido"})
			return
		}

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
