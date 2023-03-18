package models

import "github.com/jinzhu/gorm"

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}

var DB *gorm.DB

func (u *User) Create() *User {
	DB.NewRecord(u)
	DB.Create(&u)
	return u
}
