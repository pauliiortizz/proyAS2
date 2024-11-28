package repositories_users

import (
	"github.com/stretchr/testify/mock"
	dao "users/dao_users"
)

// Mock the Repository and Tokenizer interfaces
type Mock struct {
	mock.Mock
}

func NewMock() *Mock {
	return &Mock{}
}

func (m *Mock) GetUserById(id int64) (dao.User, error) {
	args := m.Called(id)
	if err := args.Error(1); err != nil {
		return dao.User{}, err // Return zero User if there's an error
	}
	return args.Get(0).(dao.User), nil
}

func (m *Mock) CreateUser(user dao.User) (int64, error) {
	args := m.Called(user)
	if err := args.Error(1); err != nil {
		return 0, err // Return 0 if there's an error
	}
	return args.Get(0).(int64), nil
}

func (m *Mock) GetUserByEmail(Email string) (dao.User, error) {
	args := m.Called(Email)
	if err := args.Error(1); err != nil {
		return dao.User{}, err // Return zero User if there's an error
	}
	return args.Get(0).(dao.User), nil
}
