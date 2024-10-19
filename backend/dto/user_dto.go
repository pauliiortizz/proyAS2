package dto

type UserDto struct {
	User_id  int    `json:"user_id"`
	Email    string `json:"username"`
	Password string `json:"password"`
	Nombre   string `json:"first_name"`
	Apellido string `json:"last_name"`
	Admin    bool   `json:"admin"`
}

type UsersDto []UserDto
