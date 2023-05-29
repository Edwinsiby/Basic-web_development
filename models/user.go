package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       int
	Role     string `gorm:"not null;default:user"`
	Name     string
	Email    string
	Password string
}
type Invalid struct {
	Errpass string
	Errmail string
}
