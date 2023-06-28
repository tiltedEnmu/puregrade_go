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

func (s *UserService) GetProfile(id int64) (puregrade.Profile, error) {
	return s.repos.User.GetById(id)
}

func (s *UserService) FollowUser(id int64, publisherId int64) error {
	return s.repos.User.AddFollower(id, publisherId)
}

func (s *UserService) UnfollowUser(id int64, publisherId int64) error {
	return s.repos.User.DeleteFollower(id, publisherId)
}

func (s *UserService) Delete(id int64, password string) error {
	return s.repos.User.Delete(id, password)
}
