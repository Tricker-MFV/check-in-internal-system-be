package repositoryimplement

import (
	"context"

	"github.com/Tricker-MFV/checkin-internal-system/internal/database"
	"github.com/Tricker-MFV/checkin-internal-system/internal/domain/entity"
	"github.com/Tricker-MFV/checkin-internal-system/internal/repository"
	"gorm.io/gorm"
)

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db database.Db) repository.EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (repo EmployeeRepository) GetAllQuery(ctx context.Context) []entity.Employee {
	var employees []entity.Employee
	result := repo.db.WithContext(ctx).Find(&employees)
	if result.Error != nil {
		return nil
	}

	return employees
}
