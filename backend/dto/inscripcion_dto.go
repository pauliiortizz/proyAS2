package dto

import "time"

type InscripcionDto struct {
	Id_inscripcion    int       `json:"id_inscripcion"`
	Id_user           int       `json:"id_user"`
	Id_course         int       `json:"id_course"`
	Fecha_inscripcion time.Time `json:"fecha_inscripcion"`
}
type inscripcionesDto []InscripcionDto
