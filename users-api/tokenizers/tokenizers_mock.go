package tokenizers

import "github.com/stretchr/testify/mock"

type Mock struct {
	mock.Mock
}

func NewMock() *Mock {
	return &Mock{}
}

func (m *Mock) GenerateToken(Email string, User_id int64) (string, error) {
	args := m.Called(Email, User_id)
	return args.String(0), args.Error(1)
}
