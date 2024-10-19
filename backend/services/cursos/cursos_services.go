package cursos

import (
	courseClient "backend/clients/cursos"
	"backend/dto"
	errores "backend/extras"
	"backend/model"
	log "github.com/sirupsen/logrus"
)

type courseService struct {
	coursesClient courseClient.CoursesClientInterface
}

type coursesServiceInterface interface {
	GetCourseById(id int) (dto.CourseDto, errores.ApiError)
	GetCourses() (dto.CoursesDto, errores.ApiError)
	GetCoursesByName(query string) (dto.CoursesDto, errores.ApiError)
	InsertCourse(creacion dto.CourseDto) (dto.CourseDto, errores.ApiError)
	EditCourse(id int, course dto.CourseDto) (dto.CourseDto, errores.ApiError)
	DeleteCourse(id int) (dto.CourseDto, errores.ApiError)
}

var (
	CourseService coursesServiceInterface
)

func initCourseService(coursesClient courseClient.CoursesClientInterface) coursesServiceInterface {
	course := new(courseService)
	course.coursesClient = coursesClient
	return course
}

func init() {
	CourseService = initCourseService(courseClient.CoursesClient)
}

func (s *courseService) GetCourseById(id int) (dto.CourseDto, errores.ApiError) {

	var course model.Courses = s.coursesClient.GetCourseById(id)
	var coursesDto dto.CourseDto

	if course.Course_id < 0 {
		return coursesDto, errores.NewBadRequestApiError("course not found")
	}

	coursesDto.Course_id = course.Course_id
	coursesDto.Nombre = course.Nombre
	coursesDto.Profesor_id = course.Profesor_id
	coursesDto.Categoria = course.Categoria
	coursesDto.Descripcion = course.Descripcion
	coursesDto.Valoracion = course.Valoracion
	coursesDto.Duracion = course.Duracion
	coursesDto.Requisitos = course.Requisitos
	coursesDto.Url_image = course.Url_image
	coursesDto.Fecha_inicio = course.Fecha_inicio
	return coursesDto, nil
}

func (s *courseService) GetCourses() (dto.CoursesDto, errores.ApiError) {

	var courses model.Coursess = s.coursesClient.GetCourses()
	var coursesDto dto.CoursesDto

	for _, course := range courses {
		var courseDto dto.CourseDto
		courseDto.Course_id = course.Course_id
		courseDto.Nombre = course.Nombre
		courseDto.Profesor_id = course.Profesor_id
		courseDto.Categoria = course.Categoria
		courseDto.Descripcion = course.Descripcion
		courseDto.Valoracion = course.Valoracion
		courseDto.Duracion = course.Duracion
		courseDto.Requisitos = course.Requisitos
		courseDto.Url_image = course.Url_image
		courseDto.Fecha_inicio = course.Fecha_inicio
		coursesDto = append(coursesDto, courseDto)
	}

	log.Debug(coursesDto)
	return coursesDto, nil
}

func (s *courseService) GetCoursesByName(query string) (dto.CoursesDto, errores.ApiError) {
	var courses model.Coursess
	courses = s.coursesClient.GetCourseByName(query)
	var coursesDto dto.CoursesDto
	for _, course := range courses {
		var courseDto dto.CourseDto
		courseDto.Course_id = course.Course_id
		courseDto.Nombre = course.Nombre
		courseDto.Profesor_id = course.Profesor_id
		courseDto.Categoria = course.Categoria
		courseDto.Descripcion = course.Descripcion
		courseDto.Valoracion = course.Valoracion
		courseDto.Duracion = course.Duracion
		courseDto.Requisitos = course.Requisitos
		courseDto.Url_image = course.Url_image
		courseDto.Fecha_inicio = course.Fecha_inicio

		coursesDto = append(coursesDto, courseDto)
	}

	log.Debug(coursesDto)
	return coursesDto, nil
}

