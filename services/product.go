package services

import (
	"go-commerce/dtos/request"
	"go-commerce/entities"
	"go-commerce/repositories"
)

type ProductService interface {
	Create(req request.ProductRequest) (entities.Product, error)
	FindAll() ([]entities.Product, error)
	FindByID(productID string) (entities.Product, error)
	Update(id string, req request.ProductRequestUpdate) (entities.Product, error)
	Delete(id string) (bool, error)
}

type productService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(productRepository repositories.ProductRepository) *productService {
	return &productService{productRepository: productRepository}
}

func (p *productService) Create(req request.ProductRequest) (entities.Product, error) {
	var err error
	product := entities.Product{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
	}
	product, err = p.productRepository.Create(product)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (p *productService) FindAll() ([]entities.Product, error) {
	products, err := p.productRepository.FindAll()
	if err != nil {
		return products, err
	}
	return products, err
}

func (p *productService) FindByID(productID string) (entities.Product, error) {
	product, err := p.productRepository.FindByID(productID)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (p *productService) Update(id string, input request.ProductRequestUpdate) (entities.Product, error) {
	oldProduct, err := p.productRepository.FindByID(id)
	if err != nil {
		return oldProduct, err
	}
	updatedProduct, err := p.productRepository.Update(oldProduct)
	if err != nil {
		return updatedProduct, err
	}
	return updatedProduct, nil
}

func (p *productService) Delete(id string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
