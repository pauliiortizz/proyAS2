package inscripcion

import (
	"backend/clients"
	"backend/model"
	log "github.com/sirupsen/logrus"
)

type inscripcionClient struct{}

type InscripcionClientInterface interface {
	InsertInscr(inscripcion model.Users_x_courses) (model.Users_x_courses, error)
	GetInscripcion() model.Userss_x_coursess
	GetCourseByUserId(id int) model.Userss_x_coursess
	InsertValoracion(id int, valoracion model.Users_x_courses) (model.Users_x_courses, error)
	InsertArchivo(id int, archivo model.Users_x_courses) (model.Users_x_courses, error)
	GetComentByCourseId(id int) model.Userss_x_coursess
}

var (
	InscripcionClient InscripcionClientInterface
)

func init() {
	InscripcionClient = &inscripcionClient{}
}

func (s *inscripcionClient) InsertInscr(inscripcion model.Users_x_courses) (model.Users_x_courses, error) {
	result := clients.Db.Create(&inscripcion)
	if result.Error != nil {
		log.Error("Error al crear la inscripción: ", result.Error)
		return inscripcion, result.Error
	}

	log.Debug("Inscripcion creada: ", inscripcion.Users_x_courses_id)
	return inscripcion, nil
}

func (s *inscripcionClient) GetInscripcion() model.Userss_x_coursess {
	var inscripciones model.Userss_x_coursess
	clients.Db.Find(&inscripciones)

	log.Debug("Inscripciones: ", inscripciones)
	return inscripciones
}

func (s *inscripcionClient) GetCourseByUserId(id int) model.Userss_x_coursess {
	var inscripciones model.Userss_x_coursess
	clients.Db.Where("user_id = ?", id).Find(&inscripciones)

	log.Debug("Inscripciones: ", inscripciones)
	return inscripciones
}

func (s *inscripcionClient) InsertValoracion(id int, valoracion model.Users_x_courses) (model.Users_x_courses, error) {
	var inscripcion model.Users_x_courses
	// Buscar la inscripción por ID
	if err := clients.Db.Where("users_x_courses_id = ?", id).First(&inscripcion).Error; err != nil {
		log.Error("Error al encontrar la inscripción: ", err)
		return inscripcion, err
	}

	// Actualizar solo las columnas valoracion y comentario
	result := clients.Db.Model(&inscripcion).Updates(map[string]interface{}{
		"valoracion": valoracion.Valoracion,
		"comentario": valoracion.Comentario,
	})
	if result.Error != nil {
		log.Error("Error al actualizar la inscripción: ", result.Error)
		return inscripcion, result.Error
	}

	log.Debug("Inscripción actualizada: ", inscripcion.Users_x_courses_id)
	return inscripcion, nil
}

func (s *inscripcionClient) InsertArchivo(id int, archivo model.Users_x_courses) (model.Users_x_courses, error) {
	var inscripcion model.Users_x_courses
	// Buscar la inscripción por ID
	if err := clients.Db.Where("users_x_courses_id = ?", id).First(&inscripcion).Error; err != nil {
		log.Error("Error al encontrar la inscripción: ", err)
		return inscripcion, err
	}

	// Actualizar solo el archivo
	result := clients.Db.Model(&inscripcion).Updates(map[string]interface{}{
		"archivo": archivo.Archivo,
	})
	if result.Error != nil {
		log.Error("Error al actualizar la entrega: ", result.Error)
		return inscripcion, result.Error
	}

	log.Debug("Entrega actualizada: ", inscripcion.Users_x_courses_id)
	return inscripcion, nil
}

func (s *inscripcionClient) GetComentByCourseId(id int) model.Userss_x_coursess {
	var inscripciones model.Userss_x_coursess
	clients.Db.Where("course_id = ?", id).Find(&inscripciones)

	log.Debug("Inscripciones: ", inscripciones)
	return inscripciones
}
