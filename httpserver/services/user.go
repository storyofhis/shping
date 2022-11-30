package services

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/storyofhis/toko-belanja/common"
	"github.com/storyofhis/toko-belanja/config"
	"github.com/storyofhis/toko-belanja/httpserver/controllers/params"
	"github.com/storyofhis/toko-belanja/httpserver/controllers/views"
	"github.com/storyofhis/toko-belanja/httpserver/repositories"
	"github.com/storyofhis/toko-belanja/httpserver/repositories/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userSvc struct {
	// db *gorm.DB
	repo repositories.UserRepo
}

func NewUserSvc(repo repositories.UserRepo) UserSvc {
	return &userSvc{
		repo: repo,
	}
}

func (svc *userSvc) Register(ctx context.Context, user *params.Register) *views.Response {
	_, err := svc.repo.FindUserByEmail(ctx, user.Email)
	if err == nil {
		return views.ErrorResponse(http.StatusBadRequest, views.M_EMAIL_ALREADY_USED, errors.New("email already used"))
	} else if err != nil && err != gorm.ErrRecordNotFound {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}
	// generate password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}
	// request from user
	input := models.Users{
		FullName: user.FullName,
		Email:    user.Email,
		Password: string(hashedPassword),
	}

	err = svc.repo.CreateUser(ctx, &input)

	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusCreated, views.M_CREATED, views.Register{
		Id:        input.Id,
		FullName:  input.FullName,
		Email:     input.Email,
		Password:  input.Password,
		Balance:   input.Balance,
		CreatedAt: input.CreatedAt,
	})
}

func (svc *userSvc) Login(ctx context.Context, user *params.Login) *views.Response {
	model, err := svc.repo.FindUserByEmail(ctx, user.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_INVALID_CREDENTIALS, err)
			// return nil
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
		// return nil
	}
	err = bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(user.Password))
	if err != nil {
		return views.ErrorResponse(http.StatusBadRequest, views.M_INVALID_CREDENTIALS, err)
		// return nil
	}

	role := string(model.Role)

	claims := &common.CustomClaims{
		Id:   int(model.Id),
		Role: role,
	}

	claims.ExpiresAt = time.Now().Add(time.Minute * time.Duration(config.GetJwtExpiredTime())).Unix()
	claims.Subject = model.Email

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(config.GetJwtSignature())

	return views.SuccessResponse(http.StatusOK, views.M_OK, views.Login{
		Token: ss,
	})
}

func (svc *userSvc) UpdateUser(ctx context.Context, id uint, user *params.UpdateUser) *views.Response {
	return nil
}
