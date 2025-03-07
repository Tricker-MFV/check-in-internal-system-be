//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/Tricker-MFV/checkin-internal-system/internal/controller"
	"github.com/Tricker-MFV/checkin-internal-system/internal/controller/http"
	v1 "github.com/Tricker-MFV/checkin-internal-system/internal/controller/http/v1"
	"github.com/Tricker-MFV/checkin-internal-system/internal/database"
	repositoryimplement "github.com/Tricker-MFV/checkin-internal-system/internal/repository/implement"
	serviceimplement "github.com/Tricker-MFV/checkin-internal-system/internal/service/implement"
	"github.com/google/wire"
)

var container = wire.NewSet(
	controller.NewApiContainer,
)

// may have grpc server in the future
var serverSet = wire.NewSet(
	http.NewServer,
)

// handler === controller | with service and repository layers to form 3 layers architecture
var handlerSet = wire.NewSet(
	v1.NewEmployeeHandler,
)

var serviceSet = wire.NewSet(
	serviceimplement.NewEmployeeService,
)

var repositorySet = wire.NewSet(
	repositoryimplement.NewEmployeeRepository,
)

func InitializeContainer(
	db database.Db,
) *controller.ApiContainer {
	wire.Build(serverSet, handlerSet, serviceSet, repositorySet, container)
	return &controller.ApiContainer{}
}
