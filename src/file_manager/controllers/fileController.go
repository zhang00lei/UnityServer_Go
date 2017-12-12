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
	"fmt"
	"io"
)

var CurVersion string

type FileController struct {
	beego.Controller
}

func (c *FileController) Get() {
	c.TplName = "file_upload.html"
}

func (c *FileController) IsContainFile(fileName string) bool{
	fileList :=[]string{".xlsx",".assetbundle"}
	for _,v :=range fileList{
		if strings.Contains(fileName,v){
			return true
		}
	}
	return false
}

func (c *FileController) Post() {
	f, h, err := c.GetFile("myfile")
	defer func() {
		f.Close()
		UpdateVer()
		c.Data["CurVer"] = "当前版本：" + CurVersion
		c.TplName = "file_upload.html"
	}()
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}
	path := conf.FileDir + h.Filename
	if (c.IsContainFile(h.Filename)==true) {
		c.SaveToFile("myfile", path)
	}else {
		c.Ctx.WriteString("文件格式不支持")
	}
	if strings.Contains(h.Filename, ".xlsx") {
		csvInfo, err := c.excelToInfo(path)
		os.Remove(path)
		if err == nil {
			var fileSuffix string
			if conf.GenerateCSV {
				fileSuffix = ".csv"
			} else {
				fileSuffix = ".lua"
			}
			fileName := conf.FileDir + strings.Replace(h.Filename, ".xlsx", fileSuffix, -1)
			if util.CheckFileIsExist(fileName) {
				os.Remove(fileName)
			}
			file, _ := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0666)
			defer file.Close()
			io.WriteString(file,csvInfo)
			c.Data["FileName"] = strings.Replace(h.Filename, ".xlsx", fileSuffix, -1)
			c.Data["FilePath"] = fileName
		} else {
			c.Ctx.WriteString("上传excel文件错误：" + err.Error())
		}
	} else if strings.Contains(h.Filename, ".assetbundle") {
		c.Data["FileName"] = h.Filename
		c.Data["FilePath"] = ""
	}
}

func (c *FileController) generaCSV(rows [][] string ) string {
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
	return csvInfo
}
func (c *FileController) generaLua(rows [][] string,fileName string) string {
	var content string
	for row,count:=2,len(rows);row<count ;row++{
		var rowInfo string = "\t{"
		for i,len:=0,len(rows[row]);i<len;i++ {
			if strings.TrimSpace(rows[1][i])!="" {
				colInfo:=rows[row][i]
				colInfo = strings.Replace(colInfo,"\"","\\\"",-1)
				rowInfo = rowInfo + fmt.Sprintf("%s = \"%s\",\t", rows[1][i], colInfo)
			}
		}
		content=content+rowInfo+"},\n"
	}
	info := fmt.Sprintf("local %s = {\t\n%s}\nreturn %s",fileName,content,fileName)
	return info
}

func (c *FileController) excelToInfo(excelPath string) (fileInfo string, err error) {
	xlsx, err := excelize.OpenFile(excelPath)
	if err != nil {
		return "", err
	}
	rows := xlsx.GetRows("Sheet1")
	if conf.GenerateCSV{
		return c.generaCSV(rows),nil
	}else {
		fileName:=c.getFileName(excelPath)
		return c.generaLua(rows,fileName),nil
	}
}

func (c *FileController) getFileName(filePath string) string  {
	stringTemp := strings.Split(filePath,"/")
	fileName:= strings.Split(stringTemp[len(stringTemp)-1],".")
	return fileName[0]
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
	file, _ := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0666)
	defer file.Close()
	wFile := bufio.NewWriter(file)
	wFile.WriteString(fileInfos)
	wFile.Flush()
	CurVersion, _ = util.Md5SumFile(filePath)
}