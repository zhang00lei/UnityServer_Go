package routers

import (
	"file_manager/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.FileController{})
	beego.Router("/GetCurVer", &controllers.GetVerController{})
}
