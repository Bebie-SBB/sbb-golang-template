package template

import "sbb-golang-template/pkg/models"

type Template struct {
	models.BaseModel
	Code   string `json:"code" validate:"required" gorm:"type:varchar(255);not null;index:idx_template_code_unique,unique;<-:create"`
	Name   string `json:"name" validate:"required"`
	Active *bool  `json:"active" gorm:"default:true"`
}

type SearchQuery struct {
	Keyword string `json:"keyword" validate:"required"`
	Limit   string
	Page    string
}

type CreateArgs struct{ Template }
type UpdateArgs struct{ Template }
type DeleteArgs struct{ Template }
type ReadArgs struct {
	Template
	AllowNotFoundFlg bool
}
