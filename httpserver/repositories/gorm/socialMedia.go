package gorm

import (
	"context"
	"time"

	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories/models"
	"gorm.io/gorm"
)

type socialMediaRepo struct {
	db *gorm.DB
}

func NewSocialMediaRepo(db *gorm.DB) repositories.SocialMediaRepo {
	return &socialMediaRepo{
		db: db,
	}
}

func (r *socialMediaRepo) CreateSocialMedia(ctx context.Context, socialMedia *models.SocialMedia) error {
	socialMedia.CreatedAt = time.Now()
	return r.db.WithContext(ctx).Preload("user_id").Create(socialMedia).Error
}

func (r *socialMediaRepo) GetSocialMedia(ctx context.Context) (*models.SocialMedia, error) {
	socialMedia := new(models.SocialMedia)
	sm := &[]models.SocialMedia{}
	return socialMedia, r.db.WithContext(ctx).Find(sm).Take(socialMedia).Error
}

func (r *socialMediaRepo) EditSocialMedia(ctx context.Context, socialMedia *models.SocialMedia) error {
	socialMedia.UpdatedAt = time.Now()
	return r.db.WithContext(ctx).Model(socialMedia).Updates(*&socialMedia).Error
}

func (r *socialMediaRepo) DeleteSocialMedia(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.SocialMedia{}, "id = ?", id).Error
}
