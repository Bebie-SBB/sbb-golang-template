package models

import (
	"sbb-golang-template/pkg/logger"

	"gorm.io/gorm"
)

type Module[H any, S any] struct {
	ModuleName string
	Handler    H
	Service    S
	Log        *logger.Logger
}

type Service[R any] struct {
	Repository R
	Log        *logger.Logger
}

type Handler[S any] struct {
	Service S
}

type Repository struct {
	DB *gorm.DB
}
