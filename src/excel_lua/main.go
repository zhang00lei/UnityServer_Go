package main

import (
	"excel_lua/excel2lua"
	export_config_info "excel_lua/export-config-info"
	"flag"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
)

//  excel目录路径
var excelDirPath = flag.String("excelPath", "", "input excel path")

// 输出lua目录路径
var outLuaDirPath = flag.String("outLuaPath", "", "out lua path")

//  输出ConfigName路径
var outConfigNamePath = flag.String("outConfigNamePath", "", "out config name path")

//  输出ConfigPath路径
var outConfigPathPath = flag.String("outConfigPathPath", "", "out config path path")

func main() {
	flag.Parse()
	excelDirPath := *excelDirPath
	outLuaDirPath := *outLuaDirPath
	outConfigNamePath := *outConfigNamePath
	outConfigPathPath := *outConfigPathPath

	//excelDirPath = "E:\\wordcollection\\ConfigData\\Excel2Lua\\excel"
	//outLuaDirPath = "E:\\wordcollection\\Assets\\LuaScripts\\Config\\Tables"
	//outConfigNamePath = "E:\\wordcollection\\Assets\\LuaScripts\\Config\\ConfigNames.lua"
	//outConfigPathPath = "E:\\wordcollection\\Assets\\LuaScripts\\Config\\ConfigPath.lua"

	fmt.Println("excelPath", excelDirPath)
	fmt.Println("outLuaPath", outLuaDirPath)

	files, _ := ioutil.ReadDir(excelDirPath)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if strings.HasPrefix(file.Name(), "~$") {
			continue
		}
		excelPath := filepath.Join(excelDirPath, file.Name())

		fileSuffix := path.Ext(excelPath)
		filenameWithSuffix := filepath.Base(excelPath)
		fileName := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
		outLuaPath := filepath.Join(outLuaDirPath, fileName+".lua")

		f, err := excelize.OpenFile(excelPath)
		if err != nil {
			fmt.Println("打开excel失败", err)
			continue
		}
		rows, err := f.GetRows("Sheet1")
		if err != nil {
			fmt.Println("打开excel失败", err)
		}
		fmt.Println("toPath", outLuaPath)
		if strings.HasSuffix(excelPath, "CommonConfig.xlsx") {
			excel2lua.ExportCommonConfig(rows, outLuaPath)
			continue
		}
		excel2lua.ExportLuaConfig(rows, outLuaPath)
		export_config_info.SaveConfigInfo(fileName)
	}
	export_config_info.ExportConfigNames(outConfigNamePath)
	export_config_info.ExportConfigPath(outConfigPathPath)
}
