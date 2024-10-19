package users

import (
	"backend/clients"
	"backend/model"
	log "github.com/sirupsen/logrus"
)

type userClient struct{}

type UserClientInterface interface {
	GetUserById(id int) (model.Users, error)
	GetUserByEmail(email string) (model.Users, error)
	CreateUser(user model.Users) (model.Users, error)
}

var (
	UserClient UserClientInterface
)

func init() {
	UserClient = &userClient{}
}

func (s *userClient) GetUserById(id int) (model.Users, error) {
	var user model.Users
	clients.Db.Where("user_id = ?", id).First(&user)
	log.Debug("user: ", user)
	return user, nil
}

func (s *userClient) GetUserByEmail(Email string) (model.Users, error) {
	var user model.Users
	result := clients.Db.Where("email = ?", Email).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (s *userClient) CreateUser(user model.Users) (model.Users, error) {
	result := clients.Db.Create(&user)

	if result.Error != nil {
		log.Debug(result.Error, user)
		user.User_id = 0
		return user, result.Error
	}
	log.Debug("Usuario creado: ", user.User_id)
	return user, nil
}
