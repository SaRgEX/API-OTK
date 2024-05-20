package service

import (
	model "github.com/SaRgEX/Diplom/Model"
	"github.com/SaRgEX/Diplom/pkg/repository"
)

type FavoriteService struct {
	repo repository.Favorite
}

func NewFavoriteService(repo repository.Favorite) *FavoriteService {
	return &FavoriteService{repo: repo}
}

func (s *FavoriteService) AddToFavorite(userId int, productId int) error {
	return s.repo.AddToFavorite(userId, productId)
}

func (s *FavoriteService) FindAll(userId int) (*model.FavoritesOutput, error) {
	return s.repo.FindAll(userId)
}

func (s *FavoriteService) RemoveFromFavorite(userId int, productId int) error {
	return s.repo.RemoveFromFavorite(userId, productId)
}
