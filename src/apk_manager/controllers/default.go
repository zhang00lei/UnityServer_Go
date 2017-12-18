package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"apk_manager/conf"
	"io/ioutil"
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
		c.Data["FileName"] ="192.168.62.29:8080/"+ path
		c.Data["FilePath"] = path
	}else {
		c.Ctx.WriteString("文件格式不支持")
	}
}

type FileInfoController struct {
	beego.Controller
}

func (c *FileInfoController)Get()  {
	dir_list,e:=ioutil.ReadDir(conf.FileDir)
	if e!=nil{
		c.Ctx.WriteString(e.Error())
		return 
	}
	fileName:=[]string{}
	for _,v:=range dir_list {
		fileName = append(fileName, v.Name())
	}
	c.Data["FileInfo"] = fileName
	c.Data["FilePath"] = conf.FileDir
	c.TplName = "file_list.html"
}