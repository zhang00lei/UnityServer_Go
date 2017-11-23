package main

import (
	"file_manager/controllers"
	_ "file_manager/routers"
	"github.com/astaxie/beego"
)

func main() {
	controllers.UpdateVer()
	beego.SetStaticPath("/config", "static/download_dir")
	beego.Run()
}
