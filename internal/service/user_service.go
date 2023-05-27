package service

import (
	"github.com/ZaiPeeKann/puregrade"
	"github.com/ZaiPeeKann/puregrade/internal/repository"
)

type UserService struct {
	repos *repository.Repository
}

func NewUserService(repos *repository.Repository) *UserService {
	return &UserService{repos: repos}
}

func (s *UserService) GetProfile(id int) (puregrade.Profile, error) {
	return s.repos.User.GetById(id)
}

func (s *UserService) FollowUser(id int, publisherId int) error {
	return s.repos.User.AddFollower(id, publisherId)
}

func (s *UserService) UnfollowUser(id int, publisherId int) error {
	return s.repos.User.DeleteFollower(id, publisherId)
}

func (s *UserService) Delete(id int, password string) error {
	return s.repos.User.Delete(id, password)
}
