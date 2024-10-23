package domain_users

type User struct {
	User_id  int64  `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Nombre   string `json:"first_name"`
	Apellido string `json:"last_name"`
	Admin    bool   `json:"admin"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User_id int64  `json:"user_id"`
	Token   string `json:"token"`
	Admin   bool   `json:"admin"`
}

type Token struct {
	Token   string `json:"token"`
	User_id int    `json:"id_user"`
	Admin   bool   `json:"admin"`
}
