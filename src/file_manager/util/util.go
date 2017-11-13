package util

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func CheckFileIsExist(filename string) bool {
	exist := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func Md5SumFile(file string) (md5Value string, err error) {
	data, err := os.Open(file)
	defer data.Close()
	md5Hash := md5.New()
	io.Copy(md5Hash, data)
	md5Value = hex.EncodeToString(md5Hash.Sum(nil))
	return
}
