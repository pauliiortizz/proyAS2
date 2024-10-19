package dto

type LoginResponseDto struct {
	User_id int    `json:"user_id"`
	Token   string `json:"token"`
	Admin   bool   `json:"admin"`
}
