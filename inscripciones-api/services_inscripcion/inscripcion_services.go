package services_inscripcion

import (
	"errors"
	"fmt"
	"inscripciones/dao_inscripcion"
	"inscripciones/domain_inscripcion"
	"inscripciones/utils"
	"strconv"
	"time"
)

type InscripcionRepository interface {
	InsertInscripcion(inscripcionDto dao_inscripcion.Inscripcion) (int64, error)
	GetInscripcionByUserID(userID int) ([]dao_inscripcion.Inscripcion, error)
	GetInscripcionByCourseID(courseID int) ([]dao_inscripcion.Inscripcion, error)
}

type InscripcionService struct {
	HTTPClient utils.HttpClient
	repository InscripcionRepository
}

func NewService(repository InscripcionRepository) *InscripcionService {
	return &InscripcionService{
		HTTPClient: utils.HttpClient{},
		repository: repository,
	}
}

func (s *InscripcionService) InsertInscripcion(inscripcionDto domain_inscripcion.InscripcionDto) (domain_inscripcion.InscripcionDto, error) {
	// Obtener el usuario a través de la API de usuario
	_, err := s.HTTPClient.GetUser(inscripcionDto.Id_user)
	if err != nil {
		return domain_inscripcion.InscripcionDto{}, fmt.Errorf("Error getting user: %w", err)
	}

	// Obtener el curso a través de la API de cursos
	course, err := s.HTTPClient.GetCourse(inscripcionDto.Id_course)
	if err != nil {
		return domain_inscripcion.InscripcionDto{}, fmt.Errorf("Error getting course: %w", err)
	}

	// Verificar que la fecha de inscripción sea anterior a la fecha de inicio del curso
	inscripcionDto.Fecha_inscripcion = time.Now()

	if !inscripcionDto.Fecha_inscripcion.Before(course.Fecha_inicio) {
		return inscripcionDto, errors.New("la inscripción debe realizarse antes de la fecha de inicio del curso")
	}

	// Crear la inscripción en la base de datos
	inscripcion := dao_inscripcion.Inscripcion{
		Fecha_inscripcion: inscripcionDto.Fecha_inscripcion,
		Id_user:           inscripcionDto.Id_user,
		Id_course:         inscripcionDto.Id_course,
	}

	id, err := s.repository.InsertInscripcion(inscripcion)
	if err != nil {
		return inscripcionDto, err
	}

	// Asignar el ID generado a inscripcionDto
	inscripcionDto.Id_inscripcion = int(id)
	return inscripcionDto, nil
}

func (s *InscripcionService) GetInscripcionByUserID(userID int) ([]dao_inscripcion.Inscripcion, error) {
	return s.repository.GetInscripcionByUserID(userID)
}

func (s *InscripcionService) GetInscripcionByCourseID(courseID int) ([]dao_inscripcion.Inscripcion, error) {
	return s.repository.GetInscripcionByCourseID(courseID)
}

func (s *InscripcionService) CheckAvailability(courseID int) (bool, error) {
	// Obtener el curso a través de la API de cursos
	course, err := s.HTTPClient.GetCourse(strconv.Itoa(courseID))
	if err != nil {
		return false, fmt.Errorf("Error getting course: %w", err)
	}

	// Obtener todas las inscripciones asociadas al curso
	inscripciones, err := s.repository.GetInscripcionByCourseID(courseID)
	if err != nil {
		return false, fmt.Errorf("Error getting inscriptions for course: %w", err)
	}

	// Calcular los cupos disponibles
	cuposDisponibles := course.Capacidad - len(inscripciones)
	if cuposDisponibles > 0 {
		return true, nil
	}

	return false, nil
}

func (s *InscripcionService) CheckAllAvailability(fechaActual time.Time) ([]domain_inscripcion.CourseDto, error) {
	// Obtener todos los cursos disponibles
	courses, err := s.HTTPClient.GetCourses()
	if err != nil {
		return nil, fmt.Errorf("Error getting all courses: %w", err)
	}

	var availableCourses []domain_inscripcion.CourseDto

	for _, course := range courses {
		// Convertir el Course_id de string a int
		courseID, err := strconv.Atoi(course.Course_id)
		if err != nil {
			return nil, fmt.Errorf("Error converting Course_id '%s' to int: %w", course.Course_id, err)
		}

		// Verificar que el curso no haya comenzado
		if !fechaActual.Before(course.Fecha_inicio) {
			continue // Saltar este curso si ya ha comenzado
		}

		// Obtener todas las inscripciones asociadas al curso
		inscripciones, err := s.repository.GetInscripcionByCourseID(courseID)
		if err != nil {
			return nil, fmt.Errorf("Error getting inscriptions for course %d: %w", courseID, err)
		}

		// Calcular cupos disponibles
		cuposDisponibles := course.Capacidad - len(inscripciones)
		if cuposDisponibles > 0 {
			availableCourses = append(availableCourses, course)
		}
	}

	return availableCourses, nil
}
