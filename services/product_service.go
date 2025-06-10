package services
import (
	"go_restful_mvc/models"
	"go_restful_mvc/repositories"
)

type ProductService interface{
	Create(product *models.Product) error
	FindByID(id uint) (*models.Product, error)
	FindAll() ([]models.Product, error)
	Update(id uint, product *models.Product) error
	Delete(id uint) error
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) Create(product *models.Product) error {
	return s.repo.Create(product)
}

func (s *productService) FindByID(id uint) (*models.Product, error) {
	return s.repo.FindByID(id)
}

func (s *productService) FindAll() ([]models.Product, error) {
	return s.repo.FindAll()
}

func (s *productService) Update(id uint, product *models.Product) error {
	return s.repo.Update(id, product)
}

func (s *productService) Delete(id uint) error {
	return s.repo.Delete(id)
}