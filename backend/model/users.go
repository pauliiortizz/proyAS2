package model

type Users struct {
	User_id  int    `gorm:"primaryKey;AUTO_INCREMENT"`
	Email    string `gorm:"type:varchar(100);not null;unique"`
	Password string `gorm:"type:varchar(100);not null"`
	Nombre   string `gorm:"type:varchar(100);not null"`
	Apellido string `gorm:"type:varchar(100);not null"`
	Admin    bool   `gorm:"type:boolean;not null"`
}
type Userss []Users
