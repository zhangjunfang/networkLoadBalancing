package main

import (
	"fmt"
	"gobconn"
	"net"
	"reflect"
	"time"
)

type Info struct {
	Name string
	Age  int
	Job  string
	Hob  []string
}
type Test struct {
	Date    int
	Login   string
	Path    string
	Servers float64
	List    []string
	Dir     string
	Stream  bool
}

//初始化要发送的类型
func init() {
	go InitListen("tcp", ":2789")
	time.Sleep(1e9)
	gobconn.RegisterType(reflect.TypeOf(Info{}))
	gobconn.RegisterType(reflect.TypeOf(Test{}))
}
func main() {
	Test_rw()
	now := time.Now().Unix()
	Benchmark_rw()
	fmt.Println(time.Now().Unix())
	fmt.Println(now)
}
func Test_rw() {
	Dail("tcp", "127.0.0.1:2789", 1)
}
func Benchmark_rw() {
	Dail("tcp", "127.0.0.1:2789", 10000)
}

//创建tcp监听的端口
func InitListen(proto, addr string) {
	lis, err := net.Listen(proto, addr)
	if err != nil {
		fmt.Println("listen error,", err.Error())
		return
	}
	defer lis.Close()
	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("接入错误:", err)
			continue
		}
		go handle(conn)
	}
}

//链接处理逻辑
func handle(conn net.Conn) {
	con := gobconn.NewGobConnection(conn)
	defer con.Close()
	for {
		msg, err := con.Read()
		if err != nil {
			fmt.Println(con.RemoteAddr())
			fmt.Println("服务端ReadError:", err)
			return
		}
		err = con.Write(msg.Interface())
		if err != nil {
			fmt.Println("服务端WriteError:", err)
			msg.Recovery()
			return
		}
		msg.Recovery()
	}
}

//创建连接.
func Dail(proto, addr string, count int) {
	con, err := net.Dial(proto, addr)
	if err != nil {
		fmt.Println("客户端连接错误:", err)
		return
	}
	conn := gobconn.NewGobConnection(con)
	defer conn.Close()
	for i := 0; i < count; i++ {
		err = conn.Write(Info{
			"testing", 25, "IT",
			[]string{"backetball", "football"},
		})
		if err != nil {
			fmt.Println("客户端WriteError:", err)
			return
		}
		msg, err := conn.Read()
		if err != nil {
			fmt.Println("客户端ReadError:", err)
			return
		}
		fmt.Println(msg, msg.Interface())
		msg.Recovery()
	}
}
