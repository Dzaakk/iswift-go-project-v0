package product

import (
	entity "iswift-go-project/internal/product/entity"
	"iswift-go-project/pkg/utils"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(offset, limit int) []entity.Product
	FindById(id int) (*entity.Product, error)
	Create(entity entity.Product) (*entity.Product, error)
	Update(entity entity.Product) (*entity.Product, error)
	Delete(entity entity.Product) error
}

type ProductRepositoryImpl struct {
	db gorm.DB
}

// Create implements ProductRepository.
func (Repository *ProductRepositoryImpl) Create(entity entity.Product) (*entity.Product, error) {

}

// Delete implements ProductRepository.
func (*ProductRepositoryImpl) Delete(entity entity.Product) error {
	panic("unimplemented")
}

// FindAll implements ProductRepository.
func (Repository *ProductRepositoryImpl) FindAll(offset int, limit int) []entity.Product {
	var products []entity.Product

	Repository.db.Scopes(utils.Paginate(offset, limit)).Find(&products)
}

// FindById implements ProductRepository.
func (*ProductRepositoryImpl) FindById(id int) (*entity.Product, error) {
	panic("unimplemented")
}

// Update implements ProductRepository.
func (*ProductRepositoryImpl) Update(entity entity.Product) (*entity.Product, error) {
	panic("unimplemented")
}

func NewProductRepository(db gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db}
}
