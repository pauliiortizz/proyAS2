package router

import (
	controllers_courses "backend/controllers/cursos"
	controllers_inscripcion "backend/controllers/inscripcion"
	controllers_users "backend/controllers/users"
)

func mapUrls() {
	//Users Mapping
	router.GET("/user/:id", controllers_users.GetUserById)
	router.POST("/login", controllers_users.Login)
	router.POST("/CreateUser", controllers_users.CreateUser)

	//Courses mapping
	router.GET("/courses/:course_id", controllers_courses.GetCourseById)
	router.GET("/search", controllers_courses.GetCourses)
	router.GET("/search/:name", controllers_courses.GetCoursesByName)
	router.POST("/createCourse", controllers_courses.InsertCourse)
	router.PUT("/edit/:course_id", controllers_courses.EditCourse)
	router.DELETE("/courses/:course_id", controllers_courses.DeleteCourse)

	//inscripcion mapping
	router.POST("/inscripcion", controllers_inscripcion.InsertInscr)
	router.GET("/inscripciones", controllers_inscripcion.GetInscripcion)
	router.GET("/inscripciones/:id_user", controllers_inscripcion.GetCourseByUserId)
	router.POST("/valorar/:id", controllers_inscripcion.InsertValoracion)
	router.POST("/archivo/:id", controllers_inscripcion.InsertArchivo)
	router.GET("/comentarios/:id_course", controllers_inscripcion.GetComentByCourseId)

}
