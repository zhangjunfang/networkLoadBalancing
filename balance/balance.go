package balance

import (
	"net"

	"github.com/zhangjunfang/networkLoadBalancing/args"
	"github.com/zhangjunfang/networkLoadBalancing/config"
	"github.com/zhangjunfang/rpc/net/tcpPool"
)

type Address string

func GetTcpPool() (tcps []tcpPool.Pool, err error) {
	addresses := args.ParseArgs().Addresses
	if len(addresses) > 0 {
		return commCode(addresses)
	}
	addresses = config.ParseIni().Addresses
	if len(config.ParseIni().Addresses) > 0 {
		return commCode(addresses)
	}
	return tcps, err
}

func (add Address) Factory() (net.Conn, error) {
	return net.Dial("tcp", string(add))
}

func commCode(addresses []string) (tcps []tcpPool.Pool, err error) {
	var pool tcpPool.Pool
	for _, address := range addresses {
		add := Address(address)
		pool, err = tcpPool.NewChannelPool(32, 64, add.Factory)
		if err == nil {
			tcps = append(tcps, pool)
		} else {
			tcps = append(tcps, nil)
		}
	}
	return tcps, err
}
