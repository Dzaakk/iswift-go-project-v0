package product_category

import (
	dto "iswift-go-project/internal/product_category/dto"
	entity "iswift-go-project/internal/product_category/entity"
	repository "iswift-go-project/internal/product_category/repository"
)

type ProductCategoryUseCase interface {
	FindAll(offset, limit int) []entity.ProductCategory
	FindById(id int) (*entity.ProductCategory, error)
	Create(dto dto.ProductCategoryRequestBody) (*entity.ProductCategory, error)
	Update(dto dto.ProductCategoryRequestBody) (*entity.ProductCategory, error)
	Delete(id int) error
}

type ProductCategoryUseCaseImpl struct {
	repository repository.ProductcategoryRepository
}

// Create implements ProductCategoryUseCase.
func (usecase *ProductCategoryUseCaseImpl) Create(dto dto.ProductCategoryRequestBody) (*entity.ProductCategory, error) {
	productCategoryEntity := entity.ProductCategory{
		Name:        dto.Name,
		CreatedByID: dto.CreatedBy,
	}
	productCategory, err := usecase.repository.Create(productCategoryEntity)

	if err != nil {
		return nil, err
	}

	return productCategory, nil
}

// Delete implements ProductCategoryUseCase.
func (*ProductCategoryUseCaseImpl) Delete(id int) error {
	panic("unimplemented")
}

// FindAll implements ProductCategoryUseCase.
func (usecase *ProductCategoryUseCaseImpl) FindAll(offset int, limit int) []entity.ProductCategory {
	return usecase.repository.FindAll(offset, limit)
}

// FindById implements ProductCategoryUseCase.
func (usecase *ProductCategoryUseCaseImpl) FindById(id int) (*entity.ProductCategory, error) {
	return usecase.repository.FindById(id)
}

// Update implements ProductCategoryUseCase.
func (*ProductCategoryUseCaseImpl) Update(dto dto.ProductCategoryRequestBody) (*entity.ProductCategory, error) {
	panic("unimplemented")
}

func NewProductCategoryUseCase(repository repository.ProductcategoryRepository) ProductCategoryUseCase {
	return &ProductCategoryUseCaseImpl{repository}
}
