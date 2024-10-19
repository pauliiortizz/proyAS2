package domain_users

type UserDto struct {
	User_id  int64  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"username"`
	Password string `json:"password"`
	Nombre   string `json:"first_name"`
	Apellido string `json:"last_name"`
	Admin    bool   `json:"admin"`
}

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseDto struct {
	User_id int64  `json:"user_id"`
	Token   string `json:"token"`
	Admin   bool   `json:"admin"`
}
