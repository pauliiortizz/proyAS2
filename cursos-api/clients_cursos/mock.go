package clients_cursos

import "cursos/domain_cursos"

type Mock struct{}

func NewMock() Mock {
	return Mock{}
}

func (Mock) Publish(cursesNew domain_cursos.CourseNew) error {
	return nil
}
