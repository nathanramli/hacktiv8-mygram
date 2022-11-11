package services

import (
	"context"
	"net/http"

	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers/params"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers/views"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories/models"
	"gorm.io/gorm"
)

type photoSvc struct {
	repo repositories.PhotoRepo
	user repositories.UserRepo
}

func NewPhotoSvc(repo repositories.PhotoRepo, user repositories.UserRepo) PhotoSvc {
	return &photoSvc{
		repo: repo,
		user: user,
	}
}

func (s *photoSvc) CreatePhoto(ctx context.Context, photo *params.CreatePhoto, UserID int) *views.Response {
	//request
	p := models.Photo{
		Title:	  photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
		UserId: UserID,
	}
	// p.UserId =idUser

	err := s.repo.CreatePhoto(ctx, &p)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusCreated, views.M_CREATED, views.CreatePhoto{
		Id:             p.Id,
		Title:          p.Title,
		Caption: 		p.Caption,
		PhotoUrl:       p.PhotoUrl,
		UserId:			p.UserId,
		CreatedAt:      p.CreatedAt,
	})
}

func (s *photoSvc) GetPhotos(ctx context.Context) *views.Response {
	p, err := s.repo.GetPhotos(ctx)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}
  
	photoViews := make([]views.GetPhotos, 0)
	// photoViews := make([]views.GetSocialMedia, 0)
	for _, v := range p {
		u, err := s.user.FindUserByID(ctx, v.UserId)
		if err != nil {
			return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
		}

		photoView := views.UserGetPhoto{
			Email: u.Email,
			Username: u.Username,
		}
		photoViews = append(photoViews, views.GetPhotos{
			Id:             v.Id,
			Title:           v.Title,
			Caption: v.Caption,
			PhotoUrl: v.PhotoUrl,
			UserId:         v.UserId,
			CreatedAt:      v.CreatedAt,
			UpdatedAt:      v.UpdatedAt,
			User:	photoView,
		})
		}
		return views.SuccessResponse(http.StatusOK, views.M_OK, photoViews)
	}

func (s *photoSvc) GetPhotoByID(ctx context.Context, id int) (*models.Photo, error){
	p, err := s.repo.FindPhotoByID(ctx, id)
	if err != nil {
		return p, err
	}

	if p.Id == 0 {
		return p, nil
	}

	return p, nil
}


func (s *photoSvc) UpdatePhoto(ctx context.Context, photo *params.UpdatePhoto, id int) *views.Response {
	p, err := s.repo.FindPhotoByID(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, err)
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	//request
	pReq := models.Photo{
		Title:	  photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
	}

	err = s.repo.UpdatePhoto(ctx, &pReq, id)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusOK, views.M_OK, views.UpdatePhoto{
		Id:             p.Id,
		Title:          p.Title,
		Caption: 		p.Caption,
		PhotoUrl:       p.PhotoUrl,
		UserId:			p.UserId,
		UpdatedAt:      p.UpdatedAt,
	})
}

func (s *photoSvc) DeletePhoto(ctx context.Context, id int) *views.Response {
	_, err := s.repo.FindPhotoByID(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, err)
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	if err = s.repo.DeletePhoto(ctx, id); err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusOK, views.M_PHOTO_SUCCESSFULLY_DELETED, nil)
}

