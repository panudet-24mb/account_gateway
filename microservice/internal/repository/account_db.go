package repository

import (
	"gorm.io/gorm"
)

type AccountRepositoryDB struct {
	db *gorm.DB
}

func NewAccountRepositoryDB(db *gorm.DB) AccountRepositoryDB {
	return AccountRepositoryDB{db: db}
}

func (r AccountRepositoryDB) CreateAccGroup(a AccountGroup) (*AccountGroup, error) {

	result := r.db.Create(&a)
	if result.Error != nil {
		return nil, result.Error
	}

	return &a, result.Error

}

func (r AccountRepositoryDB) Create(a Account) (*Account, error) {

	result := r.db.Create(&a)
	if result.Error != nil {
		return nil, result.Error
	}

	return &a, result.Error

}
