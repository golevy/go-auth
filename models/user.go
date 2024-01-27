package models

type User struct {
	Id       uint // Standard field for the primary key
	Name     string
	Email    string `gorm:"unique"` // specifies column as unique
	Password string
}
