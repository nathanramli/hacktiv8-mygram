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
}

func NewPhotoSvc(repo repositories.PhotoRepo) PhotoSvc {
	return &photoSvc{
		repo: repo,
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

	return views.SuccessResponse(http.StatusOK, views.M_OK, p)
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
	//request
	p, err := s.repo.FindPhotoByID(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, err)
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	err = s.repo.UpdatePhoto(ctx, p)
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

