package services

import (
	"context"
	"net/http"

	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers/params"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers/views"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories/models"
)

type socialMediaSvc struct {
	repo repositories.SocialMediaRepo
}

func NewSocialMediaSvc(repo repositories.SocialMediaRepo) SocialMediaSvc {
	return &socialMediaSvc{
		repo: repo,
	}
}

func (s *socialMediaSvc) CreateSocialMedia(ctx context.Context, socialMedia *params.CreateSocialMedia) *views.Response {
	// request
	sm := models.SocialMedia{
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
	}
	err := s.repo.CreateSocialMedia(ctx, &sm)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusCreated, views.M_CREATED, views.CreateSocialMedia{
		Id:             sm.Id,
		Name:           sm.Name,
		SocialMediaUrl: sm.SocialMediaUrl,
		UserId:         sm.UserId,
		CreatedAt:      sm.CreatedAt,
	})
}
func (s *socialMediaSvc) GetSocialMedia(ctx context.Context) *views.Response {
	sm := models.SocialMedia{}
	// user := models.User{}
	// err := s.repo.GetSocialMedia(ctx)
	// if err != nil {
	// 	return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	// }
	return views.SuccessResponse(http.StatusOK, views.M_OK, views.GetSocialMedia{
		Id:             sm.Id,
		Name:           sm.Name,
		SocialMediaUrl: sm.SocialMediaUrl,
		UserId:         sm.UserId,
		CreatedAt:      sm.CreatedAt,
		UpdatedAt:      sm.UpdatedAt,
	})
}

func (s *socialMediaSvc) UpdateSocialMedia(ctx context.Context, socialMedia *params.UpdateSocialMedia, id uint) *views.Response {
	sm := models.SocialMedia{
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
	}
	sm.Id = id
	err := s.repo.EditSocialMedia(ctx, &sm)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}
	return views.SuccessResponse(http.StatusOK, views.M_OK, views.EditSocialMedia{
		Id:             sm.Id,
		Name:           sm.Name,
		SocialMediaUrl: sm.SocialMediaUrl,
		UserId:         sm.UserId,
		UpdatedAt:      sm.UpdatedAt,
	})
}
