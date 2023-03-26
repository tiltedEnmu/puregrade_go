package service

import (
	puregrade "github.com/ZaiPeeKann/auth-service_pg/internal/models"
	"github.com/ZaiPeeKann/auth-service_pg/internal/repository"
)

type ProductService struct {
	repos *repository.Repository
}

func NewProductService(repos *repository.Repository) *ProductService {
	return &ProductService{repos: repos}
}

func (s *ProductService) GetAll() ([]puregrade.Product, error) {
	return s.repos.Product.GetAll()
}

func (s *ProductService) GetOneByID(id int) (puregrade.Product, error) {
	return s.repos.Product.GetOneByID(id)
}

func (s *ProductService) Create(product puregrade.Product) (int, error) {
	return s.repos.Product.Create(product)
}

func (s *ProductService) Delete(id int) error {
	return s.repos.Product.Delete(id)
}
