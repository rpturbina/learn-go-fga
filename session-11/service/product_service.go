package service

import (
	"errors"

	"github.com/rpturbina/learn-go-fga/belajar-unit-test/entity"
	"github.com/rpturbina/learn-go-fga/belajar-unit-test/repository"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (service ProductService) GetOneProduct(id string) (*entity.Product, error) {
	product := service.Repository.FindById(id)

	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}
