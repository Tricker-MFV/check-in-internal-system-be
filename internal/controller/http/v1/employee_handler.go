package v1

import (
	"github.com/Tricker-MFV/checkin-internal-system/internal/service"
	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	employeeService service.EmployeeService
}

func NewEmployeeHandler(employeeService service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{employeeService: employeeService}
}

// @BasePath /api
// @Summary Get all employees
// @Description Get all employees
// @Tags Employee
// @Produce  json
// @Router /api/employee [get]
// @Success 200 {object} []model.EmployeeResponse
func (handler *EmployeeHandler) GetAll(c *gin.Context) {
	employees := handler.employeeService.GetAllEmployee(c)
	c.JSON(200, employees)
}
