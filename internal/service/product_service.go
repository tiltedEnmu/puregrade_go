package service

import (
	"github.com/ZaiPeeKann/puregrade"
	"github.com/ZaiPeeKann/puregrade/internal/repository"
)

const Limit int = 20

type ProductService struct {
	repos *repository.Repository
}

func NewProductService(repos *repository.Repository) *ProductService {
	return &ProductService{repos: repos}
}

func (s *ProductService) GetAll(filter puregrade.ProductFilter) ([]puregrade.Product, error) {
	var offset int = 0
	if filter.Page > 1 {
		offset = (filter.Page - 1) * Limit
	}
	return s.repos.Product.GetAll(offset, Limit, filter.Genre, filter.Platform, filter.OrderBy, filter.IsAsc)
}

func (s *ProductService) GetOneByID(id int64) (puregrade.Product, error) {
	return s.repos.Product.GetOneByID(id)
}

func (s *ProductService) Create(product puregrade.CreateProductDTO) (int64, error) {
	return s.repos.Product.Create(product)
}

func (s *ProductService) AddGenres(id int64, genres []int64) error {
	return s.repos.AddGenres(id, genres)
}

func (s *ProductService) AddPlatforms(id int64, platforms []int64) error {
	return s.repos.AddPlatforms(id, platforms)
}

func (s *ProductService) DeleteGenres(id int64, genres []int64) error {
	return s.repos.DeleteGenres(id, genres)
}

func (s *ProductService) DeletePlatforms(id int64, platforms []int64) error {
	return s.repos.DeletePlatforms(id, platforms)
}

func (s *ProductService) Delete(id int64) error {
	return s.repos.Product.Delete(id)
}
