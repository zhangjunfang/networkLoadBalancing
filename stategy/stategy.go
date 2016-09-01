package stategy

import (
	"crypto/md5"
	"hash"
	"math/rand"
	"net"
	"sort"

	"github.com/zhangjunfang/networkLoadBalancing/pool"
)

const (
	RoundRobins = iota
	Hashs
	Weighteds
)

type StategyAlgorithm int

type Stategy interface {
	Select(tcps []*pool.Pool) (net.Conn, error)
}

func (s StategyAlgorithm) Select(tcps []*pool.Pool) (net.Conn, error) {
	switch int(s) {
	case RoundRobins:
		return RoundRobin(tcps)
	case Hashs:
		return Hash(tcps)
	case Weighteds:
		return Weighted(tcps)
	default:
		return nil, nil
	}
}

//轮循算法
func RoundRobin(tcps []*pool.Pool) (net.Conn, error) {

	if tcps != nil {
		if len(tcps) > 1 {
			i := rand.Intn(len(tcps))
			j := rand.Intn(len(tcps[i].Pools))
			return tcps[i].Pools[j].Get()
		} else if len(tcps) == 1 {
			j := rand.Intn(len(tcps[0].Pools))
			return tcps[0].Pools[j].Get()
		}
	}
	return nil, nil
}

//哈希算法
func Hash(tcps []*pool.Pool) (net.Conn, error) {
	if tcps != nil {
		if len(tcps) > 1 {
			addr, s := net.InterfaceAddrs()
			if s == nil {
				i := rand.Intn(len(tcps))
				var hs hash.Hash = md5.New()
				hs.Write([]byte(addr[i].String()))
				var sum int = 0
				for _, v := range hs.Sum(nil) {
					sum = sum + int(v)
				}
				j := rand.Intn(len(tcps[i].Pools))
				return tcps[sum%len(tcps)].Pools[j].Get()
			}
		} else if len(tcps) == 1 {
			j := rand.Intn(len(tcps[0].Pools))
			return tcps[0].Pools[j].Get()
		}
	}
	return nil, nil

}

//最少连接算法
func LeastConnection(tcps []*pool.Pool) (net.Conn, error) {

	return nil, nil
}

//响应速度算法
func ResponseTime(tcps []*pool.Pool) (net.Conn, error) {
	return nil, nil
}

//加权法
func Weighted(tcps []*pool.Pool) (net.Conn, error) {
	if tcps != nil && len(tcps) > 1 {
		m := len(tcps)
		sort.Sort(pool.P(tcps))
		i := uint32(rand.Int31n(100))
		for k, _ := range tcps {
			if k == 0 {
				if 0 <= i && i < tcps[k].Weight {
					j := rand.Intn(len(tcps[k].Pools))
					return tcps[k].Pools[j].Get()
				}
			}
			if k == (m - 1) {
				if tcps[k-1].Weight <= i && i < tcps[k].Weight {
					j := rand.Intn(len(tcps[k].Pools))
					return tcps[k].Pools[j].Get()
				}
			}
			if tcps[k].Weight <= i {
				j := rand.Intn(len(tcps[k].Pools))
				return tcps[k].Pools[j].Get()
			}
		}
		return nil, nil
	} else if len(tcps) == 1 {
		j := rand.Intn(len(tcps[0].Pools))
		return tcps[0].Pools[j].Get()
	}
	return nil, nil
}
