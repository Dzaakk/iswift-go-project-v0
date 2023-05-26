package admin

import (
	dto "iswift-go-project/internal/admin/dto"
	entity "iswift-go-project/internal/admin/entity"
	repository "iswift-go-project/internal/admin/repository"
)

type AdminUseCase interface {
	FindAll(offset int, limit int) []entity.Admin
	FindById(id int) (*entity.Admin, error)
	FindByEmail(email string) (*entity.Admin, error)
	Create(dto dto.AdminRequestBody) (*entity.Admin, error)
	Update(dto dto.AdminRequestBody) (*entity.Admin, error)
	Delete(id int) error
}

type AdminUseCaseImpl struct {
	repository repository.AdminRepository
}

// Create implements AdminUseCase.
func (*AdminUseCaseImpl) Create(dto dto.AdminRequestBody) (*entity.Admin, error) {
	panic("unimplemented")
}

// Delete implements AdminUseCase.
func (*AdminUseCaseImpl) Delete(id int) error {
	panic("unimplemented")
}

// FindAll implements AdminUseCase.
func (*AdminUseCaseImpl) FindAll(offset int, limit int) []entity.Admin {
	panic("unimplemented")
}

// FindByEmail implements AdminUseCase.
func (usecase *AdminUseCaseImpl) FindByEmail(email string) (*entity.Admin, error) {
	return usecase.repository.FindByEmail(email)
}

// FindById implements AdminUseCase.
func (*AdminUseCaseImpl) FindById(id int) (*entity.Admin, error) {
	panic("unimplemented")
}

// Update implements AdminUseCase.
func (*AdminUseCaseImpl) Update(dto dto.AdminRequestBody) (*entity.Admin, error) {
	panic("unimplemented")
}

func NewAdminUseCase(repository repository.AdminRepository) AdminUseCase {
	return &AdminUseCaseImpl{repository}
}
