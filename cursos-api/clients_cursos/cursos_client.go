package clients_cursos

import (
"cursos-api/clients_cursos"
"cursos-api/dao_cursos"
//"fmt"
log "github.com/sirupsen/logrus"
)

type coursesClient struct{}

type CoursesClientInterface interface {
	InsertCourse(courses dao_cursos.Curso) (dao_cursos.Curso, error)
	EditCourse(id int, updatedCourse dao_cursos.Curso) (dao_cursos.Curso, error)
	DeleteCourse(id int) error
}

var (
	CoursesClient CoursesClientInterface
)

func init() {
	CoursesClient = &coursesClient{}
}

func (s *coursesClient) InsertCourse(courses dao_cursos.Curso) (dao_cursos.Curso, error) {
	result := clients.Db.Create(&courses) //adaptar

	if result.Error != nil {
		log.Debug(result.Error, courses)
		courses.Course_id = 0
		return courses, result.Error
	}
	log.Debug("Curso creado: ", courses.Course_id)
	return courses, nil
}

func (s *coursesClient) EditCourse(id int, updatedCourse dao_cursos.Curso) (dao_cursos.Curso, error) {
	var course dao_cursos.Curso
	if err := clients.Db.Where("course_id = ?", id).First(&course).Error; err != nil  { //adaptar
		log.Debug(err, id)
		return course, err
	}

	result := clients.Db.Model(&course).Updates(updatedCourse) //adaptar
	if result.Error != nil {
		log.Debug(result.Error, id)
		return course, result.Error
	}

	log.Debug("Curso editado: ", id)
	return course, nil
}

func (s *coursesClient) DeleteCourse(id int) error {
	result := clients.Db.Delete(&model.Courses{}, "course_id = ?", id) //adaptar
	if result.Error != nil {
		log.Error("Error al eliminar el curso: ", result.Error)
		return result.Error
	}
	log.Debug("Curso eliminado: ", id)
	return nil
}
