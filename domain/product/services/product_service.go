package services

import (
	"hms-go/domain/product/models"
	"hms-go/domain/product/repositories"
)

type ProductService interface {
	GetAllProduct() ([]models.Product, error)
	GetProductById(id uint) (*models.Product, error)
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(id uint) error
	TestService() string
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetAllProduct() ([]models.Product, error) {
	return s.repo.GetAllProduct()
}

func (s *productService) GetProductById(id uint) (*models.Product, error) {
	return s.repo.GetProductById(id)
}

func (s *productService) CreateProduct(product *models.Product) error {
	return s.repo.CreateProduct(product)
}

func (s *productService) UpdateProduct(product *models.Product) error {
	return s.repo.UpdateProduct(product)
}

func (s *productService) DeleteProduct(id uint) error {
	return s.repo.DeleteProduct(id)
}

func (s *productService) TestService() string {

	useRepo := s.repo.TestRepo()
	resp := ", This is from service"

	combineResp := useRepo + resp
	return combineResp
}
