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
var err error

func init() {
	Ini = config.IniConfig{}
	workPath := os.Args[0]
	workDir := filepath.Dir(workPath)
	configer, err = Ini.Parse(path.Join(strings.Replace(workDir, "\\", "/", -1), "configs", "nlb.conf"))
	if err != nil {
		return
		//foundPath := path.Join(strings.Replace(workDir, "\\", "/", -1), "config")
		//filepath.Walk(foundPath, filter)
	}
	//fmt.Println("configer is not null ?", configer)
	//fmt.Println(configer.String("runmode"))
}
func filter(url string, info os.FileInfo, err error) error {
	if !info.IsDir() && info.Name() == "nlb.conf" {
		if configer != nil {
			temp, err := Ini.Parse(url)
			if err != nil {
				fmt.Println(err)
			}
			//			configer.Set(temp.)
			fmt.Println("temp configer:=", temp)
		} else {
			configer, err = Ini.Parse(url)
			fmt.Println("configer:=", configer)
			if err != nil {
				fmt.Println(err)
			}
		}

		//fmt.Println(configer.String("zhangsan"))
	}
	return nil
}
