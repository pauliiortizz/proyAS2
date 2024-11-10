package repositories_inscripcion

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
	log "github.com/sirupsen/logrus"
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
	db *sql.DB
}

func NewMySQL(config MySQLConfig) *MySQL {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Username, config.Password, config.Host, config.Port, config.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect to MySQL: %s", err.Error())
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping MySQL: %s", err.Error())
	}
	return &MySQL{db: db}
}

type Repository interface {
	InsertInscripcion(inscripcion inscripcionDao.Inscripcion) (int64, error)
	GetInscripcionByUserID(userID int) ([]inscripcionDao.Inscripcion, error)
	GetInscripcionByCourseID(courseID int) ([]inscripcionDao.Inscripcion, error)
}

func (repository MySQL) InsertInscripcion(inscripcion inscripcionDao.Inscripcion) (int64, error) {
	result, err := repository.db.Exec("INSERT INTO users (id_user, id_course, fecha_inscripcion) VALUES (?, ?, ?)", inscripcion.Id_user, inscripcion.Id_course, inscripcion.Fecha_inscripcion)
	if err != nil {
		return 0, errores.NewInternalServerApiError("error creating inscripcion", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, errores.NewInternalServerApiError("error getting last insert id: %w", err)
	}
	return id, nil
}

func (repository *MySQL) GetInscripcionByUserID(userID int) ([]inscripcionDao.Inscripcion, error) {
	rows, err := repository.db.Query("SELECT id_inscripcion, id_user, id_course, fecha_inscripcion FROM inscripciones WHERE id_user = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var inscripciones []inscripcionDao.Inscripcion
	for rows.Next() {
		var inscripcion inscripcionDao.Inscripcion
		if err := rows.Scan(&inscripcion.Id_inscripcion, &inscripcion.Id_user, &inscripcion.Id_course, &inscripcion.Fecha_inscripcion); err != nil {
			return nil, err
		}
		inscripciones = append(inscripciones, inscripcion)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return inscripciones, nil
}

func (repository *MySQL) GetInscripcionByCourseID(courseID int) ([]inscripcionDao.Inscripcion, error) {
	rows, err := repository.db.Query("SELECT id_inscripcion, id_user, id_course, fecha_inscripcion FROM inscripciones WHERE id_course = ?", courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var inscripciones []inscripcionDao.Inscripcion
	for rows.Next() {
		var inscripcion inscripcionDao.Inscripcion
		if err := rows.Scan(&inscripcion.Id_inscripcion, &inscripcion.Id_user, &inscripcion.Id_course, &inscripcion.Fecha_inscripcion); err != nil {
			return nil, err
		}
		inscripciones = append(inscripciones, inscripcion)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return inscripciones, nil
}
