package service

import (
	model "github.com/SaRgEX/Diplom/Model"
	"github.com/SaRgEX/Diplom/pkg/repository"
)

type AddressService struct {
	repo repository.Address
}

func NewAddressService(repo repository.Address) *AddressService {
	return &AddressService{repo: repo}
}

func (s *AddressService) Create(address model.Address) (int, error) {
	return s.repo.Create(address)
}
