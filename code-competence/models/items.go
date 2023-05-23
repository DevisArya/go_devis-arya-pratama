package models

type Items struct {
	Id          string `gorm:"primary_key"`
	Name        string `json:"Name" form:"Name" validate:"required"`
	Description string `json:"Description" form:"Description" validate:"required"`
	Stock       uint   `json:"Stock" form:"Stock" validate:"required"`
	Price       uint   `json:"Price" form:"Price" validate:"required"`
	CategoryId  uint   `json:"CategoryId" form:"CategoryId" validate:"required"`
}
