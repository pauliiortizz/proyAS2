package controllers_users

import (
	dto "backend/dto"
	service "backend/services/users"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func GetUserById(c *gin.Context) {
	log.Debug("User id: " + c.Param("id"))

	// Get Back User

	var userDto dto.UserDto
	id, _ := strconv.Atoi(c.Param("id"))
	userDto, err := service.UserService.GetUserById(id)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, userDto)
}

func Login(c *gin.Context) {
	var loginDto dto.LoginDto
	er := c.BindJSON(&loginDto)
	if er != nil {
		log.Error(er.Error())
		c.JSON(http.StatusBadRequest, er.Error())
		return
	}

	tokenDto, err := service.UserService.Login(loginDto)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, tokenDto)
}

func CreateUser(s *gin.Context) {
	var user dto.UserDto
	err := s.BindJSON(&user)
	if err != nil {
		log.Error("Error al parsear el JSON: ", err.Error())
		s.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, er := service.UserService.CreateUser(user)
	if er != nil {
		log.Error("Error al registrar el usuario: ", er.Error())
		s.JSON(http.StatusInternalServerError, er.Error())
		return
	}

	s.JSON(http.StatusCreated, user)
}
