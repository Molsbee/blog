package model

import "time"

type ServiceUser struct {
	ID        int       `gorm:"column:id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	FirstName string    `gorm:"column:first_name"`
	LastName  string    `gorm:"column:last_name"`
}

func (ServiceUser) TableName() string {
	return "service_users"
}
