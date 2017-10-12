package routers

import (
	"trace_example/components"
	"trace_example/controllers"
	"github.com/astaxie/beego"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/hprose/hprose-golang/rpc"
)

func init() {

	//分布式追踪
	trace := components.InitTrace()
	service := rpc.NewHTTPService()
	service.AddInstanceMethods(&controllers.MainController{})
	beego.Handler("/main", nethttp.Middleware(trace, service))
}
