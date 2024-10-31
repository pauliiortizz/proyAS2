package hotels

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
	Create(ctx context.Context, curso cursosDomain.CourseDto) (string, error)
	//Update(ctx context.Context, curso cursosDomain.CourseDto) error
	//Delete(ctx context.Context, id string) error
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

	// Get hotel by ID using the service
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
	// Parse hotel
	var curso cursosDomain.CourseDto
	if err := ctx.ShouldBindJSON(&curso); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("invalid request: %s", err.Error()),
		})
		return
	}

	// Create hotel
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

/*
func (controller Controller) Update(ctx *gin.Context) {
	// Validate ID param
	id := strings.TrimSpace(ctx.Param("id"))

	// Parse hotel
	var curso cursosDomain.CourseDto
	if err := ctx.ShouldBindJSON(&curso); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("invalid request: %s", err.Error()),
		})
		return
	}

	// Set the ID from the URL to the hotel object
	curso.Course_id = id

	// Update hotel
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
}*/

/*
func (controller Controller) Delete(ctx *gin.Context) {
	// Validate ID param
	id := strings.TrimSpace(ctx.Param("id"))

	// Delete hotel
	if err := controller.service.Delete(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error deleting hotel: %s", err.Error()),
		})
		return
	}

	// Send response
	ctx.JSON(http.StatusOK, gin.H{
		"message": id,
	})
}*/
