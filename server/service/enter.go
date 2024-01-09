package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/cate"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	CateServiceGroup    cate.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
