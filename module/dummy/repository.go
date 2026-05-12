package dummy

import (
	"fmt"
	"strings"

	"sbb-golang-template/pkg/database"
	"sbb-golang-template/pkg/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type dummyServiceRepository models.Repository

func (r *dummyServiceRepository) GetDatabase() (*gorm.DB, error) {
	dbService := database.GetDatabaseService()
	if dbService == nil {
		return nil, fmt.Errorf("no db")
	}
	if dbService.DefaultDatabase() == nil {
		return nil, fmt.Errorf("no default db")
	}
	return dbService.DefaultDatabase(), nil
}

func (r *dummyServiceRepository) Create(dummy *Dummy) (*Dummy, error) {
	if err := r.DB.Model(&Dummy{}).Create(dummy).Error; err != nil {
		return nil, err
	}
	return dummy, nil
}

func (r *dummyServiceRepository) Get(dummy *Dummy) (*Dummy, error) {
	result := &Dummy{}
	err := r.DB.Model(&Dummy{}).Where(&Dummy{Code: dummy.Code}).First(result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *dummyServiceRepository) List(keyword string, limit int, offset int) ([]*Dummy, error) {
	dummyList := []*Dummy{}
	queryParts := []string{}

	if keyword != "" {
		queryParts = append(queryParts, "name LIKE '%"+keyword+"%'")
		queryParts = append(queryParts, "code LIKE '%"+keyword+"%'")
	}

	orderBy := clause.OrderByColumn{Column: clause.Column{Name: "id"}, Desc: false}
	query := r.DB.Model(&Dummy{}).
		Where(strings.Join(queryParts, " OR ")).
		Order(orderBy).Limit(limit).Offset(offset).Find(&dummyList)

	return dummyList, query.Error
}

func (r *dummyServiceRepository) ListAll() (*[]Dummy, error) {
	dummyList := &[]Dummy{}
	active := true
	orderBy := clause.OrderByColumn{Column: clause.Column{Name: "id"}, Desc: false}
	query := r.DB.Model(&Dummy{}).Where(&Dummy{Active: &active}).Order(orderBy).Find(dummyList)
	return dummyList, query.Error
}

func (r *dummyServiceRepository) Update(dummy *Dummy) error {
	return r.DB.Model(&Dummy{}).Where(&Dummy{Code: dummy.Code}).Updates(dummy).Error
}

func (r *dummyServiceRepository) Delete(dummy *Dummy) error {
	return r.DB.Model(&Dummy{}).Where(&Dummy{Code: dummy.Code}).Delete(dummy).Error
}

func (r *dummyServiceRepository) CreateRepository() (*dummyServiceRepository, error) {
	r = &dummyServiceRepository{}
	db, err := r.GetDatabase()
	if err != nil {
		return nil, err
	}
	r.DB = db.Session(&gorm.Session{FullSaveAssociations: true})
	return r, r.DB.Error
}
