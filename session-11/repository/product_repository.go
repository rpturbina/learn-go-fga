package repository

import "github.com/rpturbina/learn-go-fga/belajar-unit-test/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
}
