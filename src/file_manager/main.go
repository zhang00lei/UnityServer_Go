package main

import (
	_ "file_manager/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/config","static/download_dir")
	beego.Run()
}

