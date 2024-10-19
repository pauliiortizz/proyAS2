package dto

type TokenDto struct {
	Token   string `json:"token"`
	User_id int    `json:"id_user"`
	Admin   bool   `json:"admin"`
}
