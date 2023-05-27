package service

import (
	"github.com/ZaiPeeKann/puregrade"
	"github.com/ZaiPeeKann/puregrade/internal/repository"
)

type ProductService struct {
	repos *repository.Repository
}

func NewProductService(repos *repository.Repository) *ProductService {
	return &ProductService{repos: repos}
}

func (s *ProductService) GetAll(page int, filter map[string]string) ([]puregrade.Product, error) {
	return s.repos.Product.GetAll(page, filter)
}

func (s *ProductService) GetOneByID(id int) (puregrade.Product, error) {
	return s.repos.Product.GetOneByID(id)
}

func (s *ProductService) Create(product puregrade.CreateProductDTO) (int, error) {
	return s.repos.Product.Create(product)
}

func (s *ProductService) AddGenres(id int, g []int) error {
	return s.repos.AddGenres(id, g)
}

func (s *ProductService) AddPlatforms(id int, p []int) error {
	return s.repos.AddPlatforms(id, p)
}

func (s *ProductService) DeleteGenres(id int, g []int) error {
	return s.repos.DeleteGenres(id, g)
}

func (s *ProductService) DeletePlatforms(id int, p []int) error {
	return s.repos.DeletePlatforms(id, p)
}

func (s *ProductService) Delete(id int) error {
	return s.repos.Product.Delete(id)
}
