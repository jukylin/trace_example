package routers

import (
	"github.com/astaxie/beego"
	"github.com/hprose/hprose-golang/rpc"
	"github.com/jukylin/trace_example/components"
	"github.com/jukylin/trace_example/controllers"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
)

func init() {

	//分布式追踪
	trace := components.InitTrace()
	service := rpc.NewHTTPService()
	service.AddInstanceMethods(&controllers.MainController{})
	beego.Handler("/main", nethttp.Middleware(trace, service))
}
