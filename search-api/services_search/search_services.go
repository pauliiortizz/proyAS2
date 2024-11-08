package services_search

import (
	"context"
	"fmt"
	cursosDAO "search/dao_search"
	cursosDomain "search/domain_search"
)

type Repository interface {
	Index(ctx context.Context, hotel cursosDAO.Search) (string, error)
	Update(ctx context.Context, hotel cursosDAO.Search) error
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, query string, limit int, offset int) ([]cursosDAO.Search, error) // Updated signature
}

type ExternalRepository interface {
	GetCourseByID(ctx context.Context, id string) (cursosDomain.CourseDto, error)
}

type Service struct {
	repository Repository
	cursosAPI  ExternalRepository
}

func NewService(repository Repository, cursosAPI ExternalRepository) Service {
	return Service{
		repository: repository,
		cursosAPI:  cursosAPI,
	}
}

func (service Service) Search(ctx context.Context, query string, offset int, limit int) ([]cursosDomain.CourseDto, error) {
	// Call the repository's Search method
	cursosDAOList, err := service.repository.Search(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error searching hotels: %w", err)
	}

	// Convert the dao layer hotels to domain layer hotels
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

func (service Service) HandleCursoNew(cursoNew cursosDomain.CourseNew) {
	switch cursoNew.Operation {
	case "CREATE", "UPDATE":
		// Fetch hotel details from the local service
		curso, err := service.cursosAPI.GetCourseByID(context.Background(), cursoNew.Course_id)
		if err != nil {
			fmt.Printf("Error getting hotel (%s) from API: %v\n", cursoNew.Course_id, err)
			return
		}

		cursoDAO := cursosDAO.Search{
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
		}

		// Handle Index operation
		if cursoNew.Operation == "CREATE" {
			if _, err := service.repository.Index(context.Background(), cursoDAO); err != nil {
				fmt.Printf("Error indexing hotel (%s): %v\n", cursoNew.Course_id, err)
			} else {
				fmt.Println("Hotel indexed successfully:", cursoNew.Course_id)
			}
		} else { // Handle Update operation
			if err := service.repository.Update(context.Background(), cursoDAO); err != nil {
				fmt.Printf("Error updating hotel (%s): %v\n", cursoNew.Course_id, err)
			} else {
				fmt.Println("Hotel updated successfully:", cursoNew.Course_id)
			}
		}

	case "DELETE":
		// Call Delete method directly since no hotel details are needed
		if err := service.repository.Delete(context.Background(), cursoNew.Course_id); err != nil {
			fmt.Printf("Error deleting hotel (%s): %v\n", cursoNew.Course_id, err)
		} else {
			fmt.Println("Hotel deleted successfully:", cursoNew.Course_id)
		}

	default:
		fmt.Printf("Unknown operation: %s\n", cursoNew.Operation)
	}
}
