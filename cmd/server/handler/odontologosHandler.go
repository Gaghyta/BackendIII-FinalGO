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

// NewOdontologoHandler crea un nuevo controller de odontologos
func NewOdontologotHandler(s odontologos.Service) *odontologoHandler {
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

// validateEmptyOdontologo valida que los campos no esten vacios
func validateEmptyOdontologo(odontologo *domains.Odontologo) (bool, error) {
	if odontologo.NombreOdontologo == "" || odontologo.ApellidoOdontologo == "" || odontologo.Matricula == "" {
		return false, errors.New("fields can't be empty")
	}

	return true, nil
}

/*  validateExpiration valida que la fecha de expiracion sea valida
func validateExpiration(exp string) (bool, error) {
	dates := strings.Split(exp, "/")
	list := []int{}
	if len(dates) != 3 {
		return false, errors.New("invalid expiration date, must be in format: dd/mm/yyyy")
	}
	for value := range dates {
		number, err := strconv.Atoi(dates[value])
		if err != nil {
			return false, errors.New("invalid expiration date, must be numbers")
		}
		list = append(list, number)
	}
	condition := (list[0] < 1 || list[0] > 31) && (list[1] < 1 || list[1] > 12) && (list[2] < 1 || list[2] > 9999)
	if condition {
		return false, errors.New("invalid expiration date, date must be between 1 and 31/12/9999")
	}
	return true, nil
} */

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
		/* valid, err = validateExpiration(product.Expiration)
		if !valid {
			web.Failure(c, 400, err)
			return
		} */
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
		valid, err := validateEmptyOdontologo(&odontologo)
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
