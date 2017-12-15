package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"apk_manager/conf"
)

type FileController struct {
	beego.Controller
}
func (c *FileController) IsContainFile(fileName string) bool{
	fileList :=[]string{".apk"}
	for _,v :=range fileList{
		if strings.Contains(fileName,v){
			return true
		}
	}
	return false
}
func (c *FileController) Get() {
	c.TplName = "file_upload.html"
}

func (c *FileController) Post(){
	f, h, err := c.GetFile("myfile")
	defer func() {
		f.Close()
		c.TplName = "file_upload.html"
	}()
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}
	path := conf.FileDir + h.Filename
	if (c.IsContainFile(h.Filename)) {
		c.SaveToFile("myfile", path)
		c.Data["FileName"] = "http://192.168.62.29:8080/ApkSubmit"+path
		c.Data["FilePath"] = path
	}else {
		c.Ctx.WriteString("文件格式不支持")
	}
}