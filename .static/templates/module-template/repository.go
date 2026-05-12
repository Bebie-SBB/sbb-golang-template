package template

import (
	"fmt"
	"strings"

	"sbb-golang-template/pkg/database"
	"sbb-golang-template/pkg/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type templateServiceRepository models.Repository

func (r *templateServiceRepository) GetDatabase() (*gorm.DB, error) {
	dbService := database.GetDatabaseService()
	if dbService == nil {
		return nil, fmt.Errorf("no db")
	}
	return dbService.DefaultDatabase(), nil
}

func (r *templateServiceRepository) Create(t *Template) (*Template, error) {
	if err := r.DB.Model(&Template{}).Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func (r *templateServiceRepository) Get(t *Template) (*Template, error) {
	result := &Template{}
	err := r.DB.Model(&Template{}).Where(&Template{Code: t.Code}).First(result).Error
	return result, err
}

func (r *templateServiceRepository) List(keyword string, limit, offset int) ([]*Template, error) {
	list := []*Template{}
	parts := []string{}
	if keyword != "" {
		parts = append(parts, "name LIKE '%"+keyword+"%'", "code LIKE '%"+keyword+"%'")
	}
	orderBy := clause.OrderByColumn{Column: clause.Column{Name: "id"}, Desc: false}
	return list, r.DB.Model(&Template{}).Where(strings.Join(parts, " OR ")).Order(orderBy).Limit(limit).Offset(offset).Find(&list).Error
}

func (r *templateServiceRepository) ListAll() (*[]Template, error) {
	list := &[]Template{}
	active := true
	orderBy := clause.OrderByColumn{Column: clause.Column{Name: "id"}, Desc: false}
	return list, r.DB.Model(&Template{}).Where(&Template{Active: &active}).Order(orderBy).Find(list).Error
}

func (r *templateServiceRepository) Update(t *Template) error {
	return r.DB.Model(&Template{}).Where(&Template{Code: t.Code}).Updates(t).Error
}

func (r *templateServiceRepository) Delete(t *Template) error {
	return r.DB.Model(&Template{}).Where(&Template{Code: t.Code}).Delete(t).Error
}

func (r *templateServiceRepository) CreateRepository() (*templateServiceRepository, error) {
	r = &templateServiceRepository{}
	db, err := r.GetDatabase()
	if err != nil {
		return nil, err
	}
	r.DB = db.Session(&gorm.Session{FullSaveAssociations: true})
	return r, r.DB.Error
}
