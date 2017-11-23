package controllers

import "github.com/astaxie/beego"

type GetVerController struct {
	beego.Controller
}

func (c *GetVerController) Get() {
	c.Ctx.WriteString(CurVersion)
}
