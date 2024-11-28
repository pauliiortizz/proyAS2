package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"inscripciones/domain_inscripcion"
	"net/http"
)

type HttpClient struct{}

func (h *HttpClient) GetUser(userID int) (domain_inscripcion.User, error) {
	userUrl := fmt.Sprintf("http://users-api:8080/users/%d", userID)
	userResp, err := http.Get(userUrl)
	if err != nil {
		return domain_inscripcion.User{}, fmt.Errorf("error making request to users-api: %w", err)
	}
	defer userResp.Body.Close()

	if userResp.StatusCode != http.StatusOK {
		return domain_inscripcion.User{}, errors.New("user not found")
	}

	var userDto domain_inscripcion.User
	if err := json.NewDecoder(userResp.Body).Decode(&userDto); err != nil {
		return domain_inscripcion.User{}, fmt.Errorf("error decoding user response: %w", err)
	}
	return userDto, nil
}

func (h *HttpClient) GetCourse(courseID string) (domain_inscripcion.CourseDto, error) {
	courseUrl := fmt.Sprintf("http://cursos-api:8081/courses/%s", courseID)
	courseResp, err := http.Get(courseUrl)
	if err != nil {
		return domain_inscripcion.CourseDto{}, fmt.Errorf("error making request to courses-api: %w", err)
	}
	defer courseResp.Body.Close()

	if courseResp.StatusCode != http.StatusOK {
		return domain_inscripcion.CourseDto{}, errors.New("course not found")
	}

	var courseDto domain_inscripcion.CourseDto
	if err := json.NewDecoder(courseResp.Body).Decode(&courseDto); err != nil {
		return domain_inscripcion.CourseDto{}, fmt.Errorf("error decoding course response: %w", err)
	}
	return courseDto, nil
}

func (h *HttpClient) GetCourses() ([]domain_inscripcion.CourseDto, error) {
	// Construir la URL para obtener todos los cursos
	coursesUrl := "http://cursos-api:8081/courses"

	// Realizar la solicitud HTTP GET
	coursesResp, err := http.Get(coursesUrl)
	if err != nil {
		return nil, fmt.Errorf("error making request to courses-api: %w", err)
	}
	defer coursesResp.Body.Close()

	// Verificar si la respuesta tiene un código de estado 200 OK
	if coursesResp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch courses, unexpected status code")
	}

	// Decodificar el cuerpo de la respuesta en una lista de CourseDto
	var courses []domain_inscripcion.CourseDto
	if err := json.NewDecoder(coursesResp.Body).Decode(&courses); err != nil {
		return nil, fmt.Errorf("error decoding courses response: %w", err)
	}

	return courses, nil
}

func (h *HttpClient) UpdateCourse(course domain_inscripcion.CourseDto, resultChan chan<- error) {
	go func() {
		// Construir la URL para actualizar un curso
		courseUrl := fmt.Sprintf("http://cursos-api:8081/edit/%s", course.Course_id)

		// Serializar el curso en formato JSON
		payload, err := json.Marshal(course)
		if err != nil {
			resultChan <- fmt.Errorf("error marshaling course data: %w", err)
			return
		}

		// Crear una nueva solicitud HTTP PUT
		req, err := http.NewRequest(http.MethodPut, courseUrl, bytes.NewBuffer(payload))
		if err != nil {
			resultChan <- fmt.Errorf("error creating PUT request: %w", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		// Ejecutar la solicitud
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			resultChan <- fmt.Errorf("error making PUT request to courses-api: %w", err)
			return
		}
		defer resp.Body.Close()

		// Verificar el código de respuesta
		if resp.StatusCode != http.StatusOK {
			var errorResponse map[string]interface{}
			if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
				resultChan <- fmt.Errorf("failed to update course, status code: %d", resp.StatusCode)
				return
			}
			resultChan <- fmt.Errorf("failed to update course, response: %v", errorResponse)
			return
		}

		// Si todo salió bien, enviar nil al canal
		resultChan <- nil
	}()
}
