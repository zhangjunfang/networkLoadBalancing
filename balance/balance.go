package balance

import (
	"fmt"
	"time"
	//"fmt"
	"net"
	//"net/rpc"

	"github.com/zhangjunfang/networkLoadBalancing/args"
	"github.com/zhangjunfang/networkLoadBalancing/argsStruct"
	"github.com/zhangjunfang/networkLoadBalancing/config"
	"github.com/zhangjunfang/networkLoadBalancing/pool"
	"github.com/zhangjunfang/rpc/net/tcpPool"
)

type Address string

func GetTcpPool() (pools []*pool.Pool, err error) {
	addresses := args.ParseArgs()
	if len(addresses) > 0 {
		return poolAdapter(addresses)
	}
	addresses = config.ParseIni()
	if len(addresses) > 0 {
		return poolAdapter(addresses)
	}
	return pools, err
}

func (add Address) Factory() (net.Conn, error) {
	return net.DialTimeout("tcp", string(add), 60*time.Second)
}

func poolAdapter(addresses []*argsStruct.ArgsStruct) (pools []*pool.Pool, err error) {
	var p tcpPool.Pool
	var tcps []tcpPool.Pool
	for _, address := range addresses {
		add := Address(address.Address)
		fmt.Println("add:", add)
		p, err = tcpPool.NewChannelPool(address.CoreTcp, address.MaxTcp, add.Factory)
		if err == nil {
			tcps = append(tcps, p)
		} else {
			return nil, err
		}
		pools = append(pools, &pool.Pool{Pools: tcps,
			ArgsStruct: address,
		})
	}

	return pools, nil
}
