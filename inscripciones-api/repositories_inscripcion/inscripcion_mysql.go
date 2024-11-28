package repositories_inscripcion

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	inscripcionDao "inscripciones/dao_inscripcion"
	errores "inscripciones/extras"
)

type MySQLConfig struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

type MySQL struct {
	db *gorm.DB
}

func NewMySQL(config MySQLConfig) *MySQL {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Username, config.Password, config.Host, config.Port, config.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to MySQL: %s", err.Error())
	}

	var inscripcion inscripcionDao.Inscripcion
	if err := db.AutoMigrate(inscripcion); err != nil {
		log.Fatalf("error running Automigrate: %s", err.Error())
	}

	return &MySQL{db: db}
}

// InsertInscripcion - Inserta una inscripción usando GORM
func (repository MySQL) InsertInscripcion(inscripcion inscripcionDao.Inscripcion) (int64, error) {
	if err := repository.db.Create(&inscripcion).Error; err != nil {
		return 0, errores.NewInternalServerApiError("error creating inscripcion", err)
	}
	return int64(inscripcion.Id_inscripcion), nil
}

// GetInscripcionByUserID - Busca inscripciones por ID de usuario usando GORM
func (repository *MySQL) GetInscripcionByUserID(userID int) ([]inscripcionDao.Inscripcion, error) {
	var inscripciones []inscripcionDao.Inscripcion
	if err := repository.db.Where("id_user = ?", userID).Find(&inscripciones).Error; err != nil {
		return nil, fmt.Errorf("error fetching inscripciones by user ID: %w", err)
	}
	return inscripciones, nil
}

// GetInscripcionByCourseID - Busca inscripciones por ID de curso usando GORM
func (repository *MySQL) GetInscripcionByCourseID(courseID string) ([]inscripcionDao.Inscripcion, error) {
	var inscripciones []inscripcionDao.Inscripcion
	if err := repository.db.Where("id_course = ?", courseID).Find(&inscripciones).Error; err != nil {
		return nil, fmt.Errorf("error fetching inscripciones by course ID: %w", err)
	}
	return inscripciones, nil
}
