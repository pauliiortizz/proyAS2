package dao_comentarios

import "time"

type Comentario struct {
	Comentario_id int `bson:"comentario_id"`
	Course_id       int `bson:"user_id"`
	Comentario   string  `bson:"comentario"`
	Valoracion float32 `bson:"valoracion"`
	Fecha_inscripcion date `bson:"fecha_inscripcion"`
}

type Comentarios []Comentario
