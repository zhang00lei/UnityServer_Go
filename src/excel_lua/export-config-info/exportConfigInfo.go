package export_config_info

import (
	"excel_lua/util"
	"fmt"
	"strings"
)

var allConfigName []string

func SaveConfigInfo(configName string) {
	allConfigName = append(allConfigName, configName)
}

func ExportConfigNames(outPath string) {
	write, file := util.GetFileWriter(outPath)
	defer file.Close()
	write.WriteString("local ConfigNames = {\n")
	for _, configName := range allConfigName {
		nameTemp := strings.Replace(configName, "Config", "", 1)
		info := fmt.Sprintf("    %s = \"%s\",\n", nameTemp, nameTemp)
		write.WriteString(info)
	}
	write.WriteString("}\n")
	write.WriteString("return ConfigNames")
	write.Flush()
}

func ExportConfigPath(outPath string) {
	write, file := util.GetFileWriter(outPath)
	defer file.Close()
	write.WriteString("local ConfigPath = {\n")
	for _, configName := range allConfigName {
		nameTemp := strings.Replace(configName, "Config", "", 1)
		info := fmt.Sprintf("    %s = \"Config.Tables.%s\",\n", nameTemp, configName)
		write.WriteString(info)
	}
	write.WriteString("}\n")
	write.WriteString("return ConfigPath")
	write.Flush()
}
