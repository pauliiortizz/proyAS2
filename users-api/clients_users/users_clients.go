package clients_users

import (
	log "github.com/sirupsen/logrus"
	dao_users "users/dao_users"
)

type userClient struct{}

type UserClientInterface interface {
	GetUserById(id int) (dao_users.User, error)
	GetUserByEmail(email string) (dao_users.User, error)
	CreateUser(user dao_users.User) (dao_users.User, error)
}

var (
	UserClient UserClientInterface
)

func init() {
	UserClient = &userClient{}
}

func (s *userClient) GetUserById(id int) (dao_users.User, error) {
	var user dao_users.User
	clients.Db.Where("user_id = ?", id).First(&user) //adaptar
	log.Debug("user: ", user)
	return user, nil
}

func (s *userClient) GetUserByEmail(Email string) (dao_users.User, error) {
	var user dao_users.User
	result := clients.Db.Where("email = ?", Email).First(&user) //adaptar
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (s *userClient) CreateUser(user dao_users.User) (dao_users.User, error) {
	result := clients.Db.Create(&user) //adaptar

	if result.Error != nil {
		log.Debug(result.Error, user)
		user.User_id = 0
		return user, result.Error
	}
	log.Debug("Usuario creado: ", user.User_id)
	return user, nil
}
