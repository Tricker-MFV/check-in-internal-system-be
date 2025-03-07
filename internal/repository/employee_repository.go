package repository

import (
	"context"

	"github.com/Tricker-MFV/checkin-internal-system/internal/domain/entity"
)

type EmployeeRepository interface {
	GetAllQuery(c context.Context) []entity.Employee
}
