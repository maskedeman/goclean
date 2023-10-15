package usecase

import (
	"errors"
	"goclean/internal/models"
	interfaces "goclean/pkg/v1"

	"gorm.io/gorm"
)

type UseCase struct {
	repo interfaces.RepoInterface
}

func New(repo interfaces.RepoInterface) interfaces.UseCaseInterface {
	return &UseCase{repo}
}

func (uc *UseCase) Create(user models.User) (models.User, error) {
	if _, err := uc.repo.GetByEmail(user.Email); !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.User{}, errors.New("Already existing email")

	}
	return uc.repo.Create(user)
}

func (uc *UseCase) Read(id string) (models.User, error) {
	var user models.User
	var err error

	if user, err = uc.repo.Read(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, errors.New("No existing user")
		}
		return models.User{}, err
	}
	return user, nil
}

func (uc *UseCase) Update(updateUser models.User) error {
	var user models.User
	var err error

	if user, err = uc.Read(string(updateUser.ID)); err != nil {
		return err
	}
	if user.Email != updateUser.Email {
		return errors.New("Can't change Email")
	}
	err = uc.repo.Update(updateUser)
	if err != nil {

	}
	return nil
}

func (uc *UseCase) Delete(id string) error {
	var err error
	if _, err = uc.Read(id); err != nil {
		return err
	}

	err = uc.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
