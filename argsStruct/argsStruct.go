package argsStruct

import (
	"time"
)

const (
	defa uint32 = iota //此值默认的值

	breaks //立即返回。并且返回错误信息

	retryNO //重试次数

	newObject //立即获取一个新的可以资源

	giveUp //立即放弃 什么也不做
)

type RetryStrategy uint32

type RateLimit struct {
	D time.Duration
	N int64
}

type ArgsStruct struct {
	Address   string //{ip:por}
	MaxTcp    int
	CoreTcp   int
	Weight    uint32        //weigth
	Retry     RetryStrategy //重试策略
	Interval  time.Duration //重试间隔周期
	Times     int           //重试次数
	RateLimit RateLimit     //访问频次
}
