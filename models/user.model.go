package models

type User struct {
	Base
	Name     string `gorm:"not null"`
	Username string `gorm:"not null;index;unique"`
	Password string `gorm:"not null"`

	Votes []Vote `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func NewUser(name, username, password string) User {
	return User{
		Name:     name,
		Username: username,
		Password: password,
	}
}
