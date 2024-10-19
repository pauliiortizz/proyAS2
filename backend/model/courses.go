package model

import "time"

type Courses struct {
	Course_id    int       `gorm:"primaryKey;AUTO_INCREMENT"`
	Nombre       string    `gorm:"type:varchar(100);not null"`
	Profesor_id  int       `gorm:"foreignKey:User_id"`
	Categoria    string    `gorm:"type:varchar(100);not null"`
	Descripcion  string    `gorm:"type:varchar(500);not null"`
	Valoracion   float64   `gorm:"type:decimal;unsigned;not null"`
	Duracion     int       `gorm:"type:int;unsigned;not null"`
	Requisitos   string    `gorm:"type:varchar(500);not null"`
	Url_image    string    `gorm:"type:varchar(400);not null"`
	Fecha_inicio time.Time `gorm:"type:date;not null"`
}
type Coursess []Courses
