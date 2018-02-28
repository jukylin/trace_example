package main

import (
	"github.com/astaxie/beego"
	_ "github.com/jukylin/trace_example/routers"
)

func main() {
	beego.Run()
}
