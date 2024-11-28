package domain_search

import "time"

type CourseDto struct {
	Course_id    string    `json:"course_id"`
	Nombre       string    `json:"nombre"`
	Profesor_id  int       `json:"profesor_id"`
	Categoria    string    `json:"categoria"`
	Descripcion  string    `json:"descripcion"`
	Valoracion   float64   `json:"valoracion"`
	Duracion     int       `json:"duracion"`
	Requisitos   string    `json:"requisitos"`
	Url_image    string    `json:"url_image"`
	Fecha_inicio time.Time `json:"fecha_inicio"`
	Capacidad    int       `json:"capacidad"`
}

type CoursesDto []CourseDto

type CourseNew struct {
	Operation string `json:"operation"`
	Curso_id  string `json:"course_id"`
}
