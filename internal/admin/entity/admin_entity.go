package admin

import (
	"database/sql"

	"gorm.io/gorm"
)

type Admin struct {
	ID          int64          `json:"id"`
	Email       string         `json:"email"`
	Name        string         `json:"name"`
	Password    string         `json:"password"`
	CreatedBy   *Admin         `gorm:"foreignKey:CreatedByID;references:ID"`
	CreatedByID int64          `json:"created_by"`
	UpdatedBy   *Admin         `gorm:"foreignKey:UpdatedByID;references:ID"`
	UpdatedByID *int64         `json:"updated_by"`
	CreatedAt   sql.NullTime   `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
