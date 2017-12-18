package routers

import (
	"apk_manager/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/ApkSubmit",&controllers.FileController{})
    beego.Router("/",&controllers.FileInfoController{})
}
