package models

type Category struct {
	Id    uint    `json:"Id" form:"Id" gorm:"primary_key"`
	Name  string  `json:"Name" form:"Name"`
	Items []Items `gorm:"foreignKey:CategoryId"`
}
type CategoryResponse struct {
	Id   uint
	Name string
}
