package main

import (
	"github.com/zhangjunfang/networkLoadBalancing/args"
	_ "github.com/zhangjunfang/networkLoadBalancing/config"
)

func main() {
	//C:/workspace/golang/src/github.com/zhangjunfang/networkLoadBalancing/config
	//C:/workspace/golang/src/github.com/zhangjunfang/networkLoadBalancing/conf/
	args.ParseArgs()
}
