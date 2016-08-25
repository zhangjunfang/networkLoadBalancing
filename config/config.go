package config

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego/config"
	//"github.com/juju/errors"
)

var Ini config.IniConfig
var configer config.Configer

func init() {
	Ini = config.IniConfig{}
	workPath := os.Args[0]
	workDir := filepath.Dir(workPath)
	fmt.Println(workDir)
	configer, err := Ini.Parse(path.Join(strings.Replace(workDir, "\\", "/", -1), "configs", "nlb.conf"))
	//	fmt.Println("############", err.Error())
	if err != nil {
		foundPath := path.Join(strings.Replace(workDir, "\\", "/", -1), "config")
		fmt.Println("--foundPath--", foundPath)
		filepath.Walk(foundPath, filter)
	}
	fmt.Println(configer.String("runmode"))
}
func filter(url string, info os.FileInfo, err error) error {
	if !info.IsDir() && info.Name() == "nlb.conf" {
		configer, err := Ini.Parse(path.Join(strings.Replace(url, "\\", "/", -1), "nlb.conf"))
		fmt.Println(configer)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(configer.String("zhangsan"))
	}
	return nil
}
