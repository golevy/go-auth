package models

type User struct {
	Id       uint   `json:"id"` // Standard field for the primary key
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"` // Specifies column as unique
	Password []byte `json:"-"`                   // Ignored in JSON serialization/deserialization.
}
