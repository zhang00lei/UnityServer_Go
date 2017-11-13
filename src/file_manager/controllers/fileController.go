package controllers

import (
	"bufio"
	"file_manager/conf"
	"file_manager/util"
	"github.com/Luxurioust/excelize"
	"github.com/astaxie/beego"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var CurVersion string

type FileController struct {
	beego.Controller
}

func (c *FileController) Get() {
	c.TplName = "file_upload.html"
}

func (c *FileController) Post() {
	f, h, err := c.GetFile("myfile")
	defer func() {
		f.Close()
		UpdateVer()
		c.Data["CurVer"] = "当前版本：" + CurVersion
		c.TplName = "file_upload.html"
	}()
	if err!=nil{
		c.Ctx.WriteString(err.Error())
		return
	}
	if !strings.Contains(h.Filename, ".xlsx") && !strings.Contains(h.Filename, ".assetbundle") {
		c.Ctx.WriteString("文件类型不支持.")
		return
	} else {
		path := conf.FileDir + h.Filename
		c.SaveToFile("myfile", path)
		if strings.Contains(h.Filename, ".xlsx") {
			csvInfo, err := c.excelToCsv(path)
			os.Remove(path)
			if err == nil {
				fileName := conf.FileDir + strings.Replace(h.Filename, ".xlsx", ".txt", -1)
				if util.CheckFileIsExist(fileName) {
					os.Remove(fileName)
				}
				file, _ := os.OpenFile(fileName, os.O_CREATE, 0666)
				defer file.Close()
				wFile := bufio.NewWriter(file)
				wFile.WriteString(csvInfo)
				wFile.Flush()
				c.Data["FileName"] = strings.Replace(h.Filename, ".xlsx", ".txt", -1)
				c.Data["FilePath"] = fileName
			} else {
				c.Ctx.WriteString("上传excel文件错误：" + err.Error())
			}
		} else if strings.Contains(h.Filename, ".assetbundle") {
			c.Data["FileName"] = h.Filename
			c.Data["FilePath"] = ""
		}
	}
}

func (c *FileController) excelToCsv(excelPath string) (fileInfo string, err error) {
	xlsx, err := excelize.OpenFile(excelPath)
	if err != nil {
		return "", err
	}
	rows := xlsx.GetRows("Sheet1")
	var csvInfo string
	for _, row := range rows {
		for i, len := 0, len(row); i < len; i++ {
			if strings.Contains(row[i], ",") {
				row[i] = "\"" + row[i] + "\""
			} else if strings.Contains(row[i], "\"") {
				row[i] = strings.Replace(row[i], "\"", "\"\"", -1)
				row[i] = "\"" + row[i] + "\""
			}
			csvInfo += row[i] + ","
		}
		csvInfo += "\r\n"
	}
	return csvInfo, nil
}

func UpdateVer() {
	filePath := conf.FileDir + conf.SummaryFileName
	if util.CheckFileIsExist(filePath) {
		os.Remove(filePath)
	}
	dir_list, _ := ioutil.ReadDir(conf.FileDir)
	var fileInfos string
	for _, v := range dir_list {
		if v.Name() == conf.SummaryFileName {
			continue
		}
		var fileInfo string
		md5, _ := util.Md5SumFile(conf.FileDir + v.Name())
		fileInfo += v.Name() + "|"
		fileInfo += md5 + "|"
		fileInfo += strconv.FormatInt(v.Size(), 10)
		fileInfos += fileInfo + "\n"
	}
	file, _ := os.OpenFile(filePath, os.O_CREATE, 0666)
	defer file.Close()
	wFile := bufio.NewWriter(file)
	wFile.WriteString(fileInfos)
	wFile.Flush()
	CurVersion, _ = util.Md5SumFile(filePath)
}