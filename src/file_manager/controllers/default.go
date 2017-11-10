package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/Luxurioust/excelize"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type FileController struct{
	beego.Controller
}

func (c *FileController) Get(){
	c.TplName = "file_upload.html"
}

func (c *FileController) Post()  {
	fmt.Println("test")
	f,h,_:=c.GetFile("myfile")
	path:= "static/download_dir/"+ h.Filename
	fmt.Println(path)
	f.Close()
	c.SaveToFile("myfile",path)
	c.excelToCsv(path)
	c.Data["FileName"] = h.Filename
	c.Data["FilePath"] = path
	c.TplName = "file_upload.html"
}

func (c *FileController) excelToCsv (excelPath string) error {
	xlsx,err:=excelize.OpenFile(excelPath)
	if err!=nil{
		return err
	}
	rows:=xlsx.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Println(colCell)
		}
		fmt.Println()
	}
	return nil
}