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

func ExportCSharpInit(outPath, generateFileName string) {
	write, file := util.GetFileWriter(outPath, export_type.CSharp)
	defer file.Close()
	write.WriteString(`using MxxM.GameClient;
using UnityEngine;

public partial class DataTableManager
{
    private TextAsset textAsset;

    private void InitData()
    {`)
	write.WriteString("\n")
	for _, configName := range allConfigName {
		if generateFileName != "" && !strings.Contains(generateFileName, configName+",") {
			continue
		}
		infoTemp := fmt.Sprintf("        textAsset = Context.Game.Loader.LoadAsset<TextAsset>(\"Assets/DevHere/Datas/Json/%s.json\");\n", configName)
		write.WriteString(infoTemp)
		infoTemp = fmt.Sprintf("        T%sHelper.InitData(textAsset.text);\n", configName)
		write.WriteString(infoTemp)
	}
	write.WriteString(`    }
}`)
	write.Flush()
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
