package service

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ServiceGroup struct {
	CasUserService
}

var ServiceGroupApp = new(ServiceGroup)

var (
	jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService
)
