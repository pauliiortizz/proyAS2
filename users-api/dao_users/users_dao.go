package dao_users

type User struct {
	User_id  string `bson:"user_id"`
	Username string `bson:"username"`
	Password string `bson:"password"`
	Nombre   string `bson:"nombre"`
	Apellido string `bson:"apellido"`
	Email    string `bson:"email"`
	Admin    bool   `bson:"admin"`
}

type Users []User