package args

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/zhangjunfang/networkLoadBalancing/argsStruct"
	"github.com/zhangjunfang/networkLoadBalancing/common"
)

func ParseArgs() (args []*argsStruct.ArgsStruct) {
	flag.Parse()
	arg := flag.Args()

	if len(arg) > 0 {
		param := strings.Split(strings.TrimSpace(arg[0]), ";")
		for _, v := range param {
			arg = strings.Split(v, ",")
			if len(arg) != 9 {
				fmt.Println(" param len :", len(arg))
				common.MyError(errors.New("invlidte data  format"))
			}
			rate := strings.Split(strings.TrimSpace(arg[8]), "/")
			maxtcp, e := strconv.Atoi(strings.TrimSpace(arg[2]))
			common.MyError(e)
			coretcp, e := strconv.Atoi(strings.TrimSpace(arg[3]))
			common.MyError(e)
			weight, e := strconv.Atoi(strings.TrimSpace(arg[4]))
			common.MyError(e)
			retry, e := strconv.Atoi(strings.TrimSpace(arg[5]))
			common.MyError(e)
			interval, e := strconv.Atoi(strings.TrimSpace(arg[6]))
			common.MyError(e)
			times, e := strconv.Atoi(strings.TrimSpace(arg[7]))
			common.MyError(e)
			d, e := strconv.Atoi(strings.TrimSpace(rate[0]))
			common.MyError(e)
			n, e := strconv.Atoi(strings.TrimSpace(rate[1]))
			common.MyError(e)
			g := &argsStruct.ArgsStruct{
				Address:  strings.TrimSpace(arg[0]) + ":" + strings.TrimSpace(arg[1]),
				MaxTcp:   maxtcp,
				CoreTcp:  coretcp,
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
	} else {
		fmt.Println(`
		invalid data format
	    Parameter format: "ip,port,maxtcp,coretcp,weight,retry,interval,times,rate_limit;ip1,port1,maxtcp,coretcp,weight1,retry1,interval1,times1,rate_limit1"
	    data type: ip string  port int maxtcp int coretcp int  weight int retry int interval int rate_limit int/int 10/20
	`)
		os.Exit(0)
	}
	return args
}
