package config

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/astaxie/beego/config"
	"github.com/zhangjunfang/networkLoadBalancing/argsStruct"
)

func ParseIni() *argsStruct.ArgsStruct {

	Ini := config.IniConfig{}
	workPath := os.Args[0]
	workDir := filepath.Dir(workPath)
	os.Chdir(workDir)
	configer, err := Ini.Parse(path.Join(strings.Replace(workDir, "\\", "/", -1), "config", "nlb.conf"))
	if err != nil {
		panic(err)
	}

	return &argsStruct.ArgsStruct{
		Addresses: configer.DefaultStrings("addresses", []string{"0.0.0.0:3002"}),
		Weight:    uint32(configer.DefaultInt("weight", 0)),
		Retry:     argsStruct.RetryStrategy(uint32(configer.DefaultInt("retry", 0))),
		Interval:  time.Duration(configer.DefaultInt("interval", 0)),
		Times:     configer.DefaultInt("times", 0),
	}
}
