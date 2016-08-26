package args

import (
	//"expvar"
	"flag"
	"fmt"
	"strings"
)

var addresses = flag.String("addresses", "0.0.0.0:3002", "default address ")

//
var Address []string

func ParseArgs() {
	flag.Parse()
	Address = strings.Split(*addresses, ",")
	//	for _, v := range Address {
	//		fmt.Println(v)
	//	}
}
