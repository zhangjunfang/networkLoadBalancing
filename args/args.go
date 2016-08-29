package args

import (
	"flag"
	"strings"
	"time"

	"github.com/zhangjunfang/networkLoadBalancing/argsStruct"
)

var addresses = flag.String("addresses", "61.135.169.121:80,42.236.4.32:443", "default address ")
var weight = flag.Uint("weight", 0, "weight set")
var interval = flag.Int64("interval", 0, "interval period")
var times = flag.Int("times", 0, "repeat times")
var retry = flag.Int("retry", 0, "0:default ,1:break ,2:retry,3:new resource,4:nothing")

func ParseArgs() *argsStruct.ArgsStruct {
	flag.Parse()

	Address := strings.Split(*addresses, ",")

	return &argsStruct.ArgsStruct{
		Addresses: Address,
		Weight:    uint32(*weight),
		Retry:     argsStruct.RetryStrategy(uint32(*retry)),
		Interval:  time.Duration(*interval),
		Times:     *times,
	}
}
