package services

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/nathanramli/hacktiv8-mygram/common"
	"github.com/nathanramli/hacktiv8-mygram/config"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers/params"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/controllers/views"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories"
	"github.com/nathanramli/hacktiv8-mygram/httpserver/repositories/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

type userSvc struct {
	repo repositories.UserRepo
}

func NewUserSvc(repo repositories.UserRepo) UserSvc {
	return &userSvc{
		repo: repo,
	}
}

func (s *userSvc) Register(ctx context.Context, user *params.Register) *views.Response {
	_, err := s.repo.FindUserByEmail(ctx, user.Email)
	if err == nil {
		return views.ErrorResponse(http.StatusBadRequest, views.M_EMAIL_ALREADY_USED, errors.New("email already used"))
	} else if err != nil && err != gorm.ErrRecordNotFound {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	_, err = s.repo.FindUserByUsername(ctx, user.Username)
	if err == nil {
		return views.ErrorResponse(http.StatusBadRequest, views.M_USERNAME_ALREADY_USED, errors.New("username already used"))
	} else if err != nil && err != gorm.ErrRecordNotFound {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	model := models.User{
		Age:      user.Age,
		Username: user.Username,
		Password: string(hashedPassword),
		Email:    user.Email,
	}
	err = s.repo.CreateUser(ctx, &model)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusCreated, views.M_CREATED, views.Register{
		Age:      model.Age,
		Id:       model.Id,
		Username: model.Username,
		Email:    model.Email,
	})
}

func (s *userSvc) Login(ctx context.Context, user *params.Login) *views.Response {
	model, err := s.repo.FindUserByEmail(ctx, user.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_INVALID_CREDENTIALS, err)
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(user.Password))
	if err != nil {
		return views.ErrorResponse(http.StatusBadRequest, views.M_INVALID_CREDENTIALS, err)
	}

	claims := &common.CustomClaims{
		Id: model.Id,
	}
	claims.ExpiresAt = time.Now().Add(time.Minute * time.Duration(config.GetJwtExpiredTime())).Unix()
	claims.Issuer = model.Username
	claims.Subject = model.Email

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(config.GetJwtSignature())

	return views.SuccessResponse(http.StatusOK, views.M_OK, views.Login{
		Token: ss,
	})
}

func (s *userSvc) UpdateUser(ctx context.Context, id int, params *params.UpdateUser) *views.Response {
	model, err := s.repo.FindUserByID(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, err)
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	if strings.ToLower(model.Email) != strings.ToLower(params.Email) {
		_, err = s.repo.FindUserByEmail(ctx, params.Email)
		if err == nil {
			return views.ErrorResponse(http.StatusBadRequest, views.M_EMAIL_ALREADY_USED, errors.New("email already used"))
		} else if err != nil && err != gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
		}
	}

	if strings.ToLower(model.Username) != strings.ToLower(params.Username) {
		_, err = s.repo.FindUserByUsername(ctx, params.Username)
		if err == nil {
			return views.ErrorResponse(http.StatusBadRequest, views.M_USERNAME_ALREADY_USED, errors.New("username already used"))
		} else if err != nil && err != gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
		}
	}

	model.Email = params.Email
	model.Username = params.Username
	if err = s.repo.UpdateUser(ctx, model); err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusOK, views.M_OK, views.UpdateUser{
		Id:        model.Id,
		Email:     model.Email,
		Username:  model.Username,
		Age:       model.Age,
		UpdatedAt: model.UpdatedAt,
	})
}

func (s *userSvc) DeleteUser(ctx context.Context, id int) *views.Response {
	_, err := s.repo.FindUserByID(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, err)
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	if err = s.repo.DeleteUser(ctx, id); err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusOK, views.M_ACCOUNT_SUCCESSFULLY_DELETED, nil)
}
