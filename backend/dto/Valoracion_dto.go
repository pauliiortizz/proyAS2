package dto

type ValoracionDto struct {
	Id_inscripcion int     `json:"id_inscripcion"`
	Id_user        int     `json:"id_user"`
	Id_course      int     `json:"id_course"`
	Valoracion     float64 `json:"valoracion"`
	Comentario     string  `json:"comentario"`
}

type ValoracionesDto []ValoracionDto
