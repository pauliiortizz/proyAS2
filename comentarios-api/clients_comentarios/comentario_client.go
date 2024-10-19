package clients_comentarios
import (
	"comentarios-api/clients_comentarios"
	"comentarios-api/dao_comentarios"
	log "github.com/sirupsen/logrus"
)

type comentarioClient struct{}

type ComentarioClientInterface interface {
	InsertValoracion(id int, valoracion dao_comentarios.Comentarios) (dao_comentarios.Comentarios, error)
	GetComentByCourseId(id int) dao_comentarios.Comentario
}

var (
	ComentarioClient ComentarioClientInterface
)

func init() {
	ComentarioClient = &comentarioClient{}
}


func (s *comentarioClient) InsertValoracion(id int, comentario dao_comentarios.Comentario) dao_comentarios.Comentario, error) {
	var comentario dao_comentarios.Comentario
	// Buscar la inscripción por ID
	if err := clients.Db.Where("users_x_courses_id = ?", id).First(&inscripcion).Error; err != nil { //modificar
		log.Error("Error al encontrar el curso: ", err)
		return comentario, err
	}

	// Actualizar solo las columnas valoracion y comentario
	result := clients.Db.Model(&inscripcion).Updates(map[string]interface{}{
		"valoracion": comentario.Valoracion,
		"comentario": comentario.Comentario,
	})
	if result.Error != nil {
		log.Error("Error al actualizar el comentario: ", result.Error)
		return comentario, result.Error
	}

	log.Debug("Comentario actualizado: ", comentario.Comentario_id)
	return comentario, nil
}

func (s *comentarioClient) GetComentByCourseId(id int) dao_comentarios.Comentarios {
	var comentarios dao_comentarios.Comentarios
	clients.Db.Where("course_id = ?", id).Find(&inscripciones) //cambiar

	log.Debug("Comentario: ", comentarios)
	return comentarios
}
