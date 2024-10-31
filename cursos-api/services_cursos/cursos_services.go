package services_cursos

import (
	"context"
	cursosDAO "cursos/dao_cursos"
	cursosDomain "cursos/domain_cursos"
	"fmt"
)

type Repository interface {
	GetCourseByID(ctx context.Context, id string) (cursosDAO.Curso, error)
	Create(ctx context.Context, curso cursosDAO.Curso) (string, error)
	//Update(ctx context.Context, curso cursosDAO.Curso) error
	//Delete(ctx context.Context, id string) error
}

type Queue interface {
	Publish(cursoNuevo cursosDomain.CourseNew) error
}

type Service struct {
	mainRepository Repository
	eventsQueue    Queue
}

func NewService(mainRepository Repository, eventsQueue Queue) Service {
	return Service{
		mainRepository: mainRepository,
		eventsQueue:    eventsQueue,
	}
}

func (service Service) GetCourseByID(ctx context.Context, id string) (cursosDomain.CourseDto, error) {
	cursosDAO, err := service.mainRepository.GetCourseByID(ctx, id)
	if err != nil {
		// Get curso from main repository
		cursosDAO, err = service.mainRepository.GetCourseByID(ctx, id)
		if err != nil {
			return cursosDomain.CourseDto{}, fmt.Errorf("error getting curso from repository: %v", err)
		}

	}

	// Convert DAO to DTO
	return cursosDomain.CourseDto{
		Course_id:   cursosDAO.Course_id,
		Nombre:      cursosDAO.Nombre,
		Categoria:   cursosDAO.Categoria,
		Descripcion: cursosDAO.Descripcion,
		Valoracion:  cursosDAO.Valoracion,
		Duracion:    cursosDAO.Duracion,
		Requisitos:  cursosDAO.Requisitos,
	}, nil
}

func (service Service) Create(ctx context.Context, curso cursosDomain.CourseDto) (string, error) {
	record := cursosDAO.Curso{
		Course_id:   curso.Course_id,
		Categoria:   curso.Categoria,
		Descripcion: curso.Descripcion,
		Valoracion:  curso.Valoracion,
		Duracion:    curso.Duracion,
		Requisitos:  curso.Requisitos,
	}
	id, err := service.mainRepository.Create(ctx, record)
	if err != nil {
		return "", fmt.Errorf("error creating curso in main repository: %w", err)
	}
	if err := service.eventsQueue.Publish(cursosDomain.CourseNew{
		Operation: "CREATE",
		Course_id: curso.Course_id,
	}); err != nil {
		return "", fmt.Errorf("error publishing curso new: %w", err)
	}

	return id, nil
}

// Update(ctx context.Context, curso cursosDAO.Curso) error
/*
func (service Service) Update(ctx context.Context, curso cursosDomain.CourseDto) error {
	// Convert domain model to DAO model
	record := cursosDAO.Curso{
		Course_id:   curso.Course_id,
		Categoria:   curso.Categoria,
		Descripcion: curso.Descripcion,
		Valoracion:  curso.Valoracion,
		Duracion:    curso.Duracion,
		Requisitos:  curso.Requisitos,
	}

	// Update the curso in the main repository
	err := service.mainRepository.Update(ctx, record)
	if err != nil {
		return fmt.Errorf("error updating curso in main repository: %w", err)
	}

	// Publish an event for the update operation
	if err := service.eventsQueue.Publish(cursosDomain.CourseNew{
		Operation: "UPDATE",
		Course_id: curso.Course_id,
	}); err != nil {
		return fmt.Errorf("error publishing curso update: %w", err)
	}

	return nil
}*/
