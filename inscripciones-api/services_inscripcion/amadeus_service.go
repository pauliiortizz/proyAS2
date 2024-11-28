package services_inscripcion

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

type courseService struct{}

type courseServiceInterface interface {
	CheckAvailabilityCourse(courseId string) (bool, error)
	UpdateCourseSeats(courseId string, seats int) error
}

var CourseService courseServiceInterface

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

var (
	apiToken        string
	mutex           sync.Mutex
	lastTokenUpdate time.Time
)

func init() {
	CourseService = &courseService{}
}

// Función para obtener el token
func GetAPIToken() (string, error) {
	mutex.Lock()
	defer mutex.Unlock()

	// Actualizar el token si es necesario
	if time.Since(lastTokenUpdate) > 30*time.Minute {
		err := updateAPIToken()
		if err != nil {
			return "", err
		}
	}

	return apiToken, nil
}

// Función para actualizar el token
func updateAPIToken() error {
	url := "https://test.api.amadeus.com/v1/security/oauth2/token"
	method := "POST"

	payload := strings.NewReader("grant_type=client_credentials&client_id=CLIENT_ID&client_secret=CLIENT_SECRET")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var response TokenResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	apiToken = response.AccessToken
	lastTokenUpdate = time.Now()

	fmt.Println("Token actualizado:", apiToken)
	return nil
}

// Comprobar disponibilidad de un curso
func (s *courseService) CheckAvailabilityCourse(courseId string) (bool, error) {
	url := fmt.Sprintf("http://cursos-api:8081/courses/%s", courseId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, fmt.Errorf("error al crear la solicitud: %v", err)
	}

	token, err := GetAPIToken()
	if err != nil {
		return false, fmt.Errorf("error al obtener el token: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("error al realizar la solicitud: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return false, fmt.Errorf("error en la respuesta: %s", string(body))
	}

	var response struct {
		Available bool `json:"available"`
	}

	// Leer el cuerpo de la respuesta
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return false, fmt.Errorf("error al leer el cuerpo de la respuesta: %v", err)
	}

	// Decodificar la respuesta JSON
	if err := json.Unmarshal(body, &response); err != nil {
		return false, fmt.Errorf("error al decodificar la respuesta JSON: %v", err)
	}

	return response.Available, nil
}

// Actualizar los cupos disponibles de un curso
func (s *courseService) UpdateCourseSeats(courseId string, seats int) error {
	url := fmt.Sprintf("http://cursos-api:8081/courses/%s", courseId)
	method := "PUT"

	payload := fmt.Sprintf(`{"seats": %d}`, seats)
	req, err := http.NewRequest(method, url, strings.NewReader(payload))
	if err != nil {
		return fmt.Errorf("error al crear la solicitud: %v", err)
	}

	token, err := GetAPIToken()
	if err != nil {
		return fmt.Errorf("error al obtener el token: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error al realizar la solicitud: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("error en la respuesta: %s", string(body))
	}

	return nil
}
