package service

import (
	"github.com/Molsbee/blog/model"
	"github.com/Molsbee/blog/repository"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type AuthService interface {
	Authenticate(username, password string) (*model.ServiceUser, model.ApplicationError)
}

type authService struct {
	repo repository.ServiceUserRepository
}

func NewAuthService(db *gorm.DB) AuthService {
	return &authService{
		repo: repository.NewServiceUserRepository(db),
	}
}

func (a *authService) Authenticate(username, password string) (*model.ServiceUser, model.ApplicationError) {
	user := a.repo.FindByUsername(username)
	if user == nil || bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) != nil {
		return nil, model.ApiError{
			S: http.StatusUnauthorized,
			D: map[string]interface{}{
				"error": "authentication failed",
			},
		}
	}

	return user, nil
}
