package dto

type ArchivoDto struct {
	Id_inscripcion int    `json:"id_inscripcion"`
	Id_user        int    `json:"id_user"`
	Id_course      int    `json:"id_course"`
	Archivo        []byte `json:"archivo"`
}

type ArchivosDto []ArchivoDto
