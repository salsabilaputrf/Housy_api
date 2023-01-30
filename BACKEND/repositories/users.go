package repositories

import (
	"housy/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Preload("ListAs").Find(&users).Error

	return users, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.Preload("ListAs").First(&user, ID).Error

	return user, err
}

func (r *repository) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := r.db.Preload("ListAs").First(&user, "username=?", username).Error

	return user, err
}