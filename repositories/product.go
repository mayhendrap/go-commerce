package repositories

import (
	"go-commerce/entities"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() ([]entities.Product, error)
	FindByID(id string) (entities.Product, error)
	Create(product entities.Product) (entities.Product, error)
	Update(product entities.Product) (entities.Product, error)
	Delete(id string) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db: db}
}

func (p *productRepository) FindAll() ([]entities.Product, error) {
	var products []entities.Product
	err := p.db.Find(&products).Error
	if err != nil {
		return products, err
	}
	return products, nil
}

func (p *productRepository) FindByID(id string) (entities.Product, error) {
	var product entities.Product
	err := p.db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (p *productRepository) Create(product entities.Product) (entities.Product, error) {
	err := p.db.Create(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (p *productRepository) Update(product entities.Product) (entities.Product, error) {
	err := p.db.Save(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (p *productRepository) Delete(id string) error {
	err := p.db.Where("id = ?", id).Delete(&entities.Product{}).Error
	return err
}
