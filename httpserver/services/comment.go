package services

import (
	"context"
	"net/http"

	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers/params"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers/views"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories/models"
	// "gorm.io/gorm"
)

type commentSvc struct {
	repo repositories.CommentRepo
	user repositories.UserRepo
	photo repositories.PhotoRepo
}

func NewCommentSvc(repo repositories.CommentRepo, user repositories.UserRepo, photo repositories.PhotoRepo) CommentSvc {
	return &commentSvc{
		repo: repo,
		user: user,
		photo: photo,
	}
}

func (s *commentSvc) CreateComment(ctx context.Context, comment *params.CreateComment, UserID int) *views.Response {
	//request
	p := models.Comment{
		Message:  comment.Message,
		PhotoId:  comment.PhotoId,
		UserId: UserID,
	}

	err := s.repo.CreateComment(ctx, &p)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusCreated, views.M_CREATED, views.CreateComment{
		Id:             p.Id,
		Message:        p.Message,
		PhotoId:        p.PhotoId,
		UserId:			p.UserId,
		CreatedAt:      p.CreatedAt,
	})
}

func (s *commentSvc) GetComments(ctx context.Context) *views.Response {
	c, err := s.repo.GetComments(ctx)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}
  
	commentViews := make([]views.GetComments, 0)
	for _, v := range c {
		u, err := s.user.FindUserByID(ctx, v.UserId)
		if err != nil {
			return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
		}

		p, err := s.photo.FindPhotoByID(ctx, v.PhotoId)
		if err != nil {
			return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
		}

		userComment := views.UserGetComment{
			Id   : u.Id,
			Email: u.Email,
			Username: u.Username,
		}

		photoComment := views.PhotoGetComment{
			Id		 : p.Id,
			Title 	 : p.Title,
			Caption	 : p.Caption,
			PhotoUrl : p.PhotoUrl,
			UserId	 : p.UserId,
		}

		commentViews = append(commentViews, views.GetComments{
			Id:             v.Id,
			Message:        v.Message,
			PhotoId:        v.PhotoId,
			UserId:         v.UserId,
			UpdatedAt:      v.UpdatedAt,
			CreatedAt:      v.CreatedAt,
			User:	userComment,
			Photo:  photoComment,
		})
		}

		return views.SuccessResponse(http.StatusOK, views.M_OK, commentViews)
}