package service

import "github.com/SaRgEX/Diplom/pkg/repository"

type FavoriteService struct {
	repo repository.Favorite
}

func NewFavoriteService(repo repository.Favorite) *FavoriteService {
	return &FavoriteService{repo: repo}
}

func (s *FavoriteService) AddToFavorite(userId int, productId int) error {
	return s.repo.AddToFavorite(userId, productId)
}

func (s *FavoriteService) FindAll(userId int) ([]int, error) {
	return s.repo.FindAll(userId)
}
