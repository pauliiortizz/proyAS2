package users

type User struct {
	User_id  int64 `gorm:"primaryKey;autoIncrement"`
	Password string
	Nombre   string
	Apellido string
	Email    string `gorm:"not null;unique" binding:"required"`
	Admin    bool
}

type Users []User
