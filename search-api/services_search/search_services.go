package services_search

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	cursosDomain "search/domain_search"
)

// Repository define las funciones que el repositorio debe implementar
type Repository interface {
	Index(ctx context.Context, curso cursosDomain.CourseDto) (string, error)
	Update(ctx context.Context, curso cursosDomain.CourseDto) error
	Search(ctx context.Context, query string, limit int, offset int) ([]cursosDomain.CourseDto, error)
}

// ExternalRepository define el método para obtener un curso de la API externa
type ExternalRepository interface {
	GetCourseByID(ctx context.Context, id string) (cursosDomain.CourseDto, error)
}

// Service estructura que maneja la lógica del servicio
type Service struct {
	repository Repository
	cursosAPI  ExternalRepository
}

// NewService inicializa y devuelve un Service
func NewService(repository Repository, cursosAPI ExternalRepository) Service {
	return Service{
		repository: repository,
		cursosAPI:  cursosAPI,
	}
}

// Search realiza una búsqueda de cursos
func (service Service) Search(ctx context.Context, query string, offset int, limit int) ([]cursosDomain.CourseDto, error) {
	cursosDAOList, err := service.repository.Search(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error searching courses: %w", err)
	}

	// Convertir cursosDAOList a cursosDomainList
	cursosDomainList := make([]cursosDomain.CourseDto, 0)
	for _, curso := range cursosDAOList {
		cursosDomainList = append(cursosDomainList, cursosDomain.CourseDto{
			Course_id:    curso.Course_id,
			Nombre:       curso.Nombre,
			Profesor_id:  curso.Profesor_id,
			Categoria:    curso.Categoria,
			Descripcion:  curso.Descripcion,
			Valoracion:   curso.Valoracion,
			Duracion:     curso.Duracion,
			Requisitos:   curso.Requisitos,
			Url_image:    curso.Url_image,
			Fecha_inicio: curso.Fecha_inicio,
		})
	}

	return cursosDomainList, nil
}

// HandleCourseNew maneja la creación o actualización de cursos
func (service Service) HandleCourseNew(cursoNew cursosDomain.CourseNew) {
	switch cursoNew.Operation {
	case "CREATE", "UPDATE":
		curso, err := service.cursosAPI.GetCourseByID(context.Background(), cursoNew.Curso_id)
		if err != nil {
			fmt.Printf("Error getting course (%s) from API: %v\n", cursoNew.Curso_id, err)
			return
		}

		if cursoNew.Operation == "CREATE" {
			if _, err := service.repository.Index(context.Background(), curso); err != nil {
				fmt.Printf("Error indexing course (%s): %v\n", cursoNew.Curso_id, err)
			} else {
				fmt.Println("Course indexed successfully:", cursoNew.Curso_id)
			}
		} else {
			if err := service.repository.Update(context.Background(), curso); err != nil {
				fmt.Printf("Error updating course (%s): %v\n", cursoNew.Curso_id, err)
			} else {
				fmt.Println("Course updated successfully:", cursoNew.Curso_id)
			}
		}

	default:
		fmt.Printf("Unknown operation: %s\n", cursoNew.Operation)
	}
}

// SolrRepository define el repositorio para interactuar con Solr
type SolrRepository struct {
	solrURL string
}

// NewSolrRepository inicializa y devuelve un SolrRepository
func NewSolrRepository(solrHost, solrPort, coreName string) *SolrRepository {
	solrURL := fmt.Sprintf("http://%s:%s/solr/%s/update?commit=true", solrHost, solrPort, coreName)
	return &SolrRepository{solrURL: solrURL}
}

// Index envía un documento de curso a Solr para indexación
func (repo *SolrRepository) Index(ctx context.Context, curso cursosDomain.CourseDto) (string, error) {
	doc := map[string]interface{}{
		"course_id":    curso.Course_id,
		"nombre":       curso.Nombre,
		"profesor_id":  curso.Profesor_id,
		"categoria":    curso.Categoria,
		"descripcion":  curso.Descripcion,
		"valoracion":   curso.Valoracion,
		"duracion":     curso.Duracion,
		"requisitos":   curso.Requisitos,
		"url_image":    curso.Url_image,
		"fecha_inicio": curso.Fecha_inicio,
	}

	indexRequest := map[string]interface{}{
		"add": []interface{}{doc},
	}

	jsonData, err := json.Marshal(indexRequest)
	if err != nil {
		return "", fmt.Errorf("error marshaling course data to JSON: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", repo.solrURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request to Solr: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error indexing course in Solr, status code: %d", resp.StatusCode)
	}

	return curso.Course_id, nil
}

// Update modifica un documento de curso existente en la colección de Solr
func (repo *SolrRepository) Update(ctx context.Context, curso cursosDomain.CourseDto) error {
	doc := map[string]interface{}{
		"course_id":    curso.Course_id,
		"nombre":       curso.Nombre,
		"profesor_id":  curso.Profesor_id,
		"categoria":    curso.Categoria,
		"descripcion":  curso.Descripcion,
		"valoracion":   curso.Valoracion,
		"duracion":     curso.Duracion,
		"requisitos":   curso.Requisitos,
		"url_image":    curso.Url_image,
		"fecha_inicio": curso.Fecha_inicio,
	}

	updateRequest := map[string]interface{}{
		"add": []interface{}{doc},
	}

	// Marshal the update request to JSON
	body, err := json.Marshal(updateRequest)
	if err != nil {
		return fmt.Errorf("error marshaling course document: %w", err)
	}

	// Create a new HTTP request with the body
	req, err := http.NewRequestWithContext(ctx, "POST", repo.solrURL, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error updating course: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update course in Solr, status code: %d", resp.StatusCode)
	}

	return nil
}
