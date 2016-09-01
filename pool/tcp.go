package pool

import (
	"github.com/zhangjunfang/networkLoadBalancing/argsStruct"
	"github.com/zhangjunfang/rpc/net/tcpPool"
)

type Pool struct {
	Pools []tcpPool.Pool
	*argsStruct.ArgsStruct
}
type P []*Pool

// Len is the number of elements in the collection.
func (p P) Len() int {
	return len(p)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (p P) Less(i, j int) bool {
	return (p)[i].Weight < (p)[j].Weight
}

// Swap swaps the elements with indexes i and j.
func (p P) Swap(i, j int) {
	(p)[i], (p)[j] = (p)[j], (p)[i]
}
