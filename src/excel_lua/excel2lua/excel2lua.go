package excel2lua

import (
	"bufio"
	"excel_lua/export_type"
	"excel_lua/util"
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

//
// ExportLuaConfig
//  @Description: 导出普通lua配置
//  @param excelData excel相关数据
//  @param outLuaPath 输出lua路径
//
func ExportLuaConfig(excelData [][]string, outLuaPath string) {
	write, file := util.GetFileWriter(outLuaPath, export_type.Lua)
	defer file.Close()
	fileSuffix := path.Ext(outLuaPath)
	filenameWithSuffix := filepath.Base(outLuaPath)
	fileName := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	info := fmt.Sprintf("---@class %s\n", fileName)
	write.WriteString(info)

	exportLuaNote(excelData, write)

	exportLuaData(excelData, write, fileName)

	write.Flush()
}

//
//  exportLuaNote
//  @Description: 生成lua文件注解
//  @param excelData excel相关数据
//  @param write
//
func exportLuaNote(excelData [][]string, write *bufio.Writer) {
	for i, fieldName := range excelData[2] {
		if strings.HasPrefix(fieldName, "#") {
			continue
		}
		if len(fieldName) <= 0 {
			continue
		}
		fieldType := excelData[3][i]
		if fieldType == "int" || fieldType == "float" || fieldType == "number" || fieldType == "idx1" {
			fieldType = "number"
		} else if fieldType == "bool" {
			fieldType = "boolean"
		}
		fieldDes := excelData[1][i]
		fieldDes = strings.ReplaceAll(fieldDes, "\n", " ")
		info := fmt.Sprintf("---@field %s %s @%s\n", fieldName, fieldType, fieldDes)
		write.WriteString(info)
	}
}

//
//  exportLuaData
//  @Description: 生成lua文件数据
//  @param excelData excel相关数据
//  @param write
//  @param fileName lua文件名字
//
func exportLuaData(excelData [][]string, write *bufio.Writer, fileName string) {
	info := fmt.Sprintf("local %s = {\n", fileName)
	write.WriteString(info)

	for i, rowData := range excelData {
		if i < 4 {
			continue
		}
		info = fmt.Sprintf("    [%s] = {\n", rowData[0])
		write.WriteString(info)
		dataLength := len(excelData[3])
		for j := 0; j < dataLength; j++ {
			cell := ""
			if j < len(rowData) {
				cell = rowData[j]
			}
			fieldName := excelData[2][j]
			if strings.HasPrefix(fieldName, "#") {
				continue
			}
			fieldType := excelData[3][j]
			if fieldType == "int" || fieldType == "number" || fieldType == "float" {
				if cell == "" {
					cell = "0"
				}
				info = fmt.Sprintf("        %s = %s,\n", fieldName, cell)
			} else if fieldType == "string" {
				if strings.Contains(cell, "\"") {
					cell = strings.Replace(cell, "\"", `\"`, -1)
				} else if strings.Contains(cell, "\n") {
					cell = strings.Replace(cell, "\n", `\n`, -1)
				}
				info = fmt.Sprintf("        %s = \"%s\",\n", fieldName, cell)
			} else if fieldType == "table" {
				info = fmt.Sprintf("        %s = %s,\n", fieldName, cell)
			} else if fieldType == "boolean" || fieldType == "bool" {
				temp := "false"
				if cell == "1" {
					temp = "true"
				}
				info = fmt.Sprintf("        %s = %s,\n", fieldName, temp)
			} else {
				info = ""
			}
			if info != "" {
				write.WriteString(info)
			}
		}
		write.WriteString("    },\n")
	}
	write.WriteString("}\n")
	info = fmt.Sprintf("return %s", fileName)
	write.WriteString(info)
}
