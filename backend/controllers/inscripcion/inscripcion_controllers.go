package inscripcion

import (
	"backend/dto"
	serviceInscr "backend/services/inscripcion"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func InsertInscr(c *gin.Context) {
	// Declara una variable para almacenar los datos de la inscripción
	var inscrDto dto.InscripcionDto

	// Intenta parsear el cuerpo de la solicitud JSON en la variable inscrDto
	err := c.BindJSON(&inscrDto)
	if err != nil {
		// Registra el error y envía una respuesta 400 Bad Request con el mensaje de error
		log.Error("Error al parsear JSON: ", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Llama al servicio para insertar la inscripción
	inscrDto, er := serviceInscr.InscripcionService.InsertInscr(inscrDto)
	if er != nil {
		// Si ocurre un error en el servicio, envía una respuesta con el estado y el error
		log.Error("Error al insertar inscripción: ", er.Error())
		c.JSON(http.StatusInternalServerError, er.Error())
		return
	}

	// Si la inserción fue exitosa, envía una respuesta 201 Created con los datos de la inscripción
	c.JSON(http.StatusCreated, inscrDto)
}

// GetInscripcion maneja la obtención de todas las inscripciones
func GetInscripcion(c *gin.Context) {
	inscrsDto, err := serviceInscr.InscripcionService.GetInscripcion()
	if err != nil {
		log.Error("Error al obtener inscripciones: ", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Envía una respuesta 200 OK con los datos de las inscripciones
	c.JSON(http.StatusOK, inscrsDto)
}

// GetCourseByUserId maneja la obtención de cursos por ID de usuario
func GetCourseByUserId(c *gin.Context) {
	userId := c.Param("id_user")

	var inscrDto dto.InscripcionDto
	inscrDto.Id_user, _ = strconv.Atoi(userId)

	inscrsDto, err := serviceInscr.InscripcionService.GetCourseByUserId(inscrDto)
	if err != nil {
		log.Error("Error al obtener cursos por ID de usuario: ", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Envía una respuesta 200 OK con los datos de los cursos del usuario
	c.JSON(http.StatusOK, inscrsDto)
}

func InsertValoracion(c *gin.Context) {
	var valoracionDto dto.ValoracionDto
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&valoracionDto)
	if err != nil {
		log.Error("Error al parsear JSON: ", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	valoracionResponse, er := serviceInscr.InscripcionService.InsertValoracion(id, valoracionDto)
	if er != nil {
		log.Error("Error al agregar comentario y valoración: ", er.Error())
		c.JSON(http.StatusInternalServerError, er.Error())
		return
	}

	c.JSON(http.StatusCreated, valoracionResponse)
}

func InsertArchivo(c *gin.Context) {
	var archivoDto dto.ArchivoDto
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error("Error al convertir id: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	file, err := c.FormFile("archivo")
	if err != nil {
		log.Error("Error al obtener archivo: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener archivo"})
		return
	}

	fileContent, err := file.Open()
	if err != nil {
		log.Error("Error al abrir archivo: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al abrir archivo"})
		return
	}
	defer fileContent.Close()

	archivoBytes := make([]byte, file.Size)
	_, err = fileContent.Read(archivoBytes)
	if err != nil {
		log.Error("Error al leer archivo: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer archivo"})
		return
	}

	archivoDto.Archivo = archivoBytes
	if err := c.ShouldBind(&archivoDto); err != nil {
		log.Error("Error al parsear datos adicionales: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al parsear datos adicionales"})
		return
	}

	archivoResponse, er := serviceInscr.InscripcionService.InsertArchivo(id, archivoDto)
	if er != nil {
		log.Error("Error al subir archivo: ", er.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al subir archivo"})
		return
	}

	c.JSON(http.StatusCreated, archivoResponse)
}

func GetComentByCourseId(c *gin.Context) {
	courseId := c.Param("id_course")

	var valDto dto.ValoracionDto
	valDto.Id_course, _ = strconv.Atoi(courseId)

	valsDto, err := serviceInscr.InscripcionService.GetComentByCourseId(valDto)
	if err != nil {
		log.Error("Error al obtener cursos por ID de usuario: ", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Envía una respuesta 200 OK con los datos de los cursos del usuario
	c.JSON(http.StatusOK, valsDto)
}
