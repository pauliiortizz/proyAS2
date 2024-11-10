package clients_inscripcion

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"inscripciones/dao_inscripcion"
)

var Db *gorm.DB

func InsertInscripcion(inscripcion dao_inscripcion.Inscripcion) dao_inscripcion.Inscripcion {
	result := Db.Create(&inscripcion)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("Inscripcion Creada: ", inscripcion.Id_inscripcion)
	return inscripcion
}

func GetInscripcionByCourse(idMongo string) dao_inscripcion.Inscripciones {
	var inscripciones dao_inscripcion.Inscripciones

	Db.Where("id_course = ?", idMongo).Find(&inscripciones)
	log.Debug("Inscripciones: ", inscripciones)

	return inscripciones

}

func GetInscripcionByUserId(idUser string) dao_inscripcion.Inscripciones {
	var inscripciones dao_inscripcion.Inscripciones

	Db.Where("id_user = ?", idUser).Find(&inscripciones)
	log.Debug("Inscripciones: ", inscripciones)

	return inscripciones
}
