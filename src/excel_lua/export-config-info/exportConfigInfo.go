package export_config_info

import (
	"excel_lua/export_type"
	"excel_lua/util"
	"fmt"
	"strings"
	"sync"
)

var allConfigName []string

var mutex sync.Mutex

func SaveConfigInfo(configName string) {
	mutex.Lock()
	allConfigName = append(allConfigName, configName)
	mutex.Unlock()
}

func ExportConfigNames(outPath string) {
	write, file := util.GetFileWriter(outPath, export_type.Lua)
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
	write, file := util.GetFileWriter(outPath, export_type.Lua)
	defer file.Close()
	write.WriteString("local ConfigPath = {\n")
	for _, configName := range allConfigName {
		nameTemp := strings.Replace(configName, "Config", "", 1)
		info := fmt.Sprintf("    %s = \"Configs.Tables.%s\",\n", nameTemp, configName)
		write.WriteString(info)
	}
	write.WriteString("}\n")
	write.WriteString("return ConfigPath")
	write.Flush()
}
