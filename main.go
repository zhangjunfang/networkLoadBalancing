package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net"
	"runtime"
	"time"

	"github.com/zhangjunfang/networkLoadBalancing/balance"
	"github.com/zhangjunfang/networkLoadBalancing/common"
	"github.com/zhangjunfang/networkLoadBalancing/stategy"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//ip,port,weight,retry,interval,times,rate_limit
	//61.135.169.125,80 ,96,64 ,80,20, 10, 30,40/30;221.204.14.157,80, 128,96 ,20,20, 10, 30,40/30
	pool, err := balance.GetTcpPool()
	common.MyError(err)
	var zz stategy.StategyAlgorithm = stategy.StategyAlgorithm(1)
	i := int64(0)
	for {
		i = i + 1
		conn, err := zz.Select(pool)
		common.MyError(err)
		fmt.Println(conn.RemoteAddr().String(), "----", i)
		//go conn.Read()
		//go conn.Write([]byte("sdfsdfsdfdsfsdf"))
		conn.Close()
	}
	return
}

func Lookup(host string) (string, error) {
	addrs, err := net.LookupHost(host)
	if err != nil {
		return "", err
	}
	if len(addrs) < 1 {
		return "", errors.New("unknown host")
	}
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return addrs[rd.Intn(len(addrs))], nil
}
