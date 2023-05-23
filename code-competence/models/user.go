package models

type User struct {
	Id       uint   `json:"Id" form:"Id" gorm:"primary_key"`
	Email    string `json:"Email" form:"Email" validate:"required,email"`
	Name     string `json:"Name" form:"Name" validate:"required"`
	Password string `json:"Password" form:"Password" validate:"required"`
}
type UserResponse struct {
	Id    uint
	Email string
	Name  string
}
type UserResponseLogin struct {
	Id    uint
	Email string
	Name  string
	Token string
}
