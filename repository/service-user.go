package repository

import (
	"github.com/Molsbee/blog/model"
	"github.com/jinzhu/gorm"
	"log"
)

type ServiceUserRepository interface {
	FindByUsername(username string) *model.ServiceUser
}

type serviceUserRepository struct {
	db *gorm.DB
}

func NewServiceUserRepository(db *gorm.DB) ServiceUserRepository {
	return &serviceUserRepository{
		db: db,
	}
}

func (ur *serviceUserRepository) FindByUsername(username string) *model.ServiceUser {
	serviceUser := model.ServiceUser{}
	err := ur.db.Where("username = ?", username).Find(&serviceUser).Error
	if err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			log.Printf("failed to query service user table - %s\n", err)
		}
		return nil
	}

	return &serviceUser
}
