package cursos

import (
	"backend/dto"
	servicios "backend/services/cursos"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func GetCourseById(c *gin.Context) {
	var courseDto dto.CourseDto
	id, _ := strconv.Atoi(c.Param("course_id"))
	courseDto, err := servicios.CourseService.GetCourseById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, courseDto)
}

func GetCourses(c *gin.Context) {
	var coursesDto dto.CoursesDto
	var err error

	coursesDto, err = servicios.CourseService.GetCourses()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, coursesDto)
}

func GetCoursesByName(c *gin.Context) {
	var coursesDto dto.CoursesDto
	query := c.Param("name")
	coursesDto, err := servicios.CourseService.GetCoursesByName(query)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if len(coursesDto) == 0 {
		c.JSON(http.StatusOK, []dto.CourseDto{})
		return
	}
	c.JSON(http.StatusOK, coursesDto)
}

func InsertCourse(c *gin.Context) {
	var course dto.CourseDto
	err := c.BindJSON(&course)
	if err != nil {
		log.Println("Error al parsear el JSON: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error al parsear el JSON"})
		return
	}

	createdCourse, er := servicios.CourseService.InsertCourse(course)
	if er != nil {
		log.Println("Error al crear curso: ", er.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error al crear curso"})
		return
	}

	c.JSON(http.StatusCreated, createdCourse)
}

func EditCourse(c *gin.Context) {
	var course dto.CourseDto
	id, _ := strconv.Atoi(c.Param("course_id"))

	err := c.BindJSON(&course)
	if err != nil {
		log.Error("Error al parsear el JSON: ", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	updatedCourse, er := servicios.CourseService.EditCourse(id, course)
	if er != nil {
		log.Error("Error al editar curso: ", er.Error())
		c.JSON(http.StatusInternalServerError, er.Error())
		return
	}

	c.JSON(http.StatusOK, updatedCourse)
}

func DeleteCourse(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("course_id"))

	courseDto, err := servicios.CourseService.DeleteCourse(id)
	if err != nil {
		log.Error("Error al eliminar curso: ", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, courseDto)
}
