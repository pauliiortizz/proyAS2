package dao_inscripcion

import "time"

type Inscripcion struct {
	Id_inscripcion    int       `bson:"id_inscripcion"`
	Id_user           int       `bson:"id_user"`
	Id_course         int       `bson:"id_course"`
	Fecha_inscripcion time.Time `bson:"fecha_inscripcion"`
}

type Inscripciones []Inscripcion
