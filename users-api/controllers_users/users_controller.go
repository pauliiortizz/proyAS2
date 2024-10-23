package controllers_users

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	domain "users/domain_users"
	errores "users/extras"
)

type Service interface {
	GetUserById(id int64) (domain.User, errores.ApiError)
	Login(email string, password string) (domain.LoginResponse, errores.ApiError)
	CreateUser(user domain.User) (domain.LoginResponse, errores.ApiError)
}

type Controller struct {
	service Service
}

func NewController(service Service) Controller {
	return Controller{
		service: service,
	}
}

func (controller Controller) GetUserById(c *gin.Context) {
	log.Debug("User id: " + c.Param("id"))

	// Get Back User

	var userDto domain.User
	id, _ := strconv.Atoi(c.Param("id"))
	userDto, err := controller.service.GetUserById(int64(id))
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, userDto)
}

func (controller Controller) Login(c *gin.Context) {
	var loginDto domain.Login
	er := c.BindJSON(&loginDto)
	if er != nil {
		log.Error(er.Error())
		c.JSON(http.StatusBadRequest, er.Error())
		return
	}

	tokenDto, err := controller.service.Login(loginDto.Email, loginDto.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, tokenDto)
}

func (controller Controller) CreateUser(c *gin.Context) {
	var user domain.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Error("Error al parsear el JSON: ", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, er := controller.service.CreateUser(user)
	if er != nil {
		log.Error("Error al registrar el usuario: ", er.Error())
		c.JSON(http.StatusInternalServerError, er.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}
