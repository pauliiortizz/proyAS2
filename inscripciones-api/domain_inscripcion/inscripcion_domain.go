package domain_inscripcion

import "time"

type InscripcionDto struct {
	Id_inscripcion    int       `json:"id_inscripcion"`
	Id_user           int       `json:"id_user"`
	Id_course         int       `json:"id_course"`
	Fecha_inscripcion time.Time `json:"fecha_inscripcion"`
}
type inscripcionesDto []InscripcionDto

type User struct {
	User_id  int64  `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Nombre   string `json:"first_name"`
	Apellido string `json:"last_name"`
	Admin    bool   `json:"admin"`
}

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
}
