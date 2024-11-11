package controllers_inscripcion

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"inscripciones/dao_inscripcion"
	"inscripciones/domain_inscripcion"
	"net/http"
	"strconv"
)

type InscripcionServiceInterface interface {
	InsertInscripcion(inscripcionDto domain_inscripcion.InscripcionDto) (domain_inscripcion.InscripcionDto, error)
	GetInscripcionByUserID(userID int) ([]dao_inscripcion.Inscripcion, error)
	GetInscripcionByCourseID(courseID int) ([]dao_inscripcion.Inscripcion, error)
}

type InscripcionController struct {
	service InscripcionServiceInterface
}

func NewController(service InscripcionServiceInterface) *InscripcionController {
	return &InscripcionController{service: service}
}

func (controller *InscripcionController) InsertInscripcion(c *gin.Context) {
	var inscripcionDto domain_inscripcion.InscripcionDto
	err := c.BindJSON(&inscripcionDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("inscripcionDto de controller: ", inscripcionDto)
	inscripcionDto, er := controller.service.InsertInscripcion(inscripcionDto)

	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}

	c.JSON(http.StatusCreated, inscripcionDto)
}

func (controller *InscripcionController) GetInscripcionByUserID(c *gin.Context) {
	userIDParam := c.Param("userID")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	inscripciones, err := controller.service.GetInscripcionByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, inscripciones)
}

func (controller *InscripcionController) GetInscripcionByCourseID(c *gin.Context) {
	courseIDParam := c.Param("courseID")
	courseID, err := strconv.Atoi(courseIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	inscripciones, err := controller.service.GetInscripcionByCourseID(courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, inscripciones)
}
