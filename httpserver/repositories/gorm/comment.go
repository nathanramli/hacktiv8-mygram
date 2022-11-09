package gorm

import (
	"context"
	"log"
	"time"

	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories/models"
	"gorm.io/gorm"
)

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) repositories.CommentRepo {
	return &commentRepo{
		db: db,
	}
}

func (r *commentRepo) CreateComment(ctx context.Context, comment *models.Comment) error {
	comment.CreatedAt = time.Now()
	err := r.db.WithContext(ctx).Create(comment).Error
	return err
}

func (r *commentRepo) GetComments(ctx context.Context) ([]models.Comment, error) {
	var comments []models.Comment

	err := r.db.WithContext(ctx).Find(&comments).Error
	if err != nil {
		log.Println(err)
		return comments, err
	}
	return comments, nil

}

func (r *commentRepo) GetCommentByID(ctx context.Context, id uint) (*models.Comment, error) {
	comments := new(models.Comment)
	return comments, r.db.WithContext(ctx).Where("id = ?", id).Take(comments).Error
}

func (r *commentRepo) EditComments(ctx context.Context, comment *models.Comment) error {
	comment.UpdatedAt = time.Now()
	return r.db.WithContext(ctx).Model(comment).Updates(*comment).Error
}

func (r *commentRepo) DeleteComments(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Comment{}, "id = ?", id).Error
}
