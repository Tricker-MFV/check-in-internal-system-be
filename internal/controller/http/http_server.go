package http

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	v1 "github.com/Tricker-MFV/checkin-internal-system/internal/controller/http/v1"
)

type Server struct {
	employeeHandler *v1.EmployeeHandler
}

func NewServer(employeeHandler *v1.EmployeeHandler) *Server {
	return &Server{employeeHandler: employeeHandler}
}

func (s *Server) Run() {
	router := gin.New()
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	httpServerInstance := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	v1.MapRoutes(router, s.employeeHandler)
	err := httpServerInstance.ListenAndServe()
	if err != nil {
		return
	}
	fmt.Println("Server running at " + httpServerInstance.Addr)
}
