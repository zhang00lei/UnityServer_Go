package excel2lua

import (
	"excel_lua/export_type"
	"excel_lua/util"
	"fmt"
)

// ExportCommonConfig
//  @Description: 导出CommonConfig配置
//  @param excelData excel数据
//  @param toPath 输入lua路径
//
func ExportCommonConfig(excelData [][]string, toPath string) {
	write, file := util.GetFileWriter(toPath, export_type.Lua)
	defer file.Close()
	write.WriteString("---@class CommonConfig\n")
	write.WriteString("local CommonConfig = {}\n")
	for i, row := range excelData {
		if i == 0 {
			continue
		}
		info := fmt.Sprintf("---@field %s %s @%s\n", row[0], row[2], row[1])
		write.WriteString(info)
		info = fmt.Sprintf("CommonConfig.%s = %s\n", row[0], row[3])
		write.WriteString(info)
	}
	write.WriteString("return CommonConfig")
	write.Flush()
}
