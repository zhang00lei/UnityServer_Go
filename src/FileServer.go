package main
import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func getCurrentPath()string{
	s,_ := exec.LookPath(os.Args[0])
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

func main() {
	path:=getCurrentPath()
	fmt.Println("Current Path:"+path)
	http.Handle("/", http.FileServer(http.Dir("./")))
	err := http.ListenAndServe(":8000",nil)
	if err!=nil {
		fmt.Println(err)
		panic(err)
	}
}