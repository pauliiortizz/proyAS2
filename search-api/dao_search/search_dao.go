package dao_search

import "time"

type Search struct {
	Course_id    string    `bson:"course_id"`
	Nombre       string    `bson:"nombre"`
	Profesor_id  int       `bson:"profesor_id"`
	Categoria    string    `bson:"categoria"`
	Descripcion  string    `bson:"descripcion"`
	Valoracion   float64   `bson:"valoracion"`
	Duracion     int       `bson:"duracion"`
	Requisitos   string    `bson:"requisitos"`
	Url_image    string    `bson:"url_image"`
	Fecha_inicio time.Time `bson:"fecha_inicio"`
	Capacidad    int       `bson:"capacidad"`
}

type Searchs []Search
