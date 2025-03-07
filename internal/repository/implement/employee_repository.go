package repositoryimplement

import (
	"context"

	"github.com/Tricker-MFV/checkin-internal-system/internal/database"
	"github.com/Tricker-MFV/checkin-internal-system/internal/domain/entity"
	"github.com/Tricker-MFV/checkin-internal-system/internal/repository"
	"github.com/jmoiron/sqlx"
)

type EmployeeRepository struct {
	db *sqlx.DB
}

func NewEmployeeRepository(db database.Db) repository.EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (repo EmployeeRepository) GetAllQuery(ctx context.Context) []entity.Employee {
	var employees []entity.Employee
	query := "SELECT * FROM employees"
	err := repo.db.SelectContext(ctx, &employees, query)
	if err != nil {
		return nil
	}

	return employees
}
