package services

import (
	"context"
	"net/http"

	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers/params"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers/views"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories"

	// "github.com/nathanramli/hacktiv8-mygram/httpserver/repositories/gorm"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories/models"
	"gorm.io/gorm"
)

type socialMediaSvc struct {
	repo repositories.SocialMediaRepo
}

func NewSocialMediaSvc(repo repositories.SocialMediaRepo) SocialMediaSvc {
	return &socialMediaSvc{
		repo: repo,
	}
}

func (s *socialMediaSvc) CreateSocialMedia(ctx context.Context, socialMedia *params.CreateSocialMedia, UserID uint) *views.Response {
	// request
	sm := models.SocialMedia{
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserId:         int(UserID),
	}
	err := s.repo.CreateSocialMedia(ctx, &sm)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusCreated, views.M_CREATED, views.CreateSocialMedia{
		Id:             sm.Id,
		Name:           sm.Name,
		SocialMediaUrl: sm.SocialMediaUrl,
		UserId:         UserID,
		CreatedAt:      sm.CreatedAt,
	})
}
func (s *socialMediaSvc) GetSocialMedia(ctx context.Context) *views.Response {
	// user := models.User{}
	sm, err := s.repo.GetSocialMedia(ctx)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}
	socialMediaViews := make([]views.GetSocialMedia, 0)
	// userViews := make([]views.Register, 0)
	for _, v := range sm {
		socialMediaViews = append(socialMediaViews, views.GetSocialMedia{
			Id:             v.Id,
			Name:           v.Name,
			SocialMediaUrl: v.SocialMediaUrl,
			UserId:         uint(v.UserId),
			CreatedAt:      v.CreatedAt,
			UpdatedAt:      v.UpdatedAt,
			User:           v.User,
		})
	}
	return views.SuccessResponse(http.StatusOK, views.M_OK, socialMediaViews)
}

func (s *socialMediaSvc) UpdateSocialMedia(ctx context.Context, socialMedia *params.UpdateSocialMedia, id uint) *views.Response {
	model, err := s.repo.GetSocialMediaByID(ctx, int(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, err)
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}
	model.Name = socialMedia.Name
	model.SocialMediaUrl = socialMedia.SocialMediaUrl
	model.UserId = int(id)
	// sm := models.SocialMedia{
	// 	Name:           socialMedia.Name,
	// 	SocialMediaUrl: socialMedia.SocialMediaUrl,
	// 	UserId:         int(id),
	// }
	// sm.Id = id

	if err = s.repo.EditSocialMedia(ctx, model); err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}
	return views.SuccessResponse(http.StatusOK, views.M_OK, views.EditSocialMedia{
		Id:             model.Id,
		Name:           model.Name,
		SocialMediaUrl: model.SocialMediaUrl,
		UserId:         uint(model.UserId),
		UpdatedAt:      model.UpdatedAt,
	})
}

func (s *socialMediaSvc) DeleteSocialMedia(ctx context.Context, id uint) (*views.Response, error) {
	// err := s.repo.DeleteSocialMedia(ctx, id)
	_, err := s.repo.GetSocialMediaByID(ctx, int(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, err), err
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err), err
	}
	if err = s.repo.DeleteSocialMedia(ctx, id); err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err), err
	}

	return views.SuccessResponse(http.StatusOK, views.M_ACCOUNT_SUCCESSFULLY_DELETED, nil), nil
}
