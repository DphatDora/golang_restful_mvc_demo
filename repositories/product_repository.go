package repositories

import (
	"go_restful_mvc/models"
	"go_restful_mvc/config"
)

type ProductRepository interface {
	Create(product *models.Product) error
	FindByID(id uint) (*models.Product, error)
	FindAll() ([]models.Product, error)
	Update(id uint, product *models.Product) error
	Delete(id uint) error
}

type productRepository struct{}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

func (r *productRepository) Create(product *models.Product) error {
	return config.DB.Create(product).Error
}

func (r *productRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	err := config.DB.Preload("Category").First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	err := config.DB.Preload("Category").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) Update(id uint, product *models.Product) error {
	return config.DB.Model(&models.Product{}).Where("id = ?", id).Updates(product).Error
}

func (r *productRepository) Delete(id uint) error {
	return config.DB.Delete(&models.Product{}, id).Error
}