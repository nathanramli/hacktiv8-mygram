package services

import (
	"context"
	"net/http"

	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers/params"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers/views"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories/models"
)

type photoSvc struct {
	repo repositories.PhotoRepo
}

func NewPhotoSvc(repo repositories.PhotoRepo) PhotoSvc {
	return &photoSvc{
		repo: repo,
	}
}

func (s *photoSvc) CreatePhoto(ctx context.Context, photo *params.CreatePhoto, UserID uint) *views.Response {
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

	return views.SuccessResponse(http.StatusOK, views.M_OK, views.GetPhotos{
		Id:             p.Id,
		Title:          p.Title,
		Caption: 		p.Caption,
		PhotoUrl:       p.PhotoUrl,
		UserId:			p.UserId,
		User
	})
}

func (s *photoSvc) UpdatePhoto(ctx context.Context, photo *params.UpdatePhoto, id uint) *views.Response {
	//request
	p := models.Photo{
		Title:	  photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
	}
	p.Id = id

	err := s.repo.UpdatePhoto(ctx, &p)
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