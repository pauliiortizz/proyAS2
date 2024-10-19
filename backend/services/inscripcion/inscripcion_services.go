package inscripcion

import (
	inscripcionClient "backend/clients/inscripcion"
	"backend/clients/users"
	"backend/dto"
	errores "backend/extras"
	"backend/model"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type inscripService struct {
	inscripClient inscripcionClient.InscripcionClientInterface
}

type inscrServiceInterface interface {
	InsertInscr(inscripDto dto.InscripcionDto) (dto.InscripcionDto, errores.ApiError)
	GetInscripcion() ([]dto.InscripcionDto, errores.ApiError)
	GetCourseByUserId(inscripcionDto dto.InscripcionDto) ([]dto.InscripcionDto, errores.ApiError)
	InsertValoracion(id int, valoracionDto dto.ValoracionDto) (dto.ValoracionDto, errores.ApiError)
	InsertArchivo(id int, archivoDto dto.ArchivoDto) (dto.ArchivoDto, errores.ApiError)
	GetComentByCourseId(valoracionDto dto.ValoracionDto) ([]dto.ValoracionDto, errores.ApiError)
}

var (
	InscripcionService inscrServiceInterface
)

func initInscrpService(inscripcionesClient inscripcionClient.InscripcionClientInterface) inscrServiceInterface {
	inscripcion := new(inscripService)
	inscripcion.inscripClient = inscripcionesClient
	return inscripcion
}

func init() {
	InscripcionService = initInscrpService(inscripcionClient.InscripcionClient)
}

func (s *inscripService) InsertInscr(inscripDto dto.InscripcionDto) (dto.InscripcionDto, errores.ApiError) {
	var inscripcion model.Users_x_courses

	// Asigna valores desde inscripDto a inscripcion
	inscripcion.Users_x_courses_id = inscripDto.Id_inscripcion
	inscripcion.User_id = inscripDto.Id_user
	inscripcion.Course_id = inscripDto.Id_course
	inscripcion.Fecha_inscripcion = time.Now()

	// Verifica si el usuario existe
	_, er := users.UserClient.GetUserById(inscripDto.Id_user)
	if er != nil {
		// Verificar si el error es un NotFoundApiError
		if apiErr, ok := er.(errores.ApiError); ok && apiErr.Status() == http.StatusNotFound {
			return inscripDto, apiErr
		}
		return inscripDto, errores.NewInternalServerApiError("Error al obtener usuario", er)
	}

	// Inserta la inscripción en la base de datos
	inscripcion, err := s.inscripClient.InsertInscr(inscripcion)
	if err != nil {
		return inscripDto, errores.NewInternalServerApiError("Error al insertar inscripción", err)
	}

	// Prepara el objeto de respuesta
	var inscResponseDto dto.InscripcionDto
	inscResponseDto.Id_inscripcion = inscripcion.Users_x_courses_id
	inscResponseDto.Id_user = inscripcion.User_id
	inscResponseDto.Id_course = inscripcion.Course_id
	inscResponseDto.Fecha_inscripcion = inscripcion.Fecha_inscripcion

	return inscResponseDto, nil
}

func (s *inscripService) GetInscripcion() ([]dto.InscripcionDto, errores.ApiError) {
	inscripciones := s.inscripClient.GetInscripcion()
	var inscripcionesDto []dto.InscripcionDto

	for _, inscripcion := range inscripciones {
		var inscDto dto.InscripcionDto
		inscDto.Id_inscripcion = inscripcion.Users_x_courses_id
		inscDto.Id_user = inscripcion.User_id
		inscDto.Id_course = inscripcion.Course_id
		inscDto.Fecha_inscripcion = inscripcion.Fecha_inscripcion

		inscripcionesDto = append(inscripcionesDto, inscDto)
	}

	log.Debug(inscripcionesDto)
	return inscripcionesDto, nil
}

func (s *inscripService) GetCourseByUserId(inscripcionDto dto.InscripcionDto) ([]dto.InscripcionDto, errores.ApiError) {
	inscripciones := s.inscripClient.GetCourseByUserId(inscripcionDto.Id_user)
	var inscripcionesDto []dto.InscripcionDto

	for _, inscripcion := range inscripciones {
		var inscDto dto.InscripcionDto
		inscDto.Id_inscripcion = inscripcion.Users_x_courses_id
		inscDto.Id_user = inscripcion.User_id
		inscDto.Id_course = inscripcion.Course_id
		inscDto.Fecha_inscripcion = inscripcion.Fecha_inscripcion

		inscripcionesDto = append(inscripcionesDto, inscDto)
	}

	log.Debug(inscripcionDto)
	return inscripcionesDto, nil
}

func (s *inscripService) InsertValoracion(id int, valoracionDto dto.ValoracionDto) (dto.ValoracionDto, errores.ApiError) {
	var valoracion model.Users_x_courses

	valoracion.Users_x_courses_id = id
	valoracion.User_id = valoracionDto.Id_user
	valoracion.Course_id = valoracionDto.Id_course
	valoracion.Comentario = valoracionDto.Comentario
	valoracion.Valoracion = valoracionDto.Valoracion

	valoracion, err := s.inscripClient.InsertValoracion(id, valoracion)
	if err != nil {
		return dto.ValoracionDto{}, errores.NewInternalServerApiError("Error al agregar comentario y valoración", err)
	}

	var valoracionResponse dto.ValoracionDto
	valoracionResponse.Id_inscripcion = valoracion.Users_x_courses_id
	valoracionResponse.Id_user = valoracion.User_id
	valoracionResponse.Id_course = valoracion.Course_id
	valoracionResponse.Comentario = valoracion.Comentario
	valoracionResponse.Valoracion = valoracion.Valoracion

	return valoracionResponse, nil
}

func (s *inscripService) InsertArchivo(id int, archivoDto dto.ArchivoDto) (dto.ArchivoDto, errores.ApiError) {
	var archivo model.Users_x_courses

	archivo.Users_x_courses_id = id // Aquí utilizamos el id directamente en lugar de archivoDto.Id_inscripcion
	archivo.User_id = archivoDto.Id_user
	archivo.Course_id = archivoDto.Id_course
	archivo.Archivo = archivoDto.Archivo

	archivo, err := s.inscripClient.InsertArchivo(id, archivo)
	if err != nil {
		return dto.ArchivoDto{}, errores.NewInternalServerApiError("Error al agregar el archivo", err)
	}

	var archivoResponse dto.ArchivoDto
	archivoResponse.Id_inscripcion = archivo.Users_x_courses_id
	archivoResponse.Id_user = archivo.User_id
	archivoResponse.Id_course = archivo.Course_id
	archivoResponse.Archivo = archivo.Archivo

	return archivoResponse, nil
}

func (s *inscripService) GetComentByCourseId(valoracionDto dto.ValoracionDto) ([]dto.ValoracionDto, errores.ApiError) {
	inscripciones := s.inscripClient.GetComentByCourseId(valoracionDto.Id_course)
	var valoracionesDto []dto.ValoracionDto

	for _, inscripcion := range inscripciones {
		var valDto dto.ValoracionDto
		valDto.Id_inscripcion = inscripcion.Users_x_courses_id
		valDto.Id_user = inscripcion.User_id
		valDto.Id_course = inscripcion.Course_id
		valDto.Comentario = inscripcion.Comentario
		valDto.Valoracion = inscripcion.Valoracion

		valoracionesDto = append(valoracionesDto, valDto)
	}

	log.Debug(valoracionDto)
	return valoracionesDto, nil
}
