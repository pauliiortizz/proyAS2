package dao_archivos

type Archivo struct {
	Archivo_id int  `bson:"archivo_id"`
	Course_id  int  `bson:"course_id"`
	Archivo blob `bson:"archivo"`

}

type Archivos []Archivo