package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/attendant"
	"github.com/flipped-aurora/gin-vue-admin/server/router/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Autocode autocode.RouterGroup
	Attendant attendant.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
