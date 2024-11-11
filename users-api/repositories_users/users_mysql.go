package repositories_users

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
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
	db *gorm.DB
}

func NewMySQL(config MySQLConfig) MySQL {
	// Build DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Username, config.Password, config.Host, config.Port, config.Database)

	// Open connection to MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to MySQL: %s", err.Error())
	}

	var user users.User
	if err := db.AutoMigrate(user); err != nil {
		log.Fatalf("error running Automigrate: %s", err.Error())
	}

	return MySQL{
		db: db,
	}
}

func (repository MySQL) GetUserById(id int64) (users.User, error) {
	var user users.User

	if err := repository.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, fmt.Errorf("user not found")
		}
		return user, fmt.Errorf("error fetching user by email: %w", err)
	}

	return user, nil
}

func (repository MySQL) GetUserByEmail(email string) (users.User, error) {
	var user users.User
	if err := repository.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, fmt.Errorf("user not found")
		}
		return user, fmt.Errorf("error fetching user by email: %w", err)
	}
	return user, nil
}

func (repository MySQL) CreateUser(user users.User) (int64, error) {
	if err := repository.db.Create(&user).Error; err != nil {
		return 0, errores.NewInternalServerApiError("error creating user", err)
	}
	return user.User_id, nil
}
