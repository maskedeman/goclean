package repository

import (
	"goclean/internal/models"
	interfaces "goclean/pkg/v1"

	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) interfaces.RepoInterface {
	return &Repo{db}
}

func (repo *Repo) Create(user models.User) (models.User, error) {
	err := repo.db.Create(&user).Error

	return user, err
}

func (repo *Repo) Read(id string) (models.User, error) {
	var user models.User
	err := repo.db.Where("id=?", id).First(&user).Error
	return user, err
}

func (repo *Repo) Update(user models.User) error {
	var dbUser models.User
	if err := repo.db.Where("id=?", user.ID).First(&dbUser).Error; err != nil {
		return err
	}
	dbUser.Name = user.Name
	err := repo.db.Save(dbUser).Error

	return err
}

func (repo *Repo) Delete(id string) error {
	err := repo.db.Where("id=?", id).Delete(&models.User{}).Error
	return err
}

func (repo *Repo) GetByEmail(email string) (models.User, error) {
	var user models.User
	err := repo.db.Where("email=?", email).First(&user).Error
	return user, err
}
