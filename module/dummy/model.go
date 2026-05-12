package dummy

import (
	"sbb-golang-template/pkg/models"
)

type Dummy struct {
	models.BaseModel
	Code   string `json:"code" validate:"required" gorm:"type:varchar(255);not null;index:idx_dummy_code_unique,unique;<-:create"`
	Name   string `json:"name" validate:"required"`
	Active *bool  `json:"active" gorm:"default:true"`
}

type SearchQuery struct {
	Keyword string `json:"keyword" validate:"required"`
	Limit   string
	Page    string
}

type CreateArgs struct {
	Dummy
}

type UpdateArgs struct {
	Dummy
}

type DeleteArgs struct {
	Dummy
}

type ReadArgs struct {
	Dummy
	AllowNotFoundFlg bool
}
