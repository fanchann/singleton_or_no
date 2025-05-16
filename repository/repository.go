package repository

import (
	"context"

	"gorm.io/gorm"

	"singleton_or_no/config"
	"singleton_or_no/models"
)

type Repository interface {
	All(ctx context.Context)
}

type singletonRepo struct {
	db *gorm.DB
}

func NewSingletonRepo(db *gorm.DB) Repository {
	return &singletonRepo{db: db}
}

func (r *singletonRepo) All(ctx context.Context) {
	var products []models.Product
	_ = r.db.WithContext(ctx).Find(&products).Error
}

type nonSingletonRepo struct{}

func NewNonSingletonRepo() Repository {
	return &nonSingletonRepo{}
}

func (r *nonSingletonRepo) All(ctx context.Context) {
	var products []models.Product
	_ = config.GormConnectDB().WithContext(ctx).Find(&products).Error
}
