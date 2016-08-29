package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/zhangjunfang/networkLoadBalancing/balance"
	"github.com/zhangjunfang/rpc/net/tcpPool"
)

func main() {

	pool, err := balance.GetTcpPool()
	if err != nil {
		fmt.Println(err)
	} else {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for {
			j := r.Intn(2)
			fmt.Println(j)
			fmt.Println(pool[j], "rand pool ")
		}

		return
		fmt.Println(pool)
		fmt.Println(len(pool))
		var i int = 1 //count  go number
		for k, p := range pool {
			for {
				go func(k int, p tcpPool.Pool) {
					i++
					fmt.Println("==i===", i)
					c, err := p.Get()
					fmt.Println(c.RemoteAddr().Network(), "--------------", c.RemoteAddr().String())
					fmt.Println(k, c, err)
					time.Sleep(1 * time.Millisecond)
				}(k, p)
			}
		}
	}
	time.Sleep(1 * time.Second)
}

func xxx() {
	ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(time.Second*10))
	t, ok := ctx.Deadline()
	if ok {
		fmt.Println(time.Now())
		fmt.Println(t.String())
	}
	go func(ctx context.Context) {
		fmt.Println(ctx.Value("Test"))
		<-ctx.Done()
		fmt.Println(ctx.Err())
	}(ctx)
	if ctx.Err() == nil {
		time.Sleep(11e9)
	}
	if ctx.Err() != nil {
		fmt.Println("已经退出了")
	}
	cancelFunc()
}
