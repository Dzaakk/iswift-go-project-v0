package admin

import (
	entity "iswift-go-project/internal/admin/entity"

	"gorm.io/gorm"
)

type AdminRepository interface {
	FindAll(offset int, limit int) []entity.Admin
	FindById(id int) (*entity.Admin, error)
	FindByEmail(email string) (*entity.Admin, error)
	Create(entity entity.Admin) (*entity.Admin, error)
	Update(entity entity.Admin) (*entity.Admin, error)
	Delete(entity entity.Admin) error
}

type AdminRepositoryImpl struct {
	db *gorm.DB
}

// Create implements AdminRepository.
func (*AdminRepositoryImpl) Create(entity entity.Admin) (*entity.Admin, error) {
	panic("unimplemented")
}

// Delete implements AdminRepository.
func (*AdminRepositoryImpl) Delete(entity entity.Admin) error {
	panic("unimplemented")
}

// FindAll implements AdminRepository.
func (*AdminRepositoryImpl) FindAll(offset int, limit int) []entity.Admin {
	panic("unimplemented")
}

// FindByEmail implements AdminRepository.
func (adminRepository *AdminRepositoryImpl) FindByEmail(email string) (*entity.Admin, error) {
	var admin entity.Admin

	if err := adminRepository.db.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}

	return &admin, nil
}

// FindById implements AdminRepository.
func (*AdminRepositoryImpl) FindById(id int) (*entity.Admin, error) {
	panic("unimplemented")
}

// Update implements AdminRepository.
func (*AdminRepositoryImpl) Update(entity entity.Admin) (*entity.Admin, error) {
	panic("unimplemented")
}

func NewAdminRepository() AdminRepository {
	return &AdminRepositoryImpl{}
}
