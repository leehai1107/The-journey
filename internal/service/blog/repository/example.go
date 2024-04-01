package repository

import (
	"gorm.io/gorm"
)

type IExampleRepo interface{}

type exampleRepo struct {
	db *gorm.DB
}

func NewExampleRepo(db *gorm.DB) IExampleRepo {
	return &exampleRepo{
		db: db,
	}
}
