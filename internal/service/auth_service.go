package service

import (
	"Interface_droch_3/internal/model"
	"Interface_droch_3/internal/repository"
	"fmt"
)

type AuthService struct {
	repo repository.StorageUsers
}

func NewAuthService(repo repository.StorageUsers) *AuthService {
	return &AuthService{repo: repo}
}

func (r *AuthService) Set(user *model.User) error {
	return r.repo.Set(user)
}
func (r *AuthService) Get(id int64) (*model.User, error) {
	return r.repo.Get(id)
}
func (r *AuthService) Check(id int64) (bool, error) {
	return r.repo.Check(id)
}
func (r *AuthService) Delete(id int64) error {
	exists, err := r.repo.Check(id)
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("Пользователь с ID %d не найден", id)
	}

	return r.repo.Delete(id)
}
func (r *AuthService) GetAllId() []int64 {
	return r.repo.GetAllId()
}
