package users_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	dao "users/dao_users"
	domain "users/domain_users"
	repositories "users/repositories_users"
	service "users/services_users"
	"users/tokenizers"
)

var (
	// Create mocks
	mainRepo      = repositories.NewMock()
	cacheRepo     = repositories.NewMock()
	memcachedRepo = repositories.NewMock()
	tokenizer     = tokenizers.NewMock()
	usersService  = service.NewService(mainRepo, cacheRepo, memcachedRepo, tokenizer)
)

func TestService(t *testing.T) {

	t.Run("GetUserById - Success from Cache", func(t *testing.T) {
		mockUser := dao.User{User_id: 1, Email: "user1", Password: "password1"}
		cacheRepo.On("GetUserById", int64(1)).Return(mockUser, nil).Once()

		result, err := usersService.GetUserById(1)

		assert.NoError(t, err)
		assert.Equal(t, "user1", result.Email)

		mainRepo.AssertExpectations(t)
		cacheRepo.AssertExpectations(t)
		memcachedRepo.AssertExpectations(t)
	})

	t.Run("GetUserById - Not Found in Cache, Found in Memcached", func(t *testing.T) {
		mockUser := dao.User{User_id: 1, Email: "user1", Password: "password1"}
		cacheRepo.On("GetUserById", int64(1)).Return(dao.User{}, errors.New("not found")).Once()
		memcachedRepo.On("GetUserById", int64(1)).Return(mockUser, nil).Once()
		cacheRepo.On("Create", mockUser).Return(int64(1), nil).Once()

		result, err := usersService.GetUserById(1)

		assert.NoError(t, err)
		assert.Equal(t, "user1", result.Email)

		mainRepo.AssertExpectations(t)
		cacheRepo.AssertExpectations(t)
		memcachedRepo.AssertExpectations(t)
	})

	t.Run("GetUserById - Not Found in Cache or Memcached, Found in Main Repo", func(t *testing.T) {
		mockUser := dao.User{User_id: 1, Email: "user1", Password: "password1"}
		cacheRepo.On("GetUserById", int64(1)).Return(dao.User{}, errors.New("not found")).Once()
		memcachedRepo.On("GetUserById", int64(1)).Return(dao.User{}, errors.New("not found")).Once()
		mainRepo.On("GetUserById", int64(1)).Return(mockUser, nil).Once()
		cacheRepo.On("Create", mockUser).Return(int64(1), nil).Once()
		memcachedRepo.On("Create", mockUser).Return(int64(1), nil).Once()

		result, err := usersService.GetUserById(1)

		assert.NoError(t, err)
		assert.Equal(t, "user1", result.Email)

		mainRepo.AssertExpectations(t)
		cacheRepo.AssertExpectations(t)
		memcachedRepo.AssertExpectations(t)
	})

	t.Run("GetUserById - Error in Main Repo", func(t *testing.T) {
		cacheRepo.On("GetUserById", int64(1)).Return(dao.User{}, errors.New("not found")).Once()
		memcachedRepo.On("GetUserById", int64(1)).Return(dao.User{}, errors.New("not found")).Once()
		mainRepo.On("GetUserById", int64(1)).Return(dao.User{}, errors.New("db error")).Once()

		result, err := usersService.GetUserById(1)

		assert.Error(t, err)
		assert.Equal(t, "error getting user by ID: db error", err.Error())
		assert.Equal(t, domain.User{}, result)

		mainRepo.AssertExpectations(t)
		cacheRepo.AssertExpectations(t)
		memcachedRepo.AssertExpectations(t)
	})

	t.Run("Create - Success", func(t *testing.T) {
		newUser := dao.User{Email: "newuser", Password: service.Hash("password")}
		mainRepo.On("Create", newUser).Return(int64(1), nil).Once()
		newUser.User_id = 1
		cacheRepo.On("Create", newUser).Return(int64(1), nil).Once()
		memcachedRepo.On("Create", newUser).Return(int64(1), nil).Once()

		id, err := usersService.CreateUser(domain.User{Email: "newuser", Password: "password"})

		assert.NoError(t, err)
		assert.Equal(t, int64(1), id)

		mainRepo.AssertExpectations(t)
		cacheRepo.AssertExpectations(t)
		memcachedRepo.AssertExpectations(t)
	})

	t.Run("Create - Error", func(t *testing.T) {
		newUser := dao.User{Email: "newuser", Password: service.Hash("password")}
		mainRepo.On("Create", newUser).Return(int64(0), errors.New("db error")).Once()

		id, err := usersService.CreateUser(domain.User{Email: "newuser", Password: "password"})

		assert.Error(t, err)
		assert.Equal(t, int64(0), id)
		assert.Equal(t, "error creating user: db error", err.Error())

		mainRepo.AssertExpectations(t)
		cacheRepo.AssertExpectations(t)
		memcachedRepo.AssertExpectations(t)
	})

	t.Run("Login - Success", func(t *testing.T) {
		Email := "user1"
		password := "password"
		hashedPassword := service.Hash(password)

		mockUser := dao.User{User_id: 1, Email: Email, Password: hashedPassword}
		cacheRepo.On("GetByEmail", Email).Return(mockUser, nil).Once()
		tokenizer.On("GenerateToken", Email, int64(1)).Return("token", nil).Once()

		response, err := usersService.Login(Email, password)

		assert.NoError(t, err)
		assert.Equal(t, int64(1), response.User_id)
		assert.Equal(t, "token", response.Token)

		mainRepo.AssertExpectations(t)
		cacheRepo.AssertExpectations(t)
		memcachedRepo.AssertExpectations(t)
	})

	t.Run("Login - Invalid Credentials", func(t *testing.T) {
		Email := "user1"
		password := "wrongpassword"
		hashedPassword := service.Hash("password")

		mockUser := dao.User{User_id: 1, Email: Email, Password: hashedPassword}
		cacheRepo.On("GetByEmail", Email).Return(mockUser, nil).Once()

		response, err := usersService.Login(Email, password)

		assert.Error(t, err)
		assert.Equal(t, "invalid credentials", err.Error())
		assert.Equal(t, domain.LoginResponse{}, response)

		mainRepo.AssertExpectations(t)
		cacheRepo.AssertExpectations(t)
		memcachedRepo.AssertExpectations(t)
	})

	t.Run("Login - User Not Found", func(t *testing.T) {
		Email := "user1"
		password := "password"

		cacheRepo.On("GetByEmail", Email).Return(dao.User{}, errors.New("not found")).Once()
		memcachedRepo.On("GetByEmail", Email).Return(dao.User{}, errors.New("not found")).Once()
		mainRepo.On("GetByEmail", Email).Return(dao.User{}, errors.New("not found")).Once()

		response, err := usersService.Login(Email, password)

		assert.Error(t, err)
		assert.Equal(t, "error getting user by Email from main repository: not found", err.Error())
		assert.Equal(t, domain.LoginResponse{}, response)

		mainRepo.AssertExpectations(t)
		cacheRepo.AssertExpectations(t)
		memcachedRepo.AssertExpectations(t)
	})

	t.Run("Login - Token Generation Error", func(t *testing.T) {
		Email := "user1"
		password := "password"
		hashedPassword := service.Hash(password)

		mockUser := dao.User{User_id: 1, Email: Email, Password: hashedPassword}
		cacheRepo.On("GetByEmail", Email).Return(mockUser, nil).Once()
		tokenizer.On("GenerateToken", Email, int64(1)).Return("", errors.New("token error")).Once()

		response, err := usersService.Login(Email, password)

		assert.Error(t, err)
		assert.Equal(t, "error generating token: token error", err.Error())
		assert.Equal(t, domain.LoginResponse{}, response)

		mainRepo.AssertExpectations(t)
		cacheRepo.AssertExpectations(t)
		memcachedRepo.AssertExpectations(t)
	})
}
