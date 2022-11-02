package gorm

import (
	"context"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories/models"
	"gorm.io/gorm"
	"time"
	"log"
)

type photoRepo struct {
	db *gorm.DB
}

func NewPhotoRepo(db *gorm.DB) repositories.PhotoRepo {
	return &photoRepo{
		db: db,
	}
}

func (r *photoRepo) CreatePhoto(ctx context.Context, photo *models.Photo) error {
	photo.CreatedAt = time.Now()
	err := r.db.WithContext(ctx).Create(photo).Error
	return err
}

func (r *photoRepo) GetPhotos(ctx context.Context) ([]models.Photo, error) {
	var photos []models.Photo

	err := r.db.WithContext(ctx).Find(&photos).Error
	if err != nil {
		log.Println(err)
		return photos, err
	}
	return photos, nil
	
}

func (r *photoRepo) UpdatePhoto(ctx context.Context, photo *models.Photo) error {
	photo.UpdatedAt = time.Now()
	return r.db.WithContext(ctx).Model(photo).Updates(*photo).Error
}