package cursos

import (
	"backend/clients"
	"backend/model"
	//"fmt"
	log "github.com/sirupsen/logrus"
)

type coursesClient struct{}

type CoursesClientInterface interface {
	GetCourseById(id int) model.Courses
	GetCourses() model.Coursess
	GetCourseByName(query string) model.Coursess
	InsertCourse(courses model.Courses) (model.Courses, error)
	EditCourse(id int, updatedCourse model.Courses) (model.Courses, error)
	DeleteCourse(id int) error
}

var (
	CoursesClient CoursesClientInterface
)

func init() {
	CoursesClient = &coursesClient{}
}

func (s *coursesClient) GetCourseById(id int) model.Courses {
	var course model.Courses
	clients.Db.Where("course_id = ?", id).First(&course)
	log.Debug("Course: ", course)
	return course
}

func (s *coursesClient) GetCourses() model.Coursess {
	var courses model.Coursess
	clients.Db.Find(&courses)

	log.Debug("Courses: ", courses)

	return courses
}

func (s *coursesClient) GetCourseByName(query string) model.Coursess {
	var courses model.Coursess
	clients.Db.Where("nombre LIKE ?", "%"+query+"%").Find(&courses)
	log.Debug("Courses", courses)

	return courses
}

func (s *coursesClient) InsertCourse(courses model.Courses) (model.Courses, error) {
	result := clients.Db.Create(&courses)

	if result.Error != nil {
		log.Debug(result.Error, courses)
		courses.Course_id = 0
		return courses, result.Error
	}
	log.Debug("Curso creado: ", courses.Course_id)
	return courses, nil
}

func (s *coursesClient) EditCourse(id int, updatedCourse model.Courses) (model.Courses, error) {
	var course model.Courses
	if err := clients.Db.Where("course_id = ?", id).First(&course).Error; err != nil {
		log.Debug(err, id)
		return course, err
	}

	result := clients.Db.Model(&course).Updates(updatedCourse)
	if result.Error != nil {
		log.Debug(result.Error, id)
		return course, result.Error
	}

	log.Debug("Curso editado: ", id)
	return course, nil
}

func (s *coursesClient) DeleteCourse(id int) error {
	result := clients.Db.Delete(&model.Courses{}, "course_id = ?", id)
	if result.Error != nil {
		log.Error("Error al eliminar el curso: ", result.Error)
		return result.Error
	}
	log.Debug("Curso eliminado: ", id)
	return nil
}
