package clients_inscripcion
import (
	"inscripciones-api/clients_inscripcion"
	"inscripciones-api/dao_inscripcion"
	log "github.com/sirupsen/logrus"
)

type inscripcionClient struct{}

type InscripcionClientInterface interface {
	InsertInscr(inscripcion dao_inscripcion.Inscripcion) (dao_inscripcion.Inscripcion, error)
	GetInscripcion() dao_inscripcion.Inscripciones
	GetCourseByUserId(id int) dao_inscripcion.Inscripciones
}

var (
	InscripcionClient InscripcionClientInterface
)

func init() {
	InscripcionClient = &inscripcionClient{}
}

func (s *inscripcionClient) InsertInscr(inscripcion dao_inscripcion.Inscripcion) (dao_inscripcion.Inscripcion, error) {
	result := clients.Db.Create(&inscripcion) //adaptar
	if result.Error != nil {
		log.Error("Error al crear la inscripción: ", result.Error)
		return inscripcion, result.Error
	}

	log.Debug("Inscripcion creada: ", inscripcion.Id_inscripcion)
	return inscripcion, nil
}

func (s *inscripcionClient) GetInscripcion()dao_inscripcion.Inscripciones {
	var inscripciones dao_inscripcion.Inscripciones
	clients.Db.Find(&inscripciones) //adaptar

	log.Debug("Inscripciones: ", inscripciones)
	return inscripciones
}

func (s *inscripcionClient) GetCourseByUserId(id int) dao_inscripcion.Inscripciones {
	var inscripciones dao_inscripcion.Inscripciones
	clients.Db.Where("user_id = ?", id).Find(&inscripciones) //adaptar

	log.Debug("Inscripciones: ", inscripciones)
	return inscripciones
}


