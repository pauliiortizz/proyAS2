package services_inscripcion

import (
	"errors"
	"fmt"
	"inscripciones/dao_inscripcion"
	"inscripciones/domain_inscripcion"
	"inscripciones/utils"
	//"strconv"
	"time"
)

type InscripcionRepository interface {
	InsertInscripcion(inscripcionDto dao_inscripcion.Inscripcion) (int64, error)
	GetInscripcionByUserID(userID int) ([]dao_inscripcion.Inscripcion, error)
	GetInscripcionByCourseID(courseID string) ([]dao_inscripcion.Inscripcion, error)
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

	// Verificar disponibilidad del curso antes de proceder con la inscripción
	if course.Capacidad <= 0 {
		return inscripcionDto, errors.New("no hay cupos disponibles para el curso")
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

	// Reducir la capacidad del curso
	course.Capacidad -= 1

	// Crear un canal para manejar el resultado de UpdateCourse
	resultChan := make(chan error)

	// Llamar a UpdateCourse con goroutines
	s.HTTPClient.UpdateCourse(course, resultChan)

	// Escuchar el canal en otra goroutine para manejar la respuesta sin bloquear
	go func() {
		err := <-resultChan
		if err != nil {
			// En caso de error, puedes manejarlo de forma específica o simplemente loguearlo
			fmt.Printf("Error updating course capacity: %v\n", err)
		} else {
			fmt.Println("Course capacity updated successfully!")
		}
	}()

	// Asignar el ID generado a inscripcionDto
	inscripcionDto.Id_inscripcion = int(id)

	return inscripcionDto, nil
}

func (s *InscripcionService) GetInscripcionByUserID(userID int) ([]dao_inscripcion.Inscripcion, error) {
	return s.repository.GetInscripcionByUserID(userID)
}

func (s *InscripcionService) GetInscripcionByCourseID(courseID string) ([]dao_inscripcion.Inscripcion, error) {
	return s.repository.GetInscripcionByCourseID(courseID)
}
