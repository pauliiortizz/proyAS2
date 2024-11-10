package services_inscripcion

import (
	"encoding/json"
	"errors"
	"fmt"
	"inscripciones/dao_inscripcion"
	"inscripciones/domain_inscripcion"
	"inscripciones/repositories_inscripcion"
	"inscripciones/utils"
	"net/http"
	"time"
)

func NewService(repository repositories_inscripcion.Repository) InscripcionServiceInterface {
	return &inscripcionService{
		HTTPClient: &utils.HttpClient{},
		repository: repository,
	}
}

type inscripcionService struct {
	HTTPClient utils.HttpClientInterface
	repository repositories_inscripcion.Repository
}

type InscripcionServiceInterface interface {
	InsertInscripcion(inscripcionDto domain_inscripcion.InscripcionDto) (domain_inscripcion.InscripcionDto, error)
	GetInscripcionByUserID(userID int) ([]dao_inscripcion.Inscripcion, error)
	GetInscripcionByCourseID(courseID int) ([]dao_inscripcion.Inscripcion, error)
}

func (s *inscripcionService) InsertInscripcion(inscripcionDto domain_inscripcion.InscripcionDto) (domain_inscripcion.InscripcionDto, error) {
	// Obtener el usuario a través de la API de usuario
	userUrl := fmt.Sprintf("http://localhost:8080/users/%d", inscripcionDto.Id_user)
	userResp, err := s.HTTPClient.Get(userUrl)
	if err != nil {
		return inscripcionDto, fmt.Errorf("error making request to users-api: %w", err)
	}
	defer userResp.Body.Close()

	if userResp.StatusCode != http.StatusOK {
		return inscripcionDto, errors.New("user not found")
	}

	var userDto domain_inscripcion.User
	if err := json.NewDecoder(userResp.Body).Decode(&userDto); err != nil {
		return inscripcionDto, fmt.Errorf("error decoding user response: %w", err)
	}

	// Obtener el curso a través de la API de cursos
	courseUrl := fmt.Sprintf("http://localhost:27017/courses/%s", inscripcionDto.Id_course)
	courseResp, err := s.HTTPClient.Get(courseUrl)
	if err != nil {
		return inscripcionDto, fmt.Errorf("error making request to courses-api: %w", err)
	}
	defer courseResp.Body.Close()

	if courseResp.StatusCode != http.StatusOK {
		return inscripcionDto, errors.New("course not found")
	}

	var courseDto domain_inscripcion.CourseDto
	if err := json.NewDecoder(courseResp.Body).Decode(&courseDto); err != nil {
		return inscripcionDto, fmt.Errorf("error decoding course response: %w", err)
	}

	// Verificar que la fecha de inscripción sea anterior a la fecha de inicio del curso
	inscripcionDto.Fecha_inscripcion = time.Now()

	if !inscripcionDto.Fecha_inscripcion.Before(courseDto.Fecha_inicio) {
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

func (s *inscripcionService) GetInscripcionByUserID(userID int) ([]dao_inscripcion.Inscripcion, error) {
	return s.repository.GetInscripcionByUserID(userID)
}

func (s *inscripcionService) GetInscripcionByCourseID(courseID int) ([]dao_inscripcion.Inscripcion, error) {
	return s.repository.GetInscripcionByCourseID(courseID)
}
