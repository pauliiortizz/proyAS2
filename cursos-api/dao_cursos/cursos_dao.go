package dao_cursos

import "time"

type Curso struct {
	Course_id    int       `bson:"course_id"`
	Nombre       string    `bson:"nombre"`
	Profesor_id  int       `bson:"profesor_id"`
	Categoria    string    `bson:"categoria"`
	Descripcion  string    `bson:"descripcion"`
	Valoracion   float64   `bson:"valoracion"`
	Duracion     int       `bson:"duracion"`
	Requisitos   string    `bson:"requisitos"`
	Url_image    string    `bson:"url_image"`
	Fecha_inicio time.Time `bson:"fecha_inicio"`
}

type Cursos []Curso