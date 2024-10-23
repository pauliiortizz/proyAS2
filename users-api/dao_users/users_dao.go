package users

type User struct {
	User_id  int64
	Password string
	Nombre   string
	Apellido string
	Email    string
	Admin    bool
}

type Users []User
