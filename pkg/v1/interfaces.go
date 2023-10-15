package v1

import "goclean/internal/models"

type RepoInterface interface {
	Create(models.User) (models.User, error)

	Read(id string) (models.User, error)

	Update(models.User) error

	Delete(id string) error
	GetByEmail(email string) (models.User, error)
}

type UseCaseInterface interface {
	Create(models.User) (models.User, error)
	Read(id string) (models.User, error)
	Update(models.User) error
	Delete(id string) error
}
