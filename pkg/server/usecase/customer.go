package usecase

import (
	"ganja/pkg/server/entity"
	"ganja/pkg/server/repo"
)

type CustomerUsecaseInterface interface {
	GetById(string) (*entity.Customer, error)
}

type customerUsecase struct {
	CustomerRepo repo.CustomerRepoInterface
}

func GetCustomerUsecase() CustomerUsecaseInterface {
	return &customerUsecase{
		CustomerRepo: repo.GetCustomerRepo(),
	}
}

func (u *customerUsecase) GetById(id string) (*entity.Customer, error) {
	return u.CustomerRepo.GetById(id)
}
