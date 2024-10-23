package domain_users

type TokenDto struct {
	Token   string `json:"token"`
	User_id int64  `json:"id_user"`
	Admin   bool   `json:"admin"`
}
