package models

type UserModel struct {
	Id        uint `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Password  string
}
