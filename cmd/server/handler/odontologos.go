package handler

import (
	"errors"
	"strconv"

	"github.com/Gaghyta/BackendIIIFinalGO/internal/odontologos"
	"github.com/bootcamp-go/consignas-go-db.git/pkg/web"
	"github.com/gin-gonic/gin"
)

type odontologotHandler struct {
	s odontologos.Service
}

// NewProductHandler crea un nuevo controller de productos
func NewProductHandler(s odontologos.Service) *odontologotHandler {
	return &odontologotHandler{
		s: s,
	}
}

// Get obtiene un odontólogo por id
func (h *odontologotHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("Este id es inválido"))
			return
		}
		odontologo, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("El odontólogo que estás buscando no ha sido encontrado"))
			return
		}
		web.Success(c, 200, odontologo)
	}
}

// validateEmptys valida que los campos no esten vacios
func validateEmptys(odontologo *odontologos.Odontologos) (bool, error) {
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
func (h *odontologotHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var odontologo odontologos.Odontologos
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