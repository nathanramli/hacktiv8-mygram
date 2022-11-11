package gorm

import (
	"context"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories/models"
	"gorm.io/gorm"
	"time"
	"log"
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

func (r *commentRepo) FindCommentByID(ctx context.Context, id int) (*models.Comment, error) {
	comment := new(models.Comment)
	err := r.db.WithContext(ctx).Where("id = ?", id).Take(comment).Error
	return comment, err
}

func (r *commentRepo) UpdateComment(ctx context.Context, comment *models.Comment, id int) error {
	comment.UpdatedAt = time.Now()
	err := r.db.WithContext(ctx).Model(comment).Where("id = ?", id).Updates(*comment).Error
	return err
}

func (r *commentRepo) DeleteComment(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&models.Comment{}, "id = ?", id).Error
}