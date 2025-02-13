package cursos

import (
	"context"
	cursosDomain "cursos/domain_cursos"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Service interface {
	GetCourseByID(ctx context.Context, id string) (cursosDomain.CourseDto, error)
	GetCourses(ctx context.Context) (cursosDomain.CoursesDto, error)
	Create(ctx context.Context, curso cursosDomain.CourseDto) (string, error)
	Update(ctx context.Context, curso cursosDomain.CourseDto) error
}

type Controller struct {
	service Service
}

func NewController(service Service) Controller {
	return Controller{
		service: service,
	}
}

func (controller Controller) GetCourseByID(ctx *gin.Context) {
	// Validate ID param
	cursoID := strings.TrimSpace(ctx.Param("id"))

	// Get course by ID using the service
	curso, err := controller.service.GetCourseByID(ctx.Request.Context(), cursoID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("error getting course: %s", err.Error()),
		})
		return
	}

	// Send response
	ctx.JSON(http.StatusOK, curso)
}

func (controller Controller) Create(ctx *gin.Context) {
	// Parse course
	var curso cursosDomain.CourseDto
	if err := ctx.ShouldBindJSON(&curso); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("invalid request: %s", err.Error()),
		})
		return
	}

	// Create course
	id, err := controller.service.Create(ctx.Request.Context(), curso)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error creating course: %s", err.Error()),
		})
		return
	}

	// Send ID
	ctx.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func (controller Controller) Update(ctx *gin.Context) {
	// Validate ID param
	id := strings.TrimSpace(ctx.Param("id"))

	// Parse course
	var curso cursosDomain.CourseDto
	if err := ctx.ShouldBindJSON(&curso); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("invalid request: %s", err.Error()),
		})
		return
	}

	// Set the ID from the URL to the course object
	curso.Course_id = id

	// Update course
	if err := controller.service.Update(ctx.Request.Context(), curso); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error updating course: %s", err.Error()),
		})
		return
	}

	// Send response
	ctx.JSON(http.StatusOK, gin.H{
		"message": id,
	})
}

func (controller Controller) GetCourses(c *gin.Context) {
	// Llamar al servicio pasando el contexto
	courses, err := controller.service.GetCourses(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error getting courses: %s", err.Error()),
		})
		return
	}

	// Enviar respuesta
	c.JSON(http.StatusOK, courses)
}
