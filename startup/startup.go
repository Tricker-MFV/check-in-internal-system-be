package startup

import (
	"github.com/Tricker-MFV/checkin-internal-system/internal"
	"github.com/Tricker-MFV/checkin-internal-system/internal/controller"
	"github.com/Tricker-MFV/checkin-internal-system/internal/database"
)

func registerDependencies() *controller.ApiContainer {
	// Open database connection
	db := database.Open()

	return internal.InitializeContainer(db)
}

func Execute() {
	container := registerDependencies()
	container.HttpServer.Run()
}
