package stategy

import (
	"crypto/md5"
	"hash"
	"math/rand"
	"net"

	"github.com/zhangjunfang/rpc/net/tcpPool"

	"github.com/zhangjunfang/networkLoadBalancing/util"
)

var Map util.Map

func init() {
	Map = util.NewMap()
}

type Stategy interface {
	Select(tcps []tcpPool.Pool) net.Conn
}

//轮循算法
func RoundRobin(tcps []tcpPool.Pool) net.Conn {
	if tcps != nil && len(tcps) > 1 {
		i := rand.Intn(len(tcps))
		return tcps[i].Get()
	} else if len(tcps) == 1 {
		return tcps[0]
	}
	return nil
}

//哈希算法
func Hash(tcps []tcpPool.Pool) net.Conn {
	if len(tcps) > 1 {
		addr, s := net.InterfaceAddrs()
		if s == nil {
			fmt.Println(len(addr))
			fmt.Println(addr)
			var hs hash.Hash = md5.New()
			hs.Write([]byte(addr[0].String()))
			var sum uint = 0
			for _, v := range hs.Sum(nil) {
				sum = sum + uint(v)
			}
			return tcps[sum%len(tcps)]
		}
	} else if len(tcps) == 1 {
		return tcps[0]
	}
	return nil

}

//最少连接算法
func LeastConnection(tcps []tcpPool.Pool) net.Conn {

}

//响应速度算法
func ResponseTime(tcps []tcpPool.Pool) net.Conn {

}

//加权法
func Weighted(tcps []tcpPool.Pool) net.Conn {

}
