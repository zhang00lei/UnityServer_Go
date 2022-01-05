package util

import (
	"bufio"
	"fmt"
	"os"
)

func GetFileWriter(outPath string) (*bufio.Writer, *os.File) {
	if _, err := os.Stat(outPath); err == nil {
		os.Remove(outPath)
	}
	file, err := os.OpenFile(outPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
		return nil, nil
	}
	write := bufio.NewWriter(file)
	write.WriteString("---this file is generate by tools,do not modify it.\n")
	return write, file
}
