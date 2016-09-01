package main

import (
	"errors"
	"fmt"
	//"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/zhangjunfang/networkLoadBalancing/balance"
	"github.com/zhangjunfang/networkLoadBalancing/common"
	"github.com/zhangjunfang/networkLoadBalancing/stategy"
)

func main() {
	//ip,port,weight,retry,interval,times,rate_limit
	//221.204.14.157,80, 128,96 ,100,20, 10, 30,40/30;61.135.169.125,80 ,96,64 ,100,20, 10, 30,40/30
	pool, err := balance.GetTcpPool()
	//	if err != nil {
	//		fmt.Println(err)
	//	} else {
	//		fmt.Println(pool)
	//		fmt.Println(len(pool))
	//		for k, v := range pool {
	//			for j := 0; j < 512; j = j + 1 {
	//				mm, err := v.Pools[k].Get()
	//				//fmt.Println(k, "---", v.CoreTcp, mm.SetDeadline(time.Now().Add(60*time.Second)), "-----", v.ArgsStruct.Address, "====(", mm.LocalAddr().String(), "---", mm.RemoteAddr().Network(), ")===", err)
	//				fmt.Println(mm.RemoteAddr().String(), err)
	//			}
	//		}
	//	}
	common.MyError(err)
	var zz stategy.StategyAlgorithm = stategy.StategyAlgorithm(0)
	for {
		conn, err := zz.Select(pool)
		common.MyError(err)
		fmt.Println(conn)
		conn.Write([]byte("sdfsdfsdfdsfsdf"))
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
