package clients_archivo
import (
	"archivos-api/clients_archivo"
	"archivos-api/dao_archivos"
	log "github.com/sirupsen/logrus"
)

type archivoClient struct{}

type ArchivoClientInterface interface {
	InsertArchivo(id int, archivo dao_archivos.Archivos) (dao_archivos.Archivos, error)
}

var (
	ArchivoClient ArchivoClientInterface
)

func init() {
	ArchivoClient = &archivoClient{}
}


func (s *archivoClient) InsertArchivo(id int, archivo dao_archivos.Archivos) (dao_archivos.Archivos, error) {
	var archivo dao_archivos.Archivos
	// Buscar la inscripción por ID
	if err := clients.Db.Where("users_x_courses_id = ?", id).First(&inscripcion).Error; err != nil { //cambiar consulta
		log.Error("Error al encontrar la inscripción: ", err)
		return archivo, err
	}

	// Actualizar solo el archivo
	result := clients.Db.Model(&archivos).Updates(map[string]interface{}{
		"archivo": archivo.Archivo,
	})
	if result.Error != nil {
		log.Error("Error al actualizar la carga: ", result.Error)
		return archivo, result.Error
	}

	log.Debug("Carga actualizada: ", inscripcion.Users_x_courses_id)
	return archivo, nil
}


