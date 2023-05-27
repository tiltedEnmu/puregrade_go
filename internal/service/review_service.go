package service

import (
	"errors"

	"github.com/ZaiPeeKann/puregrade"
	"github.com/ZaiPeeKann/puregrade/internal/repository"
)

type ReviewService struct {
	repos *repository.Repository
}

func NewReviewService(repos *repository.Repository) *ReviewService {
	return &ReviewService{repos: repos}
}

func (s *ReviewService) GetAll(page int, productId int) ([]puregrade.Review, error) {
	return s.repos.Review.GetAll(page, productId)
}

func (s *ReviewService) GetOneByID(id int) (puregrade.Review, error) {
	return s.repos.Review.GetOneByID(id)
}

func (s *ReviewService) Create(review puregrade.Review) (int, error) {
	return s.repos.Review.Create(review)
}

func (s *ReviewService) Update(id int, title, body string) error {
	return s.repos.Review.Update(id, title, body)
}

func (s *ReviewService) Delete(id, userId int) error {
	rewiew, err := s.repos.Review.GetOneByID(id)
	if err != nil {
		return errors.New("review not found")
	}
	if rewiew.Author.Id != id {
		return errors.New("forbidden")
	}
	return s.repos.Review.Delete(id)
}
