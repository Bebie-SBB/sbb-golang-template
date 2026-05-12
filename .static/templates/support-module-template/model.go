package template

import "sbb-golang-template/pkg/models"

type Template struct {
	models.BaseModel
	Code   string `json:"code" validate:"required" gorm:"type:varchar(255);not null"`
	Name   string `json:"name" validate:"required"`
	Active *bool  `json:"active" gorm:"default:true"`
}
