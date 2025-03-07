package serviceimplement

import (
	"context"

	"github.com/Tricker-MFV/checkin-internal-system/internal/domain/model"
	"github.com/Tricker-MFV/checkin-internal-system/internal/repository"
	"github.com/Tricker-MFV/checkin-internal-system/internal/service"
)

type EmployeeService struct {
	employeeRepository repository.EmployeeRepository
}

func NewEmployeeService(employeeRepository repository.EmployeeRepository) service.EmployeeService {
	return &EmployeeService{employeeRepository: employeeRepository}
}

func (service EmployeeService) GetAllEmployee(ctx context.Context) []model.EmployeeResponse {
	employeesFromRepo := service.employeeRepository.GetAllQuery(ctx)

	employees := make([]model.EmployeeResponse, len(employeesFromRepo))
	for i, employee := range employeesFromRepo {
		employees[i] = model.EmployeeResponse{
			ID:       employee.ID,
			Username: employee.Username,
			Password: employee.Password,
		}
	}
	return employees
}
