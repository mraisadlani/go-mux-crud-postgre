package repository

import "github.com/vanilla/go-mux-postgre/api/entities"

type ProductRepository interface {
	FindAll() ([]entities.Product, error)
	FindById(uint64) (entities.Product, error)
	Save(entities.Product) (bool, error)
	Update(uint64, entities.Product) (bool, error)
	Delete(uint64) (bool, error)
}