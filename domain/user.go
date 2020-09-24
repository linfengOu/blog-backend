package domain

import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model
	Username string `gorm:"column:username"`
	Email    string `gorm:"column:email;unique_index"`
	Bio      string `gorm:"column:bio;size:1024"`
}

type UserAbstractModel struct {
	ID       uint
	Username string
}
