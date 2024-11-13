package dao_inscripcion

import "time"

type Inscripcion struct {
	Id_inscripcion    int       `gorm:"primaryKey;autoIncrement"`
	Id_user           int       `gorm:"id_user"`
	Id_course         string    `gorm:"id_course"`
	Fecha_inscripcion time.Time `gorm:"fecha_inscripcion"`
}

type Inscripciones []Inscripcion
