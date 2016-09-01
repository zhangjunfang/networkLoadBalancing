package config

import (
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/config"
	"github.com/zhangjunfang/networkLoadBalancing/argsStruct"
	"github.com/zhangjunfang/networkLoadBalancing/common"
)

func ParseIni() (args []*argsStruct.ArgsStruct) {

	Ini := config.IniConfig{}
	workPath := os.Args[0]
	workDir := filepath.Dir(workPath)
	os.Chdir(workDir)
	configer, err := Ini.Parse(path.Join(strings.Replace(workDir, "\\", "/", -1), "config", "nlb.conf"))
	common.MyError(err)
	// data = ip,port,weight,retry,interval,times,rate_limit;ip1,port1,weight1,retry1,interval1,times1,rate_limit1
	param := strings.Split(strings.TrimSpace(configer.String("data")), ";")
	for _, v := range param {
		arg := strings.Split(v, ",")
		rate := strings.Split(strings.TrimSpace(arg[6]), "/")
		weight, e := strconv.Atoi(strings.TrimSpace(arg[2]))
		common.MyError(e)
		retry, e := strconv.Atoi(strings.TrimSpace(arg[3]))
		common.MyError(e)
		interval, e := strconv.Atoi(strings.TrimSpace(arg[4]))
		common.MyError(e)
		times, e := strconv.Atoi(strings.TrimSpace(arg[5]))
		common.MyError(e)
		d, e := strconv.Atoi(strings.TrimSpace(rate[0]))
		common.MyError(e)
		n, e := strconv.Atoi(strings.TrimSpace(rate[1]))
		common.MyError(e)
		g := &argsStruct.ArgsStruct{
			Address:  strings.TrimSpace(arg[0]) + ":" + strings.TrimSpace(arg[1]),
			Weight:   uint32(weight),
			Retry:    argsStruct.RetryStrategy(uint32(retry)),
			Interval: time.Duration(int64(interval)),
			Times:    times,
			RateLimit: argsStruct.RateLimit{
				D: time.Duration(int64(d)),
				N: int64(n),
			},
		}
		args = append(args, g)
	}
	return args
}
