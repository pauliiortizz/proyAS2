package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
	"users/config"
	controllers "users/controllers_users"
	repositories "users/repositories_users"
	services "users/services_users"
	"users/tokenizers"
)

func main() {
	// Repositories
	mySQLRepository := repositories.NewMySQL(
		repositories.MySQLConfig{
			Host:     config.MySQLHost,
			Port:     config.MySQLPort,
			Database: config.MySQLDatabase,
			Username: config.MySQLUsername,
			Password: config.MySQLPassword,
		},
	)

	cacheRepository := repositories.NewCache(repositories.CacheConfig{
		TTL: config.CacheDuration,
	})

	memcachedRepository := repositories.NewMemcached(repositories.MemcachedConfig{
		Host: config.MemcachedHost,
		Port: config.MemcachedPort,
	})

	// Tokenizer
	jwtTokenizer := tokenizers.NewTokenizer(
		tokenizers.JWTConfig{
			Key:      config.JWTKey,
			Duration: config.JWTDuration,
		},
	)

	// Services
	service := services.NewService(mySQLRepository, cacheRepository, memcachedRepository, jwtTokenizer)
	//Cannot use 'mySQLRepository' (type MySQL) as the type RepositoryType does not implement
	//'Repository' as some methods are missing:
	//GetUserById(id int64) (dao.User, errores.ApiError)
	//Login(loginDto domain.Login) (domain.TokenDto, errores.ApiError)

	// Handlers
	controller := controllers.NewController(service)

	// Create router
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// URL mappings

	router.GET("/users/:id", controller.GetUserById)
	router.POST("/createUser", controller.CreateUser)
	router.POST("/login", controller.Login)

	// Run application
	if err := router.Run(":8080"); err != nil {
		log.Panicf("Error running application: %v", err)
	}
}
