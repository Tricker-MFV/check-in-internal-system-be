package controller

import "github.com/Tricker-MFV/checkin-internal-system/internal/controller/http"

type ApiContainer struct {
	HttpServer *http.Server
}

func NewApiContainer(httpServer *http.Server) *ApiContainer {
	return &ApiContainer{HttpServer: httpServer}
}
