package repositories_users

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
	"log"
	users "users/dao_users"
	errores "users/extras"
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

func NewMySQL(config MySQLConfig) MySQL {
	// Build DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Username, config.Password, config.Host, config.Port, config.Database)

	// Open connection to MySQL
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect to MySQL: %s", err.Error())
	}

	// Ping the database to verify connection
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping MySQL: %s", err.Error())
	}

	return MySQL{
		db: db,
	}
}

func (repository MySQL) GetUserById(id int64) (users.User, error) {
	var user users.User
	if err := repository.db.
		QueryRow("SELECT id, email, password, nombre, apellido, admin FROM users WHERE id = ?", id).
		Scan(&user.User_id, &user.Email, &user.Password, &user.Nombre, &user.Apellido, &user.Admin); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, fmt.Errorf("user not found")
		}
		return user, fmt.Errorf("error fetching user by id: %w", err)
	}
	return user, nil
}

func (repository MySQL) GetUserByEmail(email string) (users.User, error) {
	var user users.User
	if err := repository.db.
		QueryRow("SELECT id, email, password FROM users WHERE email = ?", email).
		Scan(&user.User_id, &user.Email, &user.Password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, fmt.Errorf("user not found")
		}
		return user, fmt.Errorf("error fetching user by username: %w", err)
	}
	return user, nil
}

func (repository MySQL) CreateUser(user users.User) (int64, error) {
	result, err := repository.db.Exec("INSERT INTO users (nombre, apellido, email, password, admin) VALUES (?, ?, ?, ?, ?)", user.Nombre, user.Apellido, user.Email, user.Password, user.Admin)
	if err != nil {
		return 0, errores.NewInternalServerApiError("error creating user", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, errores.NewInternalServerApiError("error getting last insert id: %w", err)
	}
	return id, nil
}
