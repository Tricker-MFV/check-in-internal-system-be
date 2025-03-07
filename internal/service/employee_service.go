package service

import (
	"context"

	"github.com/Tricker-MFV/checkin-internal-system/internal/domain/model"
)

type EmployeeService interface {
	GetAllEmployee(ctx context.Context) []model.EmployeeResponse
}