func (s *courseService) InsertCourse(creacion dto.CourseDto) (dto.CourseDto, errores.ApiError) {
	var nuevo model.Courses

	nuevo.Course_id = creacion.Course_id
	nuevo.Nombre = creacion.Nombre
	nuevo.Profesor_id = creacion.Profesor_id
	nuevo.Categoria = creacion.Categoria
	nuevo.Descripcion = creacion.Descripcion
	nuevo.Valoracion = creacion.Valoracion
	nuevo.Duracion = creacion.Duracion
	nuevo.Requisitos = creacion.Requisitos
	nuevo.Url_image = creacion.Url_image
	nuevo.Fecha_inicio = creacion.Fecha_inicio

	nuevo, err := s.coursesClient.InsertCourse(nuevo)
	if err != nil {
		return creacion, errores.NewInternalServerApiError("Error al crear el curso", err)
	}

	var creacionResponse dto.CourseDto
	creacionResponse.Course_id = nuevo.Course_id
	creacionResponse.Nombre = nuevo.Nombre
	creacionResponse.Profesor_id = nuevo.Profesor_id
	creacionResponse.Duracion = nuevo.Duracion
	creacionResponse.Fecha_inicio = nuevo.Fecha_inicio
	creacionResponse.Categoria = nuevo.Categoria
	creacionResponse.Requisitos = nuevo.Requisitos
	creacionResponse.Valoracion = nuevo.Valoracion
	creacionResponse.Descripcion = nuevo.Descripcion
	creacionResponse.Url_image = nuevo.Url_image

	return creacionResponse, nil

}

func (s *courseService) EditCourse(id int, course dto.CourseDto) (dto.CourseDto, errores.ApiError) {
	var existingCourse model.Courses
	existingCourse = s.coursesClient.GetCourseById(id)
	if existingCourse.Course_id < 0 {
		return dto.CourseDto{}, errores.NewBadRequestApiError("course not found")
	}

	var updatedCourse model.Courses
	updatedCourse.Course_id = id
	updatedCourse.Nombre = course.Nombre
	updatedCourse.Profesor_id = course.Profesor_id
	updatedCourse.Categoria = course.Categoria
	updatedCourse.Descripcion = course.Descripcion
	updatedCourse.Valoracion = course.Valoracion
	updatedCourse.Duracion = course.Duracion
	updatedCourse.Requisitos = course.Requisitos
	updatedCourse.Url_image = course.Url_image
	updatedCourse.Fecha_inicio = course.Fecha_inicio

	updatedCourse, err := s.coursesClient.EditCourse(id, updatedCourse)
	if err != nil {
		return dto.CourseDto{}, errores.NewInternalServerApiError("Error al actualizar el curso", err)
	}

	var updatedCourseDto dto.CourseDto
	updatedCourseDto.Course_id = updatedCourse.Course_id
	updatedCourseDto.Nombre = updatedCourse.Nombre
	updatedCourseDto.Profesor_id = updatedCourse.Profesor_id
	updatedCourseDto.Categoria = updatedCourse.Categoria
	updatedCourseDto.Descripcion = updatedCourse.Descripcion
	updatedCourseDto.Valoracion = updatedCourse.Valoracion
	updatedCourseDto.Duracion = updatedCourse.Duracion
	updatedCourseDto.Requisitos = updatedCourse.Requisitos
	updatedCourseDto.Url_image = updatedCourse.Url_image
	updatedCourseDto.Fecha_inicio = updatedCourse.Fecha_inicio

	return updatedCourseDto, nil
}

func (s *courseService) DeleteCourse(id int) (dto.CourseDto, errores.ApiError) {
	var course model.Courses = s.coursesClient.GetCourseById(id)
	var courseDto dto.CourseDto

	if course.Course_id < 0 {
		return courseDto, errores.NewBadRequestApiError("course not found")
	}

	err := s.coursesClient.DeleteCourse(id)
	if err != nil {
		return courseDto, errores.NewInternalServerApiError("Error al eliminar el curso", err)
	}

	courseDto.Course_id = course.Course_id
	courseDto.Nombre = course.Nombre
	courseDto.Profesor_id = course.Profesor_id
	courseDto.Categoria = course.Categoria
	courseDto.Descripcion = course.Descripcion
	courseDto.Valoracion = course.Valoracion
	courseDto.Duracion = course.Duracion
	courseDto.Requisitos = course.Requisitos
	courseDto.Url_image = course.Url_image
	courseDto.Fecha_inicio = course.Fecha_inicio
	return courseDto, nil
}

/*func DeleteInscripcionesByCursoID(cursoID string) error {
	return db.DeleteInscripcionesByCursoID(cursoID)
}

// DeleteCurso elimina un curso y sus inscripciones asociadas
func DeleteCurso(cursoID string) error {
	// Primero eliminamos las inscripciones asociadas al curso
	err := DeleteInscripcionesByCursoID(cursoID)
	if err != nil {
		return err
	}

	// Luego eliminamos el curso
	return db.DeleteCursoByID(cursoID)
}
*/
