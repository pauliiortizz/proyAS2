package clients_search

import (
"search-api/clients_search"
"search-api/dao_search"
//"fmt"
log "github.com/sirupsen/logrus"
)

type coursesClient struct{}

type CoursesClientInterface interface {
	GetCourseById(id int) dao_search.Search
	GetCourses()  dao_search.Searchs
	GetCourseByName(query string)  dao_search.Searchs

}

var (
	CoursesClient CoursesClientInterface
)

func init() {
	CoursesClient = &coursesClient{}
}

func (s *coursesClient) GetCourseById(id int)  dao_search.Search {
	var course  dao_search.Search
	clients.Db.Where("course_id = ?", id).First(&course) //adaptar
	log.Debug("Course: ", course)
	return course
}

func (s *coursesClient) GetCourses()  dao_search.Searchs {
	var courses  dao_search.Searchs
	clients.Db.Find(&courses) //adaptar

	log.Debug("Courses: ", courses)

	return courses
}

func (s *coursesClient) GetCourseByName(query string)  dao_search.Searchs {
	var courses  dao_search.Searchs
	clients.Db.Where("nombre LIKE ?", "%"+query+"%").Find(&courses) //adaptar
	log.Debug("Courses", courses)

	return courses
}
