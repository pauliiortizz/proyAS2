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
		// cacheRepo.On("Create", mockUser).Return(int64(1), nil).Once()
		cacheRepo.On("CreateUser", mockUser).Return(int64(1), nil).Once()

		result, err := usersService.GetUserById(1)

		assert.NoError(t, err)
		assert.Equal(t, "user1", result.Email)

		mainRepo.AssertExpectations(t)
		cacheRepo.AssertExpectations(t)
		memcachedRepo.AssertExpectations(t)
	})

	t.Run("GetUserById - Not Found in Cache or Memcached, Found in Main Repo", func(t *testing.T) {
		mockUser := dao.User{User_id: 1, Email: "user1", Password: "password1"}

		// Configura los mocks para las llamadas esperadas
		cacheRepo.On("GetUserById", int64(1)).Return(dao.User{}, errors.New("not found")).Once()
		memcachedRepo.On("GetUserById", int64(1)).Return(dao.User{}, errors.New("not found")).Once()
		mainRepo.On("GetUserById", int64(1)).Return(mockUser, nil).Once()

		// Configura los mocks para las llamadas a CreateUser en cacheRepository y memcachedRepository
		cacheRepo.On("CreateUser", mockUser).Return(int64(1), nil).Once()
		memcachedRepo.On("CreateUser", mockUser).Return(int64(1), nil).Once()

		// Ejecuta la función bajo prueba
		result, err := usersService.GetUserById(1)

		// Verifica los resultados
		assert.NoError(t, err)
		assert.Equal(t, "user1", result.Email)

		// Verifica que los mocks cumplieron las expectativas
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
		//assert.Equal(t, "error getting user by ID: db error", err.Error())
		assert.Equal(t, "Message: user not found;Error Code: bad_request;Status: 400;Cause: []", err.Error())
		assert.Equal(t, domain.User{}, result)

		mainRepo.AssertExpectations(t)
		cacheRepo.AssertExpectations(t)
		memcachedRepo.AssertExpectations(t)
	})

	t.Run("Create - Success", func(t *testing.T) {
		newUser := dao.User{Email: "newuser", Password: service.Hash("password")}
		mainRepo.On("CreateUser", newUser).Return(int64(1), nil).Once()

		// Actualiza el ID del usuario después de ser creado en mainRepo
		newUser.User_id = 1

		cacheRepo.On("CreateUser", newUser).Return(int64(1), nil).Once()
		memcachedRepo.On("CreateUser", newUser).Return(int64(1), nil).Once()

		id, err := usersService.CreateUser(domain.User{Email: "newuser", Password: "password"})

		assert.NoError(t, err)
		assert.Equal(t, int64(1), id)

		mainRepo.AssertExpectations(t)
		cacheRepo.AssertExpectations(t)
		memcachedRepo.AssertExpectations(t)
	})

	t.Run("Create - Error", func(t *testing.T) {
		newUser := dao.User{Email: "newuser", Password: service.Hash("password")}

		// Configurar el mock para fallar en mainRepository
		mainRepo.On("CreateUser", newUser).Return(int64(0), errors.New("db error")).Once()

		// Configurar el mock para las llamadas adicionales (cacheRepository y memcachedRepository)
		cacheRepo.On("CreateUser", newUser).Return(int64(0), nil).Maybe()
		memcachedRepo.On("CreateUser", newUser).Return(int64(0), nil).Maybe()

		id, err := usersService.CreateUser(domain.User{Email: "newuser", Password: "password"})

		// Validar los resultados
		assert.Error(t, err)
		assert.Equal(t, int64(0), id)
		//assert.Equal(t, "error creating user: db error", err.Error())
		assert.Equal(t, "Message: error creating user;Error Code: internal_server_error;Status: 500;Cause: [db error]", err.Error())

		// Verificar que las expectativas del mock se cumplieron
		mainRepo.AssertExpectations(t)
		cacheRepo.AssertExpectations(t)
		memcachedRepo.AssertExpectations(t)
	})

	t.Run("Login - Success", func(t *testing.T) {
		email := "user1"
		password := "password"
		hashedPassword := service.Hash(password) // Generar el hash de la contraseña

		// Usuario con contraseña hasheada
		mockUser := dao.User{User_id: 1, Email: email, Password: hashedPassword}

		// Configurar los mocks para las llamadas a GetUserByEmail
		cacheRepo.On("GetUserByEmail", email).Return(dao.User{}, errors.New("not found")).Once()
		memcachedRepo.On("GetUserByEmail", email).Return(dao.User{}, errors.New("not found")).Once()
		mainRepo.On("GetUserByEmail", email).Return(mockUser, nil).Once()

		// Configurar el mock para CreateUser con el usuario correcto (incluye contraseña hasheada)
		cacheRepo.On("CreateUser", mockUser).Return(int64(1), nil).Maybe()
		memcachedRepo.On("CreateUser", mockUser).Return(int64(1), nil).Maybe()

		// Configurar el mock para la generación del token
		tokenizer.On("GenerateToken", email, int64(1)).Return("token", nil).Once()

		// Ejecutar el método bajo prueba
		response, err := usersService.Login(email, password)

		// Validar los resultados
		assert.NoError(t, err)
		assert.Equal(t, int64(1), response.User_id)
		assert.Equal(t, "token", response.Token)

		// Verificar expectativas
		cacheRepo.AssertExpectations(t)
		memcachedRepo.AssertExpectations(t)
		mainRepo.AssertExpectations(t)
		tokenizer.AssertExpectations(t)
	})

	t.Run("Login - Invalid Credentials", func(t *testing.T) {
		email := "user1"
		wrongPassword := "wrongpassword"
		hashedPassword := service.Hash("password") // Hash correcto de la contraseña original

		// Usuario con contraseña hasheada
		mockUser := dao.User{User_id: 1, Email: email, Password: hashedPassword}

		// Configurar los mocks para las llamadas a GetUserByEmail
		cacheRepo.On("GetUserByEmail", email).Return(dao.User{}, errors.New("not found")).Once()
		memcachedRepo.On("GetUserByEmail", email).Return(dao.User{}, errors.New("not found")).Once()
		mainRepo.On("GetUserByEmail", email).Return(mockUser, nil).Once()

		// Ejecutar el método bajo prueba
		response, err := usersService.Login(email, wrongPassword)

		// Validar los resultados
		assert.Error(t, err)
		assert.Equal(t, "invalid credentials", err.Error())
		assert.Equal(t, domain.LoginResponse{}, response)

		// Verificar expectativas
		cacheRepo.AssertExpectations(t)
		memcachedRepo.AssertExpectations(t)
		mainRepo.AssertExpectations(t)
	})

	t.Run("Login - User Not Found", func(t *testing.T) {
		email := "user1"
		password := "password"

		// Configurar los mocks para devolver errores en todas las llamadas a GetUserByEmail
		cacheRepo.On("GetUserByEmail", email).Return(dao.User{}, errors.New("not found")).Once()
		memcachedRepo.On("GetUserByEmail", email).Return(dao.User{}, errors.New("not found")).Once()
		mainRepo.On("GetUserByEmail", email).Return(dao.User{}, errors.New("not found")).Once()

		// Ejecutar el método bajo prueba
		response, err := usersService.Login(email, password)

		// Validar los resultados
		assert.Error(t, err)
		assert.Equal(t, "error getting user by email from main repository: not found", err.Error())
		assert.Equal(t, domain.LoginResponse{}, response)

		// Verificar expectativas
		cacheRepo.AssertExpectations(t)
		memcachedRepo.AssertExpectations(t)
		mainRepo.AssertExpectations(t)
	})

	t.Run("Login - Token Generation Error", func(t *testing.T) {
		email := "user1"
		password := "password"
		hashedPassword := service.Hash(password) // Generar el hash de la contraseña

		// Usuario con contraseña hasheada
		mockUser := dao.User{User_id: 1, Email: email, Password: hashedPassword}

		// Configurar los mocks para las llamadas a GetUserByEmail
		cacheRepo.On("GetUserByEmail", email).Return(dao.User{}, errors.New("not found")).Once()
		memcachedRepo.On("GetUserByEmail", email).Return(dao.User{}, errors.New("not found")).Once()
		mainRepo.On("GetUserByEmail", email).Return(mockUser, nil).Once()

		// Configurar el mock para la generación del token con un error
		tokenizer.On("GenerateToken", email, int64(1)).Return("", errors.New("token error")).Once()

		// Ejecutar el método bajo prueba
		response, err := usersService.Login(email, password)

		// Validar los resultados
		assert.Error(t, err)
		assert.Equal(t, "error generating token: token error", err.Error())
		assert.Equal(t, domain.LoginResponse{}, response)

		// Verificar expectativas
		cacheRepo.AssertExpectations(t)
		memcachedRepo.AssertExpectations(t)
		mainRepo.AssertExpectations(t)
		tokenizer.AssertExpectations(t)
	})
}
