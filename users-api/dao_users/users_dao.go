package users

type User struct {
	ID       int64
	Username string
	Password string
	Nombre   string
	Apellido string
	Email    string
	Admin    bool
}
