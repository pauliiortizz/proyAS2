package clients

import (
	"backend/model"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func ConnectDatabase() {
	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbSchema := os.Getenv("DB_NAME")
	dsn := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	connection := fmt.Sprintf(dsn, dbUsername, dbPassword, dbHost, dbPort, dbSchema)

	var err error
	for i := 0; i < 10; i++ { // Intentar conectar hasta 10 veces
		Db, err = gorm.Open(mysql.Open(connection), &gorm.Config{})
		if err == nil {
			fmt.Println("Connection Established")
			break
		}
		fmt.Println("Error connecting to DB, retrying in 5 seconds...", err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
}

func MigrateEntities() {
	err := Db.AutoMigrate(&model.Users{}, &model.Courses{}, &model.Users_x_courses{})
	if err != nil {
		fmt.Println("Error migrating to DB", err)
		panic(err)
	}
	fmt.Println("Finishing Migration Database Tables")
}
