package main

import (
	"excel_lua/excel2json"
	"excel_lua/excel2lua"
	export_config_info "excel_lua/export-config-info"
	"flag"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/fs"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

//  excel目录路径
var excelDirPath = flag.String("excelPath", "", "input excel path")

// 输出lua目录路径
var outLuaDirPath = flag.String("outLuaPath", "", "out lua path")

// 输出json目录路径
var outJsonDirPath = flag.String("outJsonPath", "", "out json path")

//  输出ConfigName路径
var outConfigNamePath = flag.String("outConfigNamePath", "", "out config name path")

//  输出ConfigPath路径
var outConfigPathPath = flag.String("outConfigPathPath", "", "out config path path")

func toJsonInfo(file fs.FileInfo, excelDirPath, outLuaDirPath, outJsonDirPath string) {
	defer waitGroup.Done()
	if file.IsDir() {
		return
	}
	if strings.HasPrefix(file.Name(), "~$") {
		return
	}
	excelPath := filepath.Join(excelDirPath, file.Name())

	fileSuffix := path.Ext(excelPath)
	filenameWithSuffix := filepath.Base(excelPath)
	fileName := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	outLuaPath := filepath.Join(outLuaDirPath, fileName+".lua")

	f, err := excelize.OpenFile(excelPath)
	if err != nil {
		fmt.Println("打开excel失败", err)
		return
	}
	rows := f.GetRows(fileName)
	//fmt.Println("toPath", outLuaPath)
	if strings.HasSuffix(excelPath, "CommonConfig.xlsx") {
		excel2lua.ExportCommonConfig(rows, outLuaPath)
		return
	}

	outJsonPath := filepath.Join(outJsonDirPath, fileName+".json")
	fmt.Println("toPath ", outJsonPath)
	excel2json.ExportJsonConfig(rows, outJsonPath)

	excel2lua.ExportLuaConfig(rows, outLuaPath)
	export_config_info.SaveConfigInfo(fileName)
}

var waitGroup sync.WaitGroup

func main() {
	flag.Parse()
	excelDirPath := *excelDirPath
	outLuaDirPath := *outLuaDirPath
	outJsonDirPath := *outJsonDirPath
	outConfigNamePath := *outConfigNamePath
	outConfigPathPath := *outConfigPathPath

	//excelDirPath = "E:\\project\\codes\\Excel2JsonTools\\Table\\1"
	//outLuaDirPath = "E:\\project\\codes\\Assets\\DevHere\\Scripts\\Lua\\Configs\\Tables"
	//outJsonDirPath = "E:\\project\\codes\\Assets\\DevHere\\Datas\\Json"
	//outConfigNamePath = "E:\\project\\codes\\Assets\\DevHere\\Scripts\\Lua\\Framework\\ConfigManager\\ConfigNames.lua"
	//outConfigPathPath = "E:\\project\\codes\\Assets\\DevHere\\Scripts\\Lua\\Framework\\ConfigManager\\ConfigPath.lua"

	fmt.Println("excelPath", excelDirPath)
	fmt.Println("outLuaPath", outLuaDirPath)
	fmt.Println("outConfigNamePath:", outConfigNamePath)
	fmt.Println("outConfigPathPath:", outConfigPathPath)

	files, _ := ioutil.ReadDir(excelDirPath)
	for _, file := range files {
		waitGroup.Add(1)
		go toJsonInfo(file, excelDirPath, outLuaDirPath, outJsonDirPath)
	}
	waitGroup.Wait()
	export_config_info.ExportConfigNames(outConfigNamePath)
	export_config_info.ExportConfigPath(outConfigPathPath)
}
